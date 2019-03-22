package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"timezones/db"
	"timezones/models"
	"timezones/schema"

	"github.com/gorilla/mux"
)

/*
Сервис реализует JSON API работающее по HTTP.
На вход принимает список зон, в ответе выдает список зон с текущим временем в них.
*/

var (
	TimeZones []string     // Список всех допустимых временных зон
	database  *db.Database // Подключение к БД
)

func main() {
	database.Initialize(TimeZones)
	defer database.Close()


	r := mux.NewRouter()
	r.HandleFunc("/time", GetTimeZone)
	var port string
	if os.Getenv("APP_PORT") == "" {
		port = ":8080"
	} else {
		port = ":" + os.Getenv("APP_PORT")
	}
	log.Fatal(http.ListenAndServe(port, r))
}

// init Получение списка временных зон
func init() {
	database = db.New()
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
		database.UpdateUserData(request.ID, request.Timezones)
		request.Timezones = database.GetTimezones(request.ID)
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
