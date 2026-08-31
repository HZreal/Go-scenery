package rabbitmqdemo

import "testing"

func TestConfigFromEnvUsesHAProxyPortByDefault(t *testing.T) {
	t.Setenv("RABBITMQ_URL", "")
	t.Setenv("RABBITMQ_QUEUE", "")

	cfg, err := ConfigFromEnv()
	if err != nil {
		t.Fatalf("ConfigFromEnv() error = %v", err)
	}

	if cfg.URL != "amqp://guest:guest@127.0.0.1:57009/" {
		t.Fatalf("URL = %q, want HAProxy listener on 57009", cfg.URL)
	}
	if cfg.Queue != "rabbitmq_cluster_demo" {
		t.Fatalf("Queue = %q", cfg.Queue)
	}
}

func TestConfigFromEnvAllowsOverrides(t *testing.T) {
	t.Setenv("RABBITMQ_URL", "amqp://guest:guest@localhost:57010/")
	t.Setenv("RABBITMQ_QUEUE", "orders")

	cfg, err := ConfigFromEnv()
	if err != nil {
		t.Fatalf("ConfigFromEnv() error = %v", err)
	}

	if cfg.URL != "amqp://guest:guest@localhost:57010/" {
		t.Fatalf("URL = %q", cfg.URL)
	}
	if cfg.Queue != "orders" {
		t.Fatalf("Queue = %q", cfg.Queue)
	}
}

func TestProducerMessageCountFromEnvUsesDefault(t *testing.T) {
	t.Setenv("RABBITMQ_MESSAGE_COUNT", "")

	count, err := ProducerMessageCountFromEnv()
	if err != nil {
		t.Fatalf("ProducerMessageCountFromEnv() error = %v", err)
	}

	if count != 10 {
		t.Fatalf("count = %d, want 10", count)
	}
}

func TestProducerMessageCountFromEnvAllowsOverride(t *testing.T) {
	t.Setenv("RABBITMQ_MESSAGE_COUNT", "3")

	count, err := ProducerMessageCountFromEnv()
	if err != nil {
		t.Fatalf("ProducerMessageCountFromEnv() error = %v", err)
	}

	if count != 3 {
		t.Fatalf("count = %d, want 3", count)
	}
}

func TestProducerMessageCountFromEnvRejectsInvalidMessageCount(t *testing.T) {
	t.Setenv("RABBITMQ_MESSAGE_COUNT", "0")

	if _, err := ProducerMessageCountFromEnv(); err == nil {
		t.Fatal("ProducerMessageCountFromEnv() error = nil, want invalid message count error")
	}
}
