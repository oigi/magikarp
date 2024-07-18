package kafak

import (
	"github.com/IBM/sarama"
	"github.com/oigi/Magikarp/config"
)

var GlobalKafka sarama.Client

// InitKafka TODO 配置Kafka
func InitKafka() {
	con := sarama.NewConfig()
	con.Producer.Return.Successes = true
	kafkaClient, err := sarama.NewClient(config.CONFIG.Kafka.Address, con)
	if err != nil {
		// TODO日志
		return
	}
	GlobalKafka = kafkaClient
}
