package main

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

var predictionCounter = prometheus.NewCounter(prometheus.CounterOpts{
    Name: "predictions_total",
    Help: "Total model predictions",
})

func init() {
    prometheus.MustRegister(predictionCounter)
}

func main() {
    http.Handle("/metrics", promhttp.Handler())
    http.HandleFunc("/predict", func(w http.ResponseWriter, r *http.Request) {
        predictionCounter.Inc()
        w.Write([]byte("Prediction made"))
    })
    http.ListenAndServe(":8080", nil)
}