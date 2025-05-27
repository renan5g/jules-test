package entity

import "time"

type PaymentMethodType string

const (
	PaymentMethodTypeCreditCard PaymentMethodType = "CREDIT_CARD"
	PaymentMethodTypeBoleto     PaymentMethodType = "BOLETO"
	PaymentMethodTypePix        PaymentMethodType = "PIX"
	// Add other types as needed
)

type PaymentMethod struct {
	ID              string
	CustomerID      string // If stored for a customer
	Type            PaymentMethodType
	Details         map[string]interface{} // e.g., card brand, last four for CC
	Token           string // If tokenized by a provider
	IsDefault       bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
