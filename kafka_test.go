package main

import (
	"github.com/segmentio/kafka-go"
	"testing"
)

func TestKafkaConnection(t *testing.T) {
	brokers := []string{"localhost:9092"}

	conn, err := kafka.Dial("tcp", brokers[0])
	if err != nil {
		t.Fatalf("Gagal terkoneksi ke Kafka: %v", err)
	}
	defer conn.Close()

	// Cek broker tersedia
	_, err = conn.Brokers()
	if err != nil {
		t.Fatalf("Gagal mendapatkan daftar broker: %v", err)
	}
}
