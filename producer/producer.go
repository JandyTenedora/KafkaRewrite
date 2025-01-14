package producer

import "kafka_rewrite/utils"

type Producer struct {
	Conf *Config
	Conn *Connection
}

func (producer *Producer) WriteMessage(message []byte) error {
	config := producer.Conf
	topic32 := utils.HashStringToInt32(config.Topic)
	return producer.Conn.WriteMessage(topic32, message)
}
