package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type apiResponse struct {
	Datetime string `json:"datetime"`
	Unixtime string `json:"unixtime"`
}

func GetTime(zone string) (string, error) {
	URL := "http://worldtimeapi.org/api/timezone/" + zone
	// Запрашиваем данные
	r, _ := http.Get(URL)
	if r.StatusCode == http.StatusOK {
		// Разбираем ответ
		jsn, _ := ioutil.ReadAll(r.Body)
		var response apiResponse
		err := json.Unmarshal(jsn, &response)
		if err != nil{
			return "", err
		}
		return response.Datetime, fmt.Errorf("")
	}
	return "", fmt.Errorf("bad status code")
}

type PTime struct {
	Zone string
	Time string
}

func GetTimeP(zone string, result chan PTime) {
	URL := "http://worldtimeapi.org/api/timezone/" + zone
	// Запрашиваем данные
	r, _ := http.Get(URL)
	if r.StatusCode == http.StatusOK {
		// Разбираем ответ
		jsn, _ := ioutil.ReadAll(r.Body)
		var response apiResponse
		err := json.Unmarshal(jsn, &response)
		if err != nil{
			log.Println(err)
		}
		result <- PTime{ zone, response.Datetime}
		return
	}
	result <- PTime{ zone, ""}
	return
}