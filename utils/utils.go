package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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


func FormResponse(w *http.ResponseWriter, status int, err error){
	(*w).WriteHeader(status)
	(*w).Header().Add("Content-Type","application/json")

}