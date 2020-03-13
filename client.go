package openweatherapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	)

type City struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Weather struct {
	Id int `json:"id"`
	Main string `json:"main"`
	Description string `json:"description"`
}

type Clouds struct {
	All int `json:"all"`
}

type Main struct {
	Temp float64 `json:"temp"`
	Pressure int `json:"pressure"`
	Humidity int `json:"humidity"`
	Temp_min float64 `json:"temp_min"`
	Temp_max float64 `json:"temp_max"`
}

type CurrentWeatherResponse struct {
	Weather []Weather `json:"weather"`
	Main `json:"main"`
	Clouds `json:"clouds"`
	Dt int `json:"dt"`
	Id int `json:"id"`
	Name string `json:"name"`
	Cod int `json:"cod"`
	Kelembapan string
}

func GetWeather(kota string, api_key string) CurrentWeatherResponse{
	var cwr CurrentWeatherResponse
	url := "http://api.openweathermap.org/data/2.5/weather?q="+kota+"&appid="+api_key
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err!=nil{
		   log.Fatal(err)
		}
		jsonErr := json.Unmarshal(body, &cwr)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		cwr.Temp=convertToCelsius(cwr.Temp)
		cwr.Temp_min=convertToCelsius(cwr.Temp_min)
		cwr.Temp_max=convertToCelsius(cwr.Temp_max)
		cwr.Kelembapan=humidityRange(cwr.Humidity)
		return cwr
	}
	return cwr
}