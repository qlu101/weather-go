package weather

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/shawntoffel/darksky"
    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
    "log"
    "net/http"
    "strconv"
    "time"
)

const (
	googleGeoKey string = "AIzaSyAo2OnaGrE1JC92gn1cKSS9MiThtaJCpP0"
	darkSkyKey string = "28896f6afb4e54d741e4d57dbe5d2749"
)

func init() {
    log.Print("Listening on http requests...")
    http.HandleFunc("/weather", getHandler)
}

/**
* Handler for http GET request. Supports two ways:
  1) https://weathergo2.appspot.com/weather?zipcode=92602
  2) https://weathergo2.appspot.com/weather?latitude=33.6846&longitude=-117.8265
*/
func getHandler(w http.ResponseWriter, r *http.Request) {
    //===== Get zipcode parameter
    zipcode, found := getParamOptional(r, "zipcode")
    if found {
		latt, lngt, err := getGeoLocation(w, r, zipcode)
	    if err != nil {
	        return
	    }
	    
		getResponse(w, r, latt, lngt)
		return
    }

    //===== Get parameters
    latt, err := getParamFloat(w, r, "latitude")
    if err != nil {
        return
    }

    lngt, err := getParamFloat(w, r, "longitude")
    if err != nil {
        return
    }

	getResponse(w, r, latt, lngt)
	return
}

/**
* Look up latitude and longitude by zipcode
*/
func getGeoLocation(w http.ResponseWriter, r *http.Request, zipcode string) (float64, float64, error) {
    fm := "https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s"
    url := fmt.Sprintf(fm, zipcode, googleGeoKey)
    // log.Print(url)

    ctx := appengine.NewContext(r)
    client := urlfetch.Client(ctx)
    resp, err := client.Get(url)
    if err != nil {
        setError(w, http.StatusFailedDependency, err.Error())
        return 0.0, 0.0, err
    }
    defer resp.Body.Close()

	geo := GeoLocation{}
    err = json.NewDecoder(resp.Body).Decode(&geo)
    if err != nil {
        setError(w, http.StatusFailedDependency, err.Error())
        return 0.0, 0.0, err
    }
	
	latt := geo.Results[0].Geometry.Location.Lat
	lngt := geo.Results[0].Geometry.Location.Lng
	addr := geo.Results[0].FormattedAddress

    msg := fmt.Sprintf("zipcode=%s, latitude=%f, longitude=%f, address=%s", zipcode, latt, lngt, addr)
    log.Print(msg)
	
    return latt, lngt, nil
}

/**
* Get weather information of the location with given latitude and longitude
*/
func getResponse(w http.ResponseWriter, r *http.Request, latt float64, lngt float64) {
    wi := WeatherInfo{
        Description: "Weather report for the past 7 days (1 week)",
        Note:        "Powered by Google Maps and Dark Sky (https://darksky.net/poweredby)",
    }

    //===== Retrieve past 7 day's weather data
    for i := -7; i <= -1; i++ {
        utcNow := time.Now().AddDate(0, 0, i).UTC().Unix()
        fm := "https://api.darksky.net/forecast/%s/%f,%f,%d?exclude=currently,hourly,minutely,alerts,flags"
        url := fmt.Sprintf(fm, darkSkyKey, latt, lngt, utcNow)
        // log.Print(url)

        ctx := appengine.NewContext(r)
        client := urlfetch.Client(ctx)
        resp, err := client.Get(url)
        if err != nil {
            setError(w, http.StatusFailedDependency, err.Error())
            return
        }
        defer resp.Body.Close()

        response := darksky.ForecastResponse{}
        err = json.NewDecoder(resp.Body).Decode(&response)
        if err != nil {
            setError(w, http.StatusFailedDependency, err.Error())
            return
        }

        wi.Days = append(wi.Days, response.Daily.Data[0])
    }

    wi.DayCount = len(wi.Days)
    
    //===== Send back response
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    json.NewEncoder(w).Encode(&wi)
}

func getParamOptional(r *http.Request, key string) (string, bool) {
    params, ok := r.URL.Query()[key]
    if !ok || len(params) < 1 {
        return "", false
    }

    param := string(params[0])
    return param, true
}

func getParamFloat(w http.ResponseWriter, r *http.Request, key string) (float64, error) {
    var errMsg string
    params, ok := r.URL.Query()[key]
    if !ok || len(params) < 1 {
        errMsg = fmt.Sprintf("Url Parameter '%s' is missing", key)
        setError(w, http.StatusBadRequest, errMsg)
        return 0.0, errors.New(errMsg)
    }

    param, err := strconv.ParseFloat(string(params[0]), 64)
    if err != nil {
        setError(w, http.StatusBadRequest, err.Error())
        return 0.0, err
    }

    return param, nil
}

func setError(w http.ResponseWriter, errCode int, errMsg string) {
    log.Print(errMsg)
    http.Error(w, errMsg, errCode)
}
