package cmd

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/trampfox/accuweather-exporter/collector"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app    = kingpin.New("accuweather-exporter", "Prometheus Exporter for AccuWeather API").Author("Davide Monfrecola")
	addr   = app.Flag("listen-address", "The address to listen on for HTTP requests.").Envar("AE_LISTEN_ADDRESS").Default(":9095").String()
	apiKey = app.Flag("api-key", "The API key for Accuweather API requests.").Envar("AE_API_KEY").Required().String()
	// TODO add language and unit (metric or imperial)
)

func Execute() {
	// Create a new instance of the AccuweatherCollector and
	// register it with the prometheus client
	accuweatherCollector := collector.NewAccuweatherCollector(*apiKey)
	prometheus.MustRegister(accuweatherCollector)

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Serving on port %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
