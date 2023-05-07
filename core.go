package main

import (
	"github.com/goodsign/monday"
	"time"
)

func CurrentTime() string {
	loc, _ := time.LoadLocation("America/Havana")
	time.Local = loc
	theTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), loc)
	date := theTime.Format("2006-01-02 03:04:05 pm")
	return date
}

func FormatDate(input string) string {
	loc, _ := time.LoadLocation("America/Havana")
	time.Local = loc
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	return monday.Format(t, "02 de January del 2006", monday.LocaleEsES)
}
