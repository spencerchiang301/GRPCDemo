package utility

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const apiURL = "https://maps.googleapis.com/maps/api/geocode/json"

type API struct {
	googleKey string
}

type GeocodingResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
	Status string `json:"status"`
}

func GetLatLongFromAddress(address string) (float64, float64, error) {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("Can't read api key!!!")
	}

	apiKey := cfg.Section("GoogleAPI").Key("Key").String()
	encodedAddress := url.QueryEscape(address)

	// Build the request URL with the encoded address and API key
	requestURL := fmt.Sprintf("%s?address=%s&key=%s", apiURL, encodedAddress, apiKey)

	// Make the HTTP request
	resp, err := http.Get(requestURL)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	// Parse the JSON response
	var geocodeResponse GeocodingResponse
	err = json.Unmarshal(body, &geocodeResponse)
	if err != nil {
		return 0, 0, err
	}

	// Check if the API returned results
	if geocodeResponse.Status != "OK" {
		return 0, 0, fmt.Errorf("geocoding failed with status: %s", geocodeResponse.Status)
	}

	// Extract the latitude and longitude from the response
	lat := geocodeResponse.Results[0].Geometry.Location.Lat
	lng := geocodeResponse.Results[0].Geometry.Location.Lng

	return lat, lng, nil
}
