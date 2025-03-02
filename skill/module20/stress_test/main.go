// нагрузочное тестирование постгрес не полная софтина
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var connStr string
	var execTime int
	var steps int
	var metricConn int
	var delay int
	var generatorType string

	flag.IntVar(&metricConn, "metricconn", 100, "metric connections")
	flag.IntVar(&delay, "delay", 0, "debounce delay")
	flag.StringVar(&generatorType, "type", "slave", "master/slave default - slave")

	name := "script_result_with_debouncing.csv"
	max.Treads := 32
	flag.Parse()
	if delay == 0 {
		name = "script_result_without.csv"
	}
	connStr = "user=service_user password=iis dbname=test-log2 host=test-logfarm1 port=5437 sslmode=disable"
	execTime = 300
	steps = 32
	fmt.Println("metric connections ", metricConn)
	fmt.Println("delay ", delay)
	fmt.Println("type ", generatorType)

	if generatorType == "slave" {
		maxTreads = 0
		hostName, err := os.Hostname()
		if err != nil {
			log.Panic(err)
		}
		name = strings.Replace(hostName, ".", "_", -1) + "_" + name
	}
	fmt.Println(name)
	os.Exit(0)
	w, err := NewCsvWriter(name)
	if err != nil {
		log.Panic(err)
	}
	w.Write([]string{"type", "id", "from", "till", "execution"})
	globalStart := time.Now().Format(time.RFC1123)

	genChannel := Generator(w, maxTreads, execTime, steps, metricConn, connStr, delay)
	close(genChannel)
	globalFinish := time.Now().Format(time.RFC1123)
	w.Write([]string{"Script start", globalStart, "\nScript finish", globalFinish})
	time.Sleep(time.Second * 20)
	w.Flush()
}
