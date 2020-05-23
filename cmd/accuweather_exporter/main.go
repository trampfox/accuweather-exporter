package cmd

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/trampfox/accuweather-exporter/collector"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func Execute() {
	flag.Parse()

	// Create a new instance of the AccuweatherCollector and
	// register it with the prometheus client
	accuweatherCollector := collector.NewAccuweatherCollector()
	prometheus.MustRegister(accuweatherCollector)

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Serving on port %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
