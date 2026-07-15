package kafka

import (
	"github.com/IBM/sarama"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
)

func GetTraceIDByProducer(headers []sarama.RecordHeader) string {
	for _, recordHeader := range headers {
		if string(recordHeader.Key) == string(log.TraceID) {
			return string(recordHeader.Value)
		}
	}
	return ""
}

func GetTraceID(headers []*sarama.RecordHeader) string {
	for _, recordHeader := range headers {
		if string(recordHeader.Key) == string(log.TraceID) {
			return string(recordHeader.Value)
		}
	}
	return ""
}
