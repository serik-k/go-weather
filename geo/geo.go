package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type  CityResponse struct {
	Error bool	 `json:"error"`
}


func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := CheckCity(city)
		if !isCity {
			panic("Город не найден: " + city)
		}
		return &GeoData{City: city}, nil
	}
	res, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("NOT200")
	}
	defer res.Body.Close()	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

var CheckCityURL = "https://countriesnow.space/api/v0.1/countries/population/cities"


func CheckCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	res, err := http.Post(CheckCityURL, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false
	}

	var cityResponse CityResponse
	json.Unmarshal(body, &cityResponse)
	return !cityResponse.Error 
}