package actuator

import (
	"encoding/json"
	"net/http"
)

//HealthBody json
type HealthBody struct {
	Status string `json: "status"`
}

//Health Return Status
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var status = HealthBody{Status: "alive"}

	_ = json.NewEncoder(w).Encode(status)
}
