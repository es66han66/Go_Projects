package producer

import (
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func Configure(kafkaBrokerUrls []string, clientId string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientId,
	}

	config := kafka.WriterConfig{
		Brokers:      kafkaBrokerUrls,
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		Dialer:       dialer,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	w = kafka.NewWriter(config)
	fmt.Println("writer connected")
	writer = w
	return w, nil
}
