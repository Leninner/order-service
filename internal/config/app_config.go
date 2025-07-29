package config

import (
	sharedConfig "github.com/leninner/shared/config"
)

type AppConfig struct {
	sharedConfig.Config
}

func LoadConfig() *AppConfig {
	config := sharedConfig.LoadConfig("order-service")
	return &AppConfig{
		Config: config,
	}
}

func (c *AppConfig) GetDatabaseConfig() sharedConfig.Config {
	return c.Config
}

func (c *AppConfig) GetKafkaConfig() sharedConfig.KafkaConfig {
	return c.Config.Kafka
}

func (c *AppConfig) GetTopicConfig() sharedConfig.TopicConfig {
	return c.Config.Kafka.Topics
} 