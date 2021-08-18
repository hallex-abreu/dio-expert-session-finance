package http

import (
	"net/http"

	"github.com/hallex-abreu/dio-expert-session-finance/adapter/http/actuator"
	"github.com/hallex-abreu/dio-expert-session-finance/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Init Routes
func Init() {

	http.HandleFunc("/transaction", transaction.GetTransations)
	http.HandleFunc("/transaction/create", transaction.CreateTransations)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)

}
