package collector

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/trampfox/accuweather-exporter/accuweather"
)

// Struct which contains pointers to all the
// prometheus descriptor defined for this exporter
type AccuweatherCollector struct {
	client                   accuweather.AccuweatherClient
	locationKey              string
	temperature              *prometheus.Desc
	realFeelTemperature      *prometheus.Desc
	apparentTemperature      *prometheus.Desc
	humidity                 *prometheus.Desc
	pressure                 *prometheus.Desc
	windDirectionDegree      *prometheus.Desc
	precipitationPastHour    *prometheus.Desc
	precipitationPast3Hours  *prometheus.Desc
	precipitationPast6Hours  *prometheus.Desc
	precipitationPast12Hours *prometheus.Desc
	precipitationPast24Hours *prometheus.Desc
}

func NewAccuweatherCollector(apiKey string, locationKey string, location string) *AccuweatherCollector {
	client := accuweather.NewAccuweatherClient(apiKey)

	// Check if the location has been provided using the location flag. 
	// If so, we need to call the City Search API in order to retrieve the
	// location key that is required to call the Current Conditions API
	if location != "" {
		location, err := client.GetLocation(location)
		if err != nil {
			log.Printf("An error occurred while retrieving city ID, the provided one or the default one will be used")
			log.Printf("Error: %s", err)
		} else {
			log.Printf("Location retrieved from City Search API. Name: %s | Key: %s | Country: %s",
				location.LocalizedName, location.Key, location.Country.LocalizedName)
			locationKey = location.Key
		}
	}

	return &AccuweatherCollector{
		client:      client,
		locationKey: locationKey,
		temperature: prometheus.NewDesc("accuweather_temperature_value",
			"Current temperature (rounded value). May be NULL.", nil, nil),
		realFeelTemperature: prometheus.NewDesc("accuweather_realfeel_temperature_metric_value",
			"Patented AccuWeather RealFeel Temperature value.", nil, nil),
		apparentTemperature: prometheus.NewDesc("accuweather_apparent_temperature_metric_value",
			"Perceived outdoor temperature caused by the combination of air temperature, relative humidity, and wind speed.", nil, nil),
		humidity: prometheus.NewDesc("accuweather_humidity_value",
			"Relative humidity. May be NULL.", nil, nil),
		pressure: prometheus.NewDesc("accuweather_pressure_value",
			"Atmospheric pressure.", nil, nil),
		windDirectionDegree: prometheus.NewDesc("accuweather_wind_direction_degree_value",
			"Wind direction in Azimuth degrees (e.g. 180 degrees is a wind coming from the south). May be NULL.", nil, nil),
		precipitationPastHour: prometheus.NewDesc("accuweather_precipitation_past_hour_value",
			"The amount of precipitation (liquid equivalent) that has fallen in the past hour.", nil, nil),
		precipitationPast3Hours: prometheus.NewDesc("accuweather_precipitation_past_3_hours_value",
			"The amount of precipitation (liquid equivalent) that has fallen in the past three hours.", nil, nil),
		precipitationPast6Hours: prometheus.NewDesc("accuweather_precipitation_past_6_hours_value",
			"The amount of precipitation (liquid equivalent) that has fallen in thE past six hours.", nil, nil),
		precipitationPast12Hours: prometheus.NewDesc("accuweather_precipitation_past_12_hours_value",
			"The amount of precipitation (liquid equivalent) that has fallen in the past twelve hours.", nil, nil),
		precipitationPast24Hours: prometheus.NewDesc("accuweather_precipitation_past_24_hours_value",
			"The amount of precipitation (liquid equivalent) that has fallen in the past twenty-four hours.", nil, nil),
	}
}

// Each and every collector must implement the Describe function.
// It writes all the descriptors to the prometheus desc channel
func (collector *AccuweatherCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.temperature
	ch <- collector.realFeelTemperature
	ch <- collector.apparentTemperature
	ch <- collector.humidity
	ch <- collector.pressure
	ch <- collector.windDirectionDegree
	ch <- collector.precipitationPastHour
	ch <- collector.precipitationPast3Hours
	ch <- collector.precipitationPast6Hours
	ch <- collector.precipitationPast12Hours
	ch <- collector.precipitationPast24Hours
}

// Collect implements the required collect function for all
// Prometheus collector
func (collector *AccuweatherCollector) Collect(ch chan<- prometheus.Metric) {
	currentConditions, err := collector.client.GetCurrentConditions(collector.locationKey)
	if err != nil {
		log.Println(err)
	}

	// Write latest value for each metric in the prometheus metric channel
	for _, currentCondition := range *currentConditions {
		ch <- prometheus.MustNewConstMetric(collector.temperature, prometheus.GaugeValue, float64(currentCondition.Temperature.Metric.Value))
		ch <- prometheus.MustNewConstMetric(collector.realFeelTemperature, prometheus.GaugeValue, float64(currentCondition.RealFeelTemperature.Metric.Value))
		ch <- prometheus.MustNewConstMetric(collector.apparentTemperature, prometheus.GaugeValue, float64(currentCondition.ApparentTemperature.Metric.Value))
		ch <- prometheus.MustNewConstMetric(collector.humidity, prometheus.GaugeValue, float64(currentCondition.RelativeHumidity))
		ch <- prometheus.MustNewConstMetric(collector.pressure, prometheus.GaugeValue, float64(currentCondition.Pressure.Metric.Value))
		ch <- prometheus.MustNewConstMetric(collector.windDirectionDegree, prometheus.GaugeValue, float64(currentCondition.Wind.Direction.Degrees))
		ch <- prometheus.MustNewConstMetric(collector.precipitationPastHour, prometheus.GaugeValue, float64(currentCondition.PrecipitationSummary.PastHour.Metric.Value))
		ch <- prometheus.MustNewConstMetric(collector.precipitationPast3Hours, prometheus.GaugeValue, float64(currentCondition.PrecipitationSummary.Past3Hours.Metric.Value))
		ch <- prometheus.MustNewConstMetric(collector.precipitationPast6Hours, prometheus.GaugeValue, float64(currentCondition.PrecipitationSummary.Past6Hours.Metric.Value))
		ch <- prometheus.MustNewConstMetric(collector.precipitationPast12Hours, prometheus.GaugeValue, float64(currentCondition.PrecipitationSummary.Past12Hours.Metric.Value))
		ch <- prometheus.MustNewConstMetric(collector.precipitationPast24Hours, prometheus.GaugeValue, float64(currentCondition.PrecipitationSummary.Past24Hours.Metric.Value))
	}
}
