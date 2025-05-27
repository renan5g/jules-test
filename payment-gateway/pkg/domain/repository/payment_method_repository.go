package repository

import (
	"context"
	"github.com/generic-org/payment-gateway/pkg/domain/entity"
)

type PaymentMethodRepository interface {
	Save(ctx context.Context, paymentMethod *entity.PaymentMethod) error
	FindByID(ctx context.Context, id string) (*entity.PaymentMethod, error)
	FindByCustomerID(ctx context.Context, customerID string) ([]*entity.PaymentMethod, error)
	// Add other methods like Update, Delete, etc.
}
