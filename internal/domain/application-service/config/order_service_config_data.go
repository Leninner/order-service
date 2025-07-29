package config

import sharedConfig "github.com/leninner/shared/config"

type OrderServiceConfigData struct {
	PaymentRequestTopicName string
	PaymentResponseTopicName string
	RestaurantApprovalRequestTopicName string
	RestaurantApprovalResponseTopicName string
}

func NewOrderServiceConfigData() *OrderServiceConfigData {
	return &OrderServiceConfigData{
		PaymentRequestTopicName: "payment-request",
		PaymentResponseTopicName: "payment-response",
		RestaurantApprovalRequestTopicName: "restaurant-approval-request",
		RestaurantApprovalResponseTopicName: "restaurant-approval-response",
	}
}

func NewOrderServiceConfigDataFromShared(sharedConfig sharedConfig.Config) *OrderServiceConfigData {
	return &OrderServiceConfigData{
		PaymentRequestTopicName: sharedConfig.Kafka.Topics.PaymentRequest,
		PaymentResponseTopicName: sharedConfig.Kafka.Topics.PaymentResponse,
		RestaurantApprovalRequestTopicName: sharedConfig.Kafka.Topics.RestaurantApprovalRequest,
		RestaurantApprovalResponseTopicName: sharedConfig.Kafka.Topics.RestaurantApprovalResponse,
	}
}