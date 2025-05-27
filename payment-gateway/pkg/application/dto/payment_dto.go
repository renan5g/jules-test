package dto

import "time"
// import "github.com/shopspring/decimal"

type CreatePaymentInput struct {
	MerchantID      string
	Amount          float64 // Placeholder; use decimal.Decimal
	Currency        string
	PaymentMethod   PaymentMethodInputDetails
	Customer        *CustomerInputDetails // Optional
	OrderID         string // Merchant's order ID
	Description     string
}

type PaymentMethodInputDetails struct {
	Type    string                 // "CREDIT_CARD", "PIX", "BOLETO"
	Details map[string]interface{} // e.g., card details, pix key
}

type CustomerInputDetails struct {
	ID      string
	Name    string
	Email   string
	Document string
}

type PaymentOutput struct {
	ID              string
	MerchantID      string
	Amount          float64 // Placeholder
	Currency        string
	Status          string
	OrderID         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	// Add other relevant fields like transaction IDs, payment method confirmation
}
