package repository

import (
	"context"
	"github.com/generic-org/payment-gateway/pkg/domain/entity"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *entity.Payment) error
	FindByID(ctx context.Context, id string) (*entity.Payment, error)
	Update(ctx context.Context, payment *entity.Payment) error
	// Add other methods like FindByMerchantID, etc.
}
