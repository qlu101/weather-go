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

const ApiKey string = "28896f6afb4e54d741e4d57dbe5d2749"

type WeatherInfo struct {
    Description string              `json:"description,omitempty"`
    Note        string              `json:"note,omitempty"`
    DayCount    int                 `json:"daycount,omitempty"`
    Days        []darksky.DataPoint `json:"days,omitempty"`
}

func init() {
    log.Print("Listening on http requests...")
    http.HandleFunc("/weather", getHandler)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    //===== Get parameters
    latt, err := getParamFloat(w, r, "latitude")
    if err != nil {
        return
    }

    lngt, err := getParamFloat(w, r, "longitude")
    if err != nil {
        return
    }

    wi := WeatherInfo{
        Description: "Weather report for the past 7 days (1 week)",
        Note:        "Powered by Dark Sky (https://darksky.net/poweredby)",
    }

    //===== Retrieve past 7 day's weather data
    for i := -7; i <= -1; i++ {
        utcNow := time.Now().AddDate(0, 0, i).UTC().Unix()
        fm := "https://api.darksky.net/forecast/%s/%f,%f,%d?exclude=currently,hourly,minutely,alerts,flags"
        url := fmt.Sprintf(fm, ApiKey, latt, lngt, utcNow)
        // log.Print(url)

        ctx := appengine.NewContext(r)
        client := urlfetch.Client(ctx)
        resp, err := client.Get(url)
        if err != nil {
            errorOut(w, http.StatusFailedDependency, err.Error())
            return
        }
        defer resp.Body.Close()

        response := darksky.ForecastResponse{}
        err = json.NewDecoder(resp.Body).Decode(&response)
        if err != nil {
            errorOut(w, http.StatusFailedDependency, err.Error())
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

func getParamFloat(w http.ResponseWriter, r *http.Request, key string) (float64, error) {
    var errMsg string
    params, ok := r.URL.Query()[key]
    if !ok || len(params) < 1 {
        errMsg = fmt.Sprintf("Url Param '%s' is missing", key)
        errorOut(w, http.StatusBadRequest, errMsg)
        return 0.0, errors.New(errMsg)
    }

    param, err := strconv.ParseFloat(string(params[0]), 64)
    if err != nil {
        errorOut(w, http.StatusBadRequest, err.Error())
        return 0.0, err
    }

    return param, nil
}

func errorOut(w http.ResponseWriter, errCode int, errMsg string) {
    log.Print(errMsg)
    http.Error(w, errMsg, errCode)
}
