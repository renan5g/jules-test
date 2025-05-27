package port

import (
	"context"
)

type ProcessorPaymentRequest struct {
	Amount             int64    // Amount in cents
	Currency           string
	OrderID            string   // Internal payment ID
	Description        string
	PaymentMethodToken string   // Or more detailed payment method data
	CustomerInfo       map[string]string // e.g. email, name, ip_address
	// ... other common fields
}

type ProcessorPaymentResponse struct {
	TransactionID     string // Provider's transaction ID
	Status            string // e.g., "succeeded", "failed", "pending", "requires_action"
	AuthorizationCode string // If applicable
	RawResponse       []byte 
	ErrorCode         string 
	ErrorMessage      string 
	RedirectURL       string // If redirection is needed (e.g. 3DS)
}

type ProcessorRefundRequest struct {
	ProviderTransactionID string 
	Amount                int64  
	Currency              string
	Reason                string 
}

type ProcessorRefundResponse struct {
	RefundID      string 
	Status        string 
	RawResponse   []byte
	ErrorCode     string
	ErrorMessage  string
}

type PaymentProcessor interface {
	Process(ctx context.Context, request ProcessorPaymentRequest) (ProcessorPaymentResponse, error)
	Refund(ctx context.Context, request ProcessorRefundRequest) (ProcessorRefundResponse, error)
	// GetStatus, Cancel, Capture etc. might be added
}
