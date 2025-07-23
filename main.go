package main

import (
	"flag"
	"fmt"
	"myproject/geo"
	"myproject/weather"
)


func main() {
    fmt.Println("Hello, World!")
    city := flag.String("city", "Алматы", "Город пользователя")
    format := flag.Int("format", 1, "Формат вывода погоды")
    flag.Parse()
    geoData, err := geo.GetMyLocation(*city)
    if err != nil {
        fmt.Println("Ошибка геолокации:", err)
        return
    }
    if geoData == nil {
        fmt.Println("Геоданные не получены")
        return
    }
    fmt.Println(geoData)
    weatherData := weather.GetWeather(*geoData, *format)
    fmt.Println(weatherData)
}

