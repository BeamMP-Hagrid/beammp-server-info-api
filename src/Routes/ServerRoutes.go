package routes

import (
	"net/http"

	server_services "github.com/ignaciochemes/beammp-server-info-api/src/Services"
)

func GetDataFromBMPBackendRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	server_services.GetDataFromBMPBackendService(w, r)
}
