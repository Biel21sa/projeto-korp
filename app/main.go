package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requisições recebidas",
	},
)

func init() {
	prometheus.MustRegister(requestCounter)
}

func projetoKorp(w http.ResponseWriter, r *http.Request) {
	requestCounter.Inc()

	response := map[string]string{
		"nome":    "Projeto Korp",
		"horario": time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UP"))
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorp)
	http.HandleFunc("/health", health)
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}