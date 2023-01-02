package server_services

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	structs "github.com/ignaciochemes/beammp-server-info-api/src/Structs"
)

var uri = "https://backend.beammp.com/servers-info"

func GetDataFromBMPBackendService(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if !query.Has("serverIp") {
		newError := structs.Error{
			Error:   "No serverIp query param.",
			Message: "You need to send a serverIp query param.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(newError)
		return
	}
	response, err := http.Get(uri)
	if err != nil {
		newError := structs.Error{
			Error:   "Error getting data from beam mp backend",
			Message: "Error getting data from beam mp backend",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(newError)
		panic(err)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		newError := structs.Error{
			Error:   "Error reading response body.",
			Message: "Error reading response body.",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(newError)
		panic(err)
	}
	var servers []structs.Server
	json.Unmarshal([]byte(responseData), &servers)

	var newServers []structs.Server

	for _, server := range servers {
		if server.IP == query.Get("serverIp") {
			modTotalSize, err := strconv.ParseInt(server.ModsTotalSize, 10, 64)
			if err != nil {
				return
			}
			parseToGb := strconv.FormatFloat(float64(modTotalSize)/1000000000, 'f', 2, 64)
			server.ModsTotalSize = parseToGb + " GB"
			newServers = append(newServers, server)
		}
	}

	if len(newServers) == 0 {
		newError := structs.Error{
			Error: "No servers found",
			Message: "No servers found with the serverIp query param. " +
				"Please check the serverIp query param.",
		}
		json.NewEncoder(w).Encode(newError)
		return
	}
	json.NewEncoder(w).Encode(newServers)
}
