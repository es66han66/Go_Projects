package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

func main() {
	statsd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
	var i float64
	for {
		i += 1
		statsd.Gauge("example_metric.gauge", i, []string{"environment:dev"}, 1)
		statsd.Flush()
		fmt.Println("Sent ", i)
		time.Sleep(2 * time.Second)
	}
}
