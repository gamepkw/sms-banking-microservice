package service

import (
	"context"
	"fmt"
	"strings"

	consumer "github.com/gamepkw/sms-banking-microservice/pkg/kafka/consumer"
	"gopkg.in/Shopify/sarama.v1"
)

func (a *smsService) SendSms(ctx context.Context) {
	topic := "sms"
	var partition int32 = 0
	var offset int64 = sarama.OffsetNewest
	for {
		select {
		case <-ctx.Done():
			fmt.Println("SendSms stopped")
			return
		default:
			// Consume a Kafka message
			message := consumer.RunKafkaConsumer(a.kafkaClient, topic, partition, offset)
			message_split := strings.Split(message, "|")
			fmt.Printf("OTP='%s' for reset password\n%s\n",
				message_split[0],
				strings.Repeat("-", 60))
			// time.Sleep(time.Second)
		}
	}
}
