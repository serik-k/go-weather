package geo_test

import (
	"encoding/json"
	"myproject/geo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMyLocation_WithValidCity(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]bool{"error": false}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	originalURL := geo.CheckCityURL
	geo.CheckCityURL = server.URL
	defer func() { geo.CheckCityURL = originalURL }()

	result, err := geo.GetMyLocation("Almaty")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.City != "Almaty" {
		t.Errorf("expected city to be 'Almaty', got %s", result.City)
	}
}

func TestGetMyLocation_WithInvalidCity_ShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for invalid city, but got none")
		}
	}()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]bool{"error": true}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	originalURL := geo.CheckCityURL
	geo.CheckCityURL = server.URL
	defer func() { geo.CheckCityURL = originalURL }()

	geo.GetMyLocation("InvalidCity")
}
