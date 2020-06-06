# accuweather-exporter

Prometheus exporter for AccuWeather [Current Conditions](https://developer.accuweather.com/accuweather-current-conditions-api/apis) API to gather weather metrics for a specific location.

## Requirements

You have to [signup here](https://developer.accuweather.com/user/register) for an AccuWeather APIs account and then create a new App to get the API key needed by this exporter.

Accuweather exporter uses the following APIs

* [Current Conditions API](https://developer.accuweather.com/accuweather-current-conditions-api/apis)
* [Locations API](https://developer.accuweather.com/accuweather-locations-api/apis)

## Configuration

Accuweather exporter is setup to take run configuration from environment variables or CLI flags, which are listed in the table below.

| Environment variable | Flag             | Description                                                  | Default |
| :------------------- | ---------------- | ------------------------------------------------------------ | ------- |
| AE_LISTEN_ADDRESS    | --listen-address | The address to listen on for HTTP requests                   | :9095   |
| AE_API_KEY           | --api-key        | The API key for Accuweather API requests                     | -       |
| AE_LOCATION          | --location       | The location for which you want to retrieve current conditions data (e.g. `Turin, IT`) | -       |
| AE_LOCATION_KEY      | --location-key   | The location key of the city for which you want to retrieve current conditions data | 214753  |

### Location configuration

The location can be specified in two ways:

* providing a location key using the `location-key` flag or the `AE_LOCATION_KEY` environment variable

* providing a location string, using the `location` or the `AE_LOCATION` environment variable

In the latter way the exporter calls the [City search API](https://developer.accuweather.com/accuweather-locations-api/apis/get/locations/v1/cities/search) and it retrieves the location key from the **first result**.

## Usage

### Binary usage

Export the weather metrics for the city of Turin (Italy) using the binary and `location-key` flag.

```bash
./accuweather-exporter --api-key <apikey> --location-key 214753
```

Export the weather metrics for the city of Turin (Italy) using the binary and `location` flags (the exporter will call the City search API to retrieve the Turin's key).

```bash
./accuweather-exporter --api-key <apikey> --location "Turin, IT"
```
