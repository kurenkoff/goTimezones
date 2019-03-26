package models

import (
	"fmt"
	"timezones/utils"
)

// Request структура - тело запроса
type Request struct {
	ID        int      `json:"id"`
	Timezones []string `json:"timezones"`
}

// Response структура - тело ответа
type Response struct {
	TimeInZones map[string]string `json:"timezones"`
}

// GetTime записывает в r.TimeInZones время в зонах, которые находятся в Request.Timezones
func (r *Response) GetTime(request Request) {
	// переделать в горутины
	fmt.Println(request.Timezones)
	r.TimeInZones = make(map[string]string)
	for _, zone := range request.Timezones {
		r.TimeInZones[zone], _ = utils.GetTime(zone)
	}
}

// GetTime записывает в r.TimeInZones время в зонах, которые находятся в Request.Timezones. Реализовано с помощью горутин
func (r *Response) GetTimeP(request Request) {
	r.TimeInZones = make(map[string]string)
	time := make(chan utils.PTime)

	for _, zone := range request.Timezones {
		go utils.GetTimeP(zone, time)
	}
	for i := 0; i < len(request.Timezones); i++ {
		rsp := <-time
		r.TimeInZones[rsp.Zone] = rsp.Time
	}
}
