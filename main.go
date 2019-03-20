package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"timezones/models"
	"timezones/schema"

	"github.com/gorilla/mux"
)

/*
Сервис реализует JSON API работающее по HTTP.
На вход принимает список зон, в ответе выдает список зон с текущим временем в них.

   Заметки:
	1. 2.03.19 Решено использовать API http://worldtimeapi.org/
	2. 4.03.19 Для реализации JsonSchema выбрана библиотека github.com/xeipuuv/gojsonschema

	TO-DO
	1. Обработка ошибок (нормальная)
	2. GET запрос
	3. Main: mux/router -> стандартная библиотека
	4. Разбиение на файлы
	5. БД для хранения пользователей

*/

// Список всех допустимых временных зон
var TimeZones []string


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/time", GetTimeZone)
	var port string
	if os.Getenv("APP_PORT") == ""{
		port = ":8080"
	} else {
		port = ":" + os.Getenv("APP_PORT")
	}
	log.Fatal(http.ListenAndServe(port, r))
}

// Инициализация(выполняется автоматически в самом начале работы программы).
// Получение списка временных зон
func init() {
	// Очень некрасивый код. Поправить
	r, err := http.Get("http://worldtimeapi.org/api/timezone")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	err = json.Unmarshal(jsn, &TimeZones)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

// Плохая реакция на пустое тело запроса

// Handler.
// Body - models.Request
// Response - models.Response
func GetTimeZone(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// т.к. в теле запроса массив находится массив таймзон
		// Необходимо достать его
		jsn, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("GetTimeZone: Ошибка при попытке прочитать тело запроса: %v\n", err)
		}
		// Парсинг тела запроса
		var request models.Request
		err = json.Unmarshal(jsn, &request)
		if err != nil {
			log.Printf("GetTimeZone: Ошибка при попытке парсинга тела: %v\n", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		// Проверка запроса по JsonSchema
		validator := schema.NewValidator()
		err = validator.ValidateRequest(string(jsn))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "%s", err.Error())
			return
		}

		// Здесь должна быть  логика хранилища
		// Конец логики хранилища

		// Теперь будем запрашивать время для каждой из зон
		// Или можно для всех сразу? (проверить на сайте, который предлагает API)
		var response models.Response

		response.GetTimeP(request)

		jsn, err = json.Marshal(response)

		if err != nil {
			log.Printf("GetTimeZone: Ошибка при попытке замаршалить ответ: %v\n", err)
		}
		err = validator.ValidateResponse(string(jsn))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "%s", err.Error())
			return
		}
		fmt.Fprintf(w, "%s", string(jsn))

	}
}
