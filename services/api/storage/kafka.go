package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rafaelbreno/go-bot/api/internal"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// KafkaClient manages kafka connection
// and actions.
type KafkaClient struct {
	Ctx   *internal.Context
	P     *kafka.Producer
	topic *string
}

func newKafkaClient(ctx *internal.Context) *KafkaClient {
	client := &KafkaClient{
		Ctx: ctx,
	}

	client.
		createTopic().
		setProducer()

	return client
}

func (k *KafkaClient) createTopic() *KafkaClient {
	t := new(string)
	*t = k.Ctx.Env["KAFKA_TOPIC"]
	k.topic = t

	kafkaAdmin, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.server": fmt.Sprintf("%s:%s", k.Ctx.Env["KAFKA_URL"], k.Ctx.Env["KAFKA_PORT"]),
	})

	defer kafkaAdmin.Close()

	if err != nil {
		k.Ctx.Logger.Error(err.Error())
		return k
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	results, err := kafkaAdmin.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{{
			Topic:             k.Ctx.Env["KAFKA_TOPIC"],
			NumPartitions:     2,
			ReplicationFactor: 2,
		}},
		kafka.SetAdminOperationTimeout(60*time.Second),
	)

	if err != nil {
		k.Ctx.Logger.Error(err.Error())
		return k
	}

	for _, res := range results {
		k.Ctx.Logger.Info(res.String())
	}

	return k
}

func (k *KafkaClient) setProducer() *KafkaClient {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":    fmt.Sprintf("%s:%s", k.Ctx.Env["KAFKA_URL"], k.Ctx.Env["KAFKA_PORT"]),
		"group.id":             0,
		"enable.partition.eof": true,
		"group.instance.id":    "1",
		"log_level":            2,
		"debug":                "all",
	})

	k.Ctx.Logger.Info("Connecting into Kafka...")

	if err != nil {
		k.Ctx.Logger.Error(err.Error())
	}

	if err := p.GetFatalError(); err != nil {
		k.Ctx.Logger.Error(err.Error())
		os.Exit(0)
	}

	return k
}

func (k *KafkaClient) Produce(key, value []byte) {
	deliveryChan := make(chan kafka.Event, 3)

	if err := k.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     k.topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
		Key:   key,
	}, deliveryChan); err != nil {
		if err != nil {
			k.Ctx.Logger.Error(err.Error())
			os.Exit(0)
		}
	}
}
