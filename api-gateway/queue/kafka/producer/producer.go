package producer

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// type KafkaProducer interface {
// 	ProducerMassages(topic string, handler func(massage []byte)) error
// 	Close() error
// }

type KafkaProducer struct {
	wrider *kafka.Writer
}

func NewKafkaProducerInit(broker []string) (KafkaProducer, error) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(broker...),
		AllowAutoTopicCreation: true,
	}

	return KafkaProducer{wrider: writer}, nil
}

func (k *KafkaProducer) ProducerMessage(topic string, massage []byte) error {
	return k.wrider.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Value: massage,
	})
}

func (k *KafkaProducer) Close() error {
	return k.wrider.Close()
}
