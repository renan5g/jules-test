package dto

// Using application DTOs directly for now, can be specific API DTOs if needed for decoupling
// from internal application DTOs or for API versioning.
// For example, we might want to expose only certain fields or structure them differently.

import (
	appDTO "github.com/generic-org/payment-gateway/pkg/application/dto"
)

// CreatePaymentRequest mirrors appDTO.CreatePaymentInput but for the API layer
type CreatePaymentRequest struct {
	MerchantID    string                         `json:"merchant_id" binding:"required"`
	Amount        float64                        `json:"amount" binding:"required,gt=0"` // Amount in major units
	Currency      string                         `json:"currency" binding:"required,len=3"`
	PaymentMethod PaymentMethodRequestDetails    `json:"payment_method" binding:"required"`
	Customer      *CustomerRequestDetails        `json:"customer,omitempty"`
	OrderID       string                         `json:"order_id,omitempty"`
	Description   string                         `json:"description,omitempty"`
}

type PaymentMethodRequestDetails struct {
	Type    string                 `json:"type" binding:"required"` // e.g., "CREDIT_CARD", "PIX", "BOLETO"
	Details map[string]interface{} `json:"details" binding:"required"` // e.g., card details, pix key
}

type CustomerRequestDetails struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty" binding:"omitempty,email"`
	Document string `json:"document,omitempty"`
}

// PaymentResponse mirrors appDTO.PaymentOutput but for the API layer
type PaymentResponse struct {
	ID         string  `json:"id"`
	MerchantID string  `json:"merchant_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Status     string  `json:"status"`
	OrderID    string  `json:"order_id,omitempty"`
	CreatedAt  string  `json:"created_at"` // Using string for simplicity, can be time.Time
	// Add other relevant fields
}

// RefundRequest mirrors appDTO.RefundPaymentInput
type RefundRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"` // Amount in major units
	Reason string  `json:"reason,omitempty"`
}

// RefundResponse mirrors appDTO.RefundOutput
type RefundResponse struct {
	RefundID       string  `json:"refund_id"`
	PaymentID      string  `json:"payment_id"`
	Status         string  `json:"status"`
	AmountRefunded float64 `json:"amount_refunded"`
	Currency       string  `json:"currency"`
	ProcessedAt    string  `json:"processed_at"` // Using string for simplicity
}

// Helper function to convert from application DTO to API DTO (example)
func ToPaymentResponse(appOutput *appDTO.PaymentOutput) *PaymentResponse {
	if appOutput == nil {
		return nil
	}
	return &PaymentResponse{
		ID:         appOutput.ID,
		MerchantID: appOutput.MerchantID,
		Amount:     appOutput.Amount,
		Currency:   appOutput.Currency,
		Status:     appOutput.Status,
		OrderID:    appOutput.OrderID,
		CreatedAt:  appOutput.CreatedAt.String(), // Or format as desired
	}
}

// Helper function to convert from API DTO to application DTO (example)
func (cpr *CreatePaymentRequest) ToApplicationInput() *appDTO.CreatePaymentInput {
	// Basic conversion, real implementation might need more logic
	return &appDTO.CreatePaymentInput{
		MerchantID:  cpr.MerchantID,
		Amount:      cpr.Amount,
		Currency:    cpr.Currency,
		PaymentMethod: appDTO.PaymentMethodInputDetails{
			Type: cpr.PaymentMethod.Type,
			Details: cpr.PaymentMethod.Details,
		},
		// Customer details conversion
		OrderID:     cpr.OrderID,
		Description: cpr.Description,
	}
}

func (rr *RefundRequest) ToApplicationInput(paymentID, merchantID string) *appDTO.RefundPaymentInput {
	return &appDTO.RefundPaymentInput{
		PaymentID:  paymentID,
		Amount:     rr.Amount,
		Reason:     rr.Reason,
		MerchantID: merchantID, // Assuming merchant ID is passed for validation
	}
}
