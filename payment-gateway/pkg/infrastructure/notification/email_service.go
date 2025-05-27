package notification

import (
	"context"
	"github.com/generic-org/payment-gateway/pkg/application/port"
)

type EmailNotificationService struct {
	// Email client config, e.g., SMTP server details
}

func NewEmailNotificationService(/* config */) port.NotificationService {
	return &EmailNotificationService{}
}

func (s *EmailNotificationService) SendNotification(ctx context.Context, message port.NotificationMessage) error {
	// TODO: Implement email sending logic
	return nil
}
