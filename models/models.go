package models

// Запрос
type Request struct{
	Timezones []string `json:"timezones"`
}



// Ответ
type Response struct{
	TimeInZones map[string]string `json:"timezones"` 
}