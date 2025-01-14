package main

import "kafka_rewrite/producer"

func NewProducer(config *producer.Config, conn *producer.Connection) *producer.Producer {
	return &producer.Producer{
		Conf: config,
		Conn: conn,
	}
}

func main() {
	config := producer.Config{BrokerAddress: "localhost:9091", Topic: "First Topic"}
	conn, _ := producer.NewConnection(config.BrokerAddress)
	producer := NewProducer(&config, conn)
	producer.WriteMessage("This is a test message")
}
