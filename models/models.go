package models

import (
	"fmt"
	"timezones/utils"
)

// Запрос
type Request struct{
	ID int `json:"id"`
	Timezones []string `json:"timezones"`
}



// Ответ
type Response struct{
	TimeInZones map[string]string `json:"timezones"` 
}

func (r *Response) GetTime(request Request){
	// переделать в горутины
	fmt.Println(request.Timezones)
	r.TimeInZones = make(map[string]string)
	for _, zone := range request.Timezones {
		r.TimeInZones[zone], _ = utils.GetTime(zone)
	}
}



func (r *Response) GetTimeP(request Request){
	r.TimeInZones = make(map[string]string)
	time := make(chan utils.PTime)

	for _, zone := range request.Timezones  {
		go utils.GetTimeP(zone, time)
	}
	for i := 0; i < len(request.Timezones); i++{
		rsp := <-time
		r.TimeInZones[rsp.Zone] = rsp.Time
	}
}
