Weather Web Service
===================

A sample web service which provides the past 7 days' weather information. It is written in golang, and deployed to Google Cloud. The source code utilizes some data structures from `github.com/shawntoffel/darksky`.

------------------

### Usage

You can request the past 7 days' weather information with the following URL.
```
https://weathergo2.appspot.com/weather?latitude=<LATITUDE>&longitude=<LONGITUDE>
```

For example, the latitude and longitude of Irvine city is 33.6846 and -117.8265 respectively. The URL is:

```
https://weathergo2.appspot.com/weather?latitude=33.6846&longitude=-117.8265
```

The service will send back the past 7 day's weather information in JSon format. Below illustrates an example.

```
{
    "description": "Weather report for the past 7 days (1 week)",
    "note": "Powered by Dark Sky (https://darksky.net/poweredby)",
    "daycount": 7,
    "days": [
        {
            "apparentTemperatureMax": 61.67,
            "apparentTemperatureMaxTime": 1516917600,
            "apparentTemperatureMin": 46.65,
            "apparentTemperatureMinTime": 1516885200,
            "cloudCover": 0.14,
            "dewPoint": 43.65,
            "humidity": 0.68,
            "icon": "partly-cloudy-day",
            "moonPhase": 0.29,
            "ozone": 303.07,
            "pressure": 1018,
            "summary": "Partly cloudy until afternoon.",
            "sunriseTime": 1516891995,
            "sunsetTime": 1516929384,
            "temperatureMax": 61.67,
            "temperatureMaxTime": 1516917600,
            "temperatureMin": 46.65,
            "temperatureMinTime": 1516885200,
            "time": 1516867200,
            "uvIndex": 2,
            "uvIndexTime": 1516906800,
            "visibility": 9.92,
            "windBearing": 204,
            "windGust": 9.33,
            "windGustTime": 1516921200,
            "windSpeed": 2.6
        },
        {
            "apparentTemperatureMax": 65.42,
            "apparentTemperatureMaxTime": 1517007600,
            "apparentTemperatureMin": 41.36,
            "apparentTemperatureMinTime": 1516975200,
            "dewPoint": 40.83,
            "humidity": 0.63,
            "icon": "clear-day",
            "moonPhase": 0.33,
            "ozone": 310.03,
            "pressure": 1021.61,
            "summary": "Clear throughout the day.",
            "sunriseTime": 1516978363,
            "sunsetTime": 1517015844,
            "temperatureMax": 65.42,
            "temperatureMaxTime": 1517007600,
            "temperatureMin": 43.83,
            "temperatureMinTime": 1516975200,
            "time": 1516953600,
            "uvIndex": 3,
            "uvIndexTime": 1516993200,
            "visibility": 9.87,
            "windBearing": 324,
            "windGust": 8.26,
            "windGustTime": 1516953600,
            "windSpeed": 0.69
        },
        {
            "apparentTemperatureMax": 75.72,
            "apparentTemperatureMaxTime": 1517097600,
            "apparentTemperatureMin": 42.89,
            "apparentTemperatureMinTime": 1517058000,
            "cloudCover": 0.03,
            "dewPoint": 32.2,
            "humidity": 0.43,
            "icon": "clear-day",
            "moonPhase": 0.36,
            "ozone": 299.33,
            "pressure": 1021.65,
            "summary": "Clear throughout the day.",
            "sunriseTime": 1517064730,
            "sunsetTime": 1517102303,
            "temperatureMax": 75.72,
            "temperatureMaxTime": 1517097600,
            "temperatureMin": 44.67,
            "temperatureMinTime": 1517065200,
            "time": 1517040000,
            "uvIndex": 4,
            "uvIndexTime": 1517083200,
            "visibility": 9.91,
            "windBearing": 334,
            "windGust": 6.69,
            "windGustTime": 1517094000,
            "windSpeed": 0.89
        },
        {
            "apparentTemperatureMax": 84.93,
            "apparentTemperatureMaxTime": 1517180400,
            "apparentTemperatureMin": 58.45,
            "apparentTemperatureMinTime": 1517137200,
            "cloudCover": 0.02,
            "dewPoint": 28.74,
            "humidity": 0.21,
            "icon": "clear-day",
            "moonPhase": 0.4,
            "ozone": 300.75,
            "pressure": 1020.15,
            "summary": "Clear throughout the day.",
            "sunriseTime": 1517151095,
            "sunsetTime": 1517188762,
            "temperatureMax": 84.93,
            "temperatureMaxTime": 1517180400,
            "temperatureMin": 58.45,
            "temperatureMinTime": 1517137200,
            "time": 1517126400,
            "uvIndex": 4,
            "uvIndexTime": 1517169600,
            "visibility": 9.72,
            "windBearing": 43,
            "windGust": 15.06,
            "windGustTime": 1517169600,
            "windSpeed": 5.55
        },
        {
            "apparentTemperatureMax": 85.38,
            "apparentTemperatureMaxTime": 1517266800,
            "apparentTemperatureMin": 67.44,
            "apparentTemperatureMinTime": 1517295600,
            "cloudCover": 0.11,
            "dewPoint": 27.74,
            "humidity": 0.18,
            "icon": "clear-day",
            "moonPhase": 0.44,
            "ozone": 299.4,
            "pressure": 1018.26,
            "summary": "Clear throughout the day.",
            "sunriseTime": 1517237459,
            "sunsetTime": 1517275222,
            "temperatureMax": 85.38,
            "temperatureMaxTime": 1517266800,
            "temperatureMin": 67.44,
            "temperatureMinTime": 1517295600,
            "time": 1517212800,
            "uvIndex": 3,
            "uvIndexTime": 1517252400,
            "visibility": 9.72,
            "windBearing": 43,
            "windGust": 11.47,
            "windGustTime": 1517234400,
            "windSpeed": 3.38
        },
        {
            "apparentTemperatureMax": 80.16,
            "apparentTemperatureMaxTime": 1517349600,
            "apparentTemperatureMin": 58.24,
            "apparentTemperatureMinTime": 1517382000,
            "cloudCover": 0.17,
            "dewPoint": 31.72,
            "humidity": 0.27,
            "icon": "partly-cloudy-day",
            "moonPhase": 0.48,
            "ozone": 296.28,
            "precipIntensityMax": 0.001,
            "precipIntensityMaxTime": 1517320800,
            "pressure": 1016.23,
            "summary": "Partly cloudy in the morning.",
            "sunriseTime": 1517323821,
            "sunsetTime": 1517361681,
            "temperatureMax": 80.16,
            "temperatureMaxTime": 1517349600,
            "temperatureMin": 58.24,
            "temperatureMinTime": 1517382000,
            "time": 1517299200,
            "uvIndex": 3,
            "uvIndexTime": 1517338800,
            "visibility": 10,
            "windBearing": 349,
            "windGust": 6.32,
            "windGustTime": 1517302800,
            "windSpeed": 1.19
        },
        {
            "apparentTemperatureMax": 76.51,
            "apparentTemperatureMaxTime": 1517436000,
            "apparentTemperatureMin": 53.57,
            "apparentTemperatureMinTime": 1517400000,
            "cloudCover": 0.03,
            "dewPoint": 35.34,
            "humidity": 0.36,
            "icon": "clear-day",
            "moonPhase": 0.51,
            "ozone": 296.28,
            "precipIntensityMax": 0.0002,
            "precipIntensityMaxTime": 1517410800,
            "pressure": 1014.16,
            "summary": "Clear throughout the day.",
            "sunriseTime": 1517410181,
            "sunsetTime": 1517448140,
            "temperatureMax": 76.51,
            "temperatureMaxTime": 1517436000,
            "temperatureMin": 53.57,
            "temperatureMinTime": 1517400000,
            "time": 1517385600,
            "uvIndex": 4,
            "uvIndexTime": 1517428800,
            "visibility": 10,
            "windBearing": 111,
            "windGust": 6.62,
            "windGustTime": 1517428800,
            "windSpeed": 0.46
        }
    ]
}
```
