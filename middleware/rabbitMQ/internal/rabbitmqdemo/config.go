package rabbitmqdemo

import (
	"fmt"
	"os"
	"strconv"
)

const (
	defaultURL           = "amqp://guest:guest@127.0.0.1:57009/"
	defaultQueue         = "rabbitmq_cluster_demo"
	defaultProducerCount = 10
)

type Config struct {
	URL   string
	Queue string
}

func ConfigFromEnv() (Config, error) {
	cfg := Config{
		URL:   envOrDefault("RABBITMQ_URL", defaultURL),
		Queue: envOrDefault("RABBITMQ_QUEUE", defaultQueue),
	}

	return cfg, nil
}

func ProducerMessageCountFromEnv() (int, error) {
	count := defaultProducerCount
	if raw := os.Getenv("RABBITMQ_MESSAGE_COUNT"); raw != "" {
		n, err := strconv.Atoi(raw)
		if err != nil || n <= 0 {
			return 0, fmt.Errorf("RABBITMQ_MESSAGE_COUNT must be a positive integer")
		}
		count = n
	}

	return count, nil
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
