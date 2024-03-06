package kafak

import (
    "github.com/IBM/sarama"
    "github.com/oigi/Magikarp/config"
)

var GobalKafka sarama.Client

func InitKafka() {
    con := sarama.NewConfig()
    con.Producer.Return.Successes = true
    kafkaClient, err := sarama.NewClient(config.CONFIG.Kafka, con)
    if err != nil {
        // TODO日志
        return
    }
    GobalKafka = kafkaClient
}
