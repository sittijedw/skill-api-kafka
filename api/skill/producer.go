package skill

import (
	"log"
	"os"
	"strings"

	_ "net/http/pprof"

	"github.com/IBM/sarama"
)

// Sarama configuration options
var (
	brokerURLS = os.Getenv("BROKER_URLS")
	topic      = os.Getenv("TOPIC")
)

func init() {
	if len(topic) == 0 {
		panic("no topic given to be consumed, please set the -topic flag")
	}
}

func sendMessage(message, action string) {
	config := configProducer()

	producer, err := sarama.NewSyncProducer(strings.Split(brokerURLS, ","), config)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(action), Value: sarama.StringEncoder(message)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}

func configProducer() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.NoResponse

	return config
}
