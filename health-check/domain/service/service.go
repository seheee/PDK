package service

import "github.com/KumKeeHyun/PDK/health-check/adapter.go"

type ElasticClient interface {
	GetInput() chan<- adapter.Document
}

type KafkaConsumer interface {
	GetOutput() <-chan adapter.States
}
