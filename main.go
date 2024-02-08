package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const API_KEY = "d1a89e10ca02435"

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
	ForeCast struct {
		Forecastday []struct {
			Day struct {
				Maxtemp string `json:"maxtemp_c"`
				Mintemp string `json:"mintemp_c"`
				RainP   string `json:"daily_chance_of_rain"`
			} `json:"day"`
			Astro struct {
				Sunrise   string `json:"sunrise"`
				Sunset    string `json:"sunset"`
				MoonPhase string `json:"moon_phase"`
			} `json:"astro"`
			Hour []struct {
				Time      string `json:"time"`
				TempC     string `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				WindD string `json:"wind_dir"`
				Rain  int    `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	fmt.Println("weather cli")

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + API_KEY + "&q=Kathmandu")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// resp, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + API_KEY + "&q=Kathmandu")
	// if err != nil {
	// 	panic(err)
	// }

	// b, err := io.ReadAll(resp.Body)
	// fmt.Println(string(b))

	if res.StatusCode != 200 {
		panic("Weather Api not Available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))

	var weather Weather

	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic("error while Unmarshal json")
	}

	fmt.Println(weather)
	// city, time, temp, cond := weather.Location.Name, weather.Location.LocalTime, weather.Current.Temp_c, weather.Current.Condition.Text

	// fmt.Printf("%s, %s, %.0fC, %s, \n ", city, time, temp, cond)

}
