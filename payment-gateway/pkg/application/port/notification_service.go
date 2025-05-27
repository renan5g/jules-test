package port

import "context"

type NotificationMessage struct {
	To      string
	From    string
	Subject string
	Body    string
	Type    string // "EMAIL", "SMS"
}

type NotificationService interface {
	SendNotification(ctx context.Context, message NotificationMessage) error
}
