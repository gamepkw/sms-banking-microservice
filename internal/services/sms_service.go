package service

import (
	"context"
	"time"

	"gopkg.in/Shopify/sarama.v1"
)

type smsService struct {
	contextTimeout time.Duration
	kafkaClient    sarama.Client
}

func NewSmsService(timeout time.Duration, kafka sarama.Client) SmsService {
	return &smsService{
		contextTimeout: timeout,
		kafkaClient:    kafka,
	}
}

type SmsService interface {
	SendSms(ctx context.Context)
}
