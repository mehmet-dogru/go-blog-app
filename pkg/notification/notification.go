package notification

import (
	"fmt"
	"go-blog-app/config"
)

type NotificationClient interface {
	SendEmail(email string, subject string, content string) error
}

type notificationClient struct {
	config config.AppConfig
}

func NewNotificationClient(config config.AppConfig) NotificationClient {
	return notificationClient{
		config: config,
	}
}

func (n notificationClient) SendEmail(email string, subject string, content string) error {
	//TODO implement me
	fmt.Printf("Sent email to %v subject: '%v' content: '%v'\n", email, subject, content)
	return nil
}
