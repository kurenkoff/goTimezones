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

func GetTime(zone string) string {
	URL := "http://worldtimeapi.org/api/timezone/" + zone
	// Запрашиваем данные
	r, err := http.Get(URL)
	if r.StatusCode == http.StatusOK {
		if err != nil {
			fmt.Printf("getTime: Ошибка при попытке получить время в таймзоне: %v\n", err)
		}
		// Разбираем ответ
		jsn, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("getTime: Ошибка при попытке прочитать тело запроса: %v\n", err)
		}
		var response apiResponse
		err = json.Unmarshal(jsn, &response)
		if err != nil {
			fmt.Printf("getTime: Ошибка при попытке парсинга тела: %v\n", err)
		}
		return response.Datetime
	}
	return "ErrorOccurred"
}


func FormResponse(w *http.ResponseWriter, status int, err error){
	(*w).WriteHeader(status)
	(*w).Header().Add("Content-Type","application/json")

}