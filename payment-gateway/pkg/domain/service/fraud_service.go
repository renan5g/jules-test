package service

import (
	"context"
	"github.com/generic-org/payment-gateway/pkg/domain/entity"
)

type FraudRiskLevel string

const (
	RiskLow    FraudRiskLevel = "LOW"
	RiskMedium FraudRiskLevel = "MEDIUM"
	RiskHigh   FraudRiskLevel = "HIGH"
)

type FraudAssessment struct {
	Score     int
	RiskLevel FraudRiskLevel
	Reason    string
}

type FraudDetectionService interface {
	AssessRisk(ctx context.Context, payment *entity.Payment) (FraudAssessment, error)
}
