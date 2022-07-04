package repositories

import "message/module/domain/model"

type NotificationRepository interface {
	Create(notification *model.Notification) error
	GetAllForUser(username string) ([]*model.Notification, error)
}