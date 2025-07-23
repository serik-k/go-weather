package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		return &GeoData{City: city}, nil
	}
	res, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("NOT200")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}