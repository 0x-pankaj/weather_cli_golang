package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const API_KEY = "d1a89e10caxxx356b7b11062424xxxx"

type Weather struct {
	Location struct {
		Name      string `json:"name"`
		Country   string `json:"country"`
		LocalTime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		Last_Updated string  `json:"last_updated"`
		Temp_c       float32 `json:"temp_c"`
		Condition    struct {
			Text string `json:"text"`
		} `json:"condition"`
		Wind_kph float32 `json:"wind_kph"`
		Humidity int     `json:"humidity"`
	} `json:"current"`
}

func main() {
	fmt.Println("weather cli")

	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + API_KEY + "&q=London")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather Api not Available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather

	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	// fmt.Println(weather)
	city, time, temp, cond := weather.Location.Name, weather.Location.LocalTime, weather.Current.Temp_c, weather.Current.Condition.Text

	fmt.Printf("%s, %s, %.0fC, %s, ", city, time, temp, cond)

}
