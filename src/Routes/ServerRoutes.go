package routes

import (
	"net/http"

	services "github.com/ignaciochemes/beammp-server-info-api/src/Services"
)

func GetDataFromBMPBackendRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	services.GetDataFromBMPBackendService(w, r)
}

func GetNewVersionOfServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	services.GetNewVersionOfServer()
}

func SendBeamMPServerFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=BeamMP_Server.zip")
	services.SendBeamMPServerFilesService(w, r)
}
