package repository

import (
	"context"
	"github.com/generic-org/payment-gateway/pkg/domain/entity"
)

type MerchantRepository interface {
	Save(ctx context.Context, merchant *entity.Merchant) error
	FindByID(ctx context.Context, id string) (*entity.Merchant, error)
	FindByAPIKey(ctx context.Context, apiKey string) (*entity.Merchant, error)
}
