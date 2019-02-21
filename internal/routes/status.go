package routes

import (
	"encoding/json"
	"net/http"

	"github.com/MSLacerda/status-hub/internal/services"
)

type DataJson struct {
	Data []NormalizedStatus
}

func makeRequest(el services.Service) NormalizedStatus {
	return HandlerRequest(el)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	var endpoints []services.Service = *services.GetServices()
	// var arr []string
	var response DataJson

	for _, element := range endpoints {
		status := makeRequest(element)
		response.Data = append(response.Data, status)
	}

	parsed, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(parsed)
}
