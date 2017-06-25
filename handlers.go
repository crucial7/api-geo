package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"googlemaps.github.io/maps"
)

type LatLongByAddressResponse struct {
	Address string
	Lat     float64
	Long    float64
}

type HTMLError struct {
	Err  string
	Code int
}

func LatLongByAddress(w http.ResponseWriter, req *http.Request) {

	client, err := maps.NewClient(maps.WithAPIKey("AIzaSyBKJ5hORqnuHg33bzcB7AyZBDr7jKzhwsc"))
	if err != nil {

		log.Fatalf("fatal error: %s", err)
	}

	r := &maps.GeocodingRequest{
		Address: req.URL.Path[len("/latlongByAddress/"):],
	}

	resp, err := client.Geocode(context.Background(), r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil && !strings.Contains(err.Error(), "ZERO_RESULTS") {

		log.Fatalf("fatal error: %s", err)
	} else {

		if err != nil && strings.Contains(err.Error(), "ZERO_RESULTS") {

			errResult := &HTMLError{
				Err:  err.Error(),
				Code: 500,
			}
			json.NewEncoder(w).Encode(errResult)
		} else {

			result := &LatLongByAddressResponse{
				Address: resp[0].FormattedAddress,
				Lat:     resp[0].Geometry.Location.Lat,
				Long:    resp[0].Geometry.Location.Lng,
			}
			json.NewEncoder(w).Encode(result)
		}
	}
}
