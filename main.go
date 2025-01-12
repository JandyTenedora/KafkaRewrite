package main

import (
	"kafka_rewrite/producer"
)

func NewProducer(config *producer.Config, conn *producer.Connection) *producer.Producer {
			return &producer.Producer{
							config:  config ,
							conn: conn}

}
