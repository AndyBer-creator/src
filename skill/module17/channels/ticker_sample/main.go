package main

import (
	"fmt"
	"time"
)

func main() {
	var ticker *time.Ticker = time.NewTicker(time.Second * 1)
	var t time.Time
	for {
		t = <-ticker.C
		outputMessage := []byte("Время: ")
		// Метод AppendFormat преобразует объект time.Time
		// к заданному строковому формату (второй аргумент)
		// и добавляет полученную строку к строке, переданной в первом
		// аргументе
		outputMessage = t.AppendFormat(outputMessage, "15:04:05")
		fmt.Println(string(outputMessage))
	}
}
