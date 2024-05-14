package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/oriiyx/arso/queries"
)

func HandleMaxStrength(w http.ResponseWriter, r *http.Request) error {
	// Retrieve latitude and longitude from the query string
	latStr := r.URL.Query().Get("lat")
	lonStr := r.URL.Query().Get("lon")

	// Check if latitude and longitude are not empty
	if latStr == "" || lonStr == "" {
		http.Error(w, "Latitude and longitude are required", http.StatusBadRequest)
		return nil
	}

	// Try to parse latitude and longitude as floats
	lat, errLat := strconv.ParseFloat(latStr, 64)
	lon, errLon := strconv.ParseFloat(lonStr, 64)
	if errLat != nil || errLon != nil {
		http.Error(w, "Latitude and longitude must be valid floating-point numbers", http.StatusBadRequest)
		return nil
	}

	// If everything is okay, proceed with the rest of the handler logic
	// (Implementation depends on further requirements)
	// @todo - Implement the rest of the handler logic
	ms := queries.MaxStrength{
		Lat: lat,
		Lon: lon,
	}

	fmt.Println(lat)
	fmt.Println(lon)

	// Execute the query
	msData, err := ms.Execute(ms)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return nil
	}

	if len(msData) == 0 {
		msData = append(msData, queries.MaxStrengthData{
			ID:       0,
			Name:     "No data found",
			Area:     0,
			Strength: 0,
		})
	}

	msSingleData := msData[0]
	msJson, err := json.Marshal(msSingleData)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(msJson)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return err
	}

	return nil
}
