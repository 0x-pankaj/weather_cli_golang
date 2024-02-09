package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const API_KEY = "d1a89e10ca024356b7"

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
			Hour []struct {
				Time      string  `json:"time"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				Rain int `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	// fmt.Println("weather cli")

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + API_KEY + "&q=Kathmandu")
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

	// fmt.Println(string(body))

	var weather Weather

	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic("error while Unmarshal json")
	}

	/*
		weatherJSON, err := json.MarshalIndent(weather, "", "\t")
		fmt.Println(string(weatherJSON))
	*/

	location, current, hours := weather.Location, weather.Current, weather.ForeCast.Forecastday[0].Hour

	// kathmandu, nepal : temp, condition
	fmt.Printf(" %s, %s, : %.0fC, %s  \n", location.Name, location.Country, current.Temp_c, current.Condition.Text)

	// fmt.Println(hours)
	for _, val := range hours {
		// t, err := time.Parse("2024-02-09 00:00", val.Time)
		// if err != nil {
		// 	panic("Error parsing time")
		// }
		// formattedTime := t.Format("00:00")
		// fmt.Println(val)
		fmt.Printf("Time: %s , Temp: %.0fC , Cond: %s , RainChance: %d \n", val.Time, val.TempC, val.Condition.Text, val.Rain)
	}

}
