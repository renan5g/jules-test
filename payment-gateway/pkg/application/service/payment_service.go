package service

import (
	"context"
	"github.com/generic-org/payment-gateway/pkg/application/dto"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, input dto.CreatePaymentInput) (*dto.PaymentOutput, error)
	GetPayment(ctx context.Context, paymentID string) (*dto.PaymentOutput, error)
	RefundPayment(ctx context.Context, input dto.RefundPaymentInput) (*dto.RefundOutput, error)
	// Potentially ProcessPayment, CancelPayment etc.
}
