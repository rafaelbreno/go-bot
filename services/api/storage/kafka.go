package storage

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rafaelbreno/go-bot/api/internal"
)

// Prod is Kafka's producer
type Prod struct {
	Ctx *internal.Context
	P   *kafka.Producer
}

func newProd(ctx *internal.Context) *Prod {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":    "localhost:9092",
		"group.id":             0,
		"enable.partition.eof": true,
		"group.instance.id":    "1",
		"log_level":            7,
		"debug":                "all",
	})

	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return &Prod{
		Ctx: ctx,
		P:   p,
	}
}
