package handlers

import (
	saga "common/module/saga/messaging"
	events "common/module/saga/post_events"
	"message/module/application"
	"message/module/infrastructure/api"
)

type NotificationCommandHandler struct {
	service           *application.NotificationService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewNotificationCommandHandler(service *application.NotificationService, publisher saga.Publisher, subscriber saga.Subscriber) (*NotificationCommandHandler, error) {
	o := &NotificationCommandHandler{
		service:           service,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *NotificationCommandHandler) handle(command *events.PostNotificationCommand) {

	notification := api.MapNewPostNotification(command)
	var reply = &events.PostNotificationReply{}
	switch command.Type {
	case events.LikePost:
		_, err := handler.service.Create(notification)
		if err != nil {
			reply = api.MapPostNotificationReply(events.NotificationNotSent)
		}
		reply = api.MapPostNotificationReply(events.NotificationSent)
	case events.DislikePost:
		_, err := handler.service.Create(notification)
		if err != nil {
			reply = api.MapPostNotificationReply(events.NotificationNotSent)
		}
		reply = api.MapPostNotificationReply(events.NotificationSent)
	case events.CommentPost:
		_, err := handler.service.Create(notification)
		if err != nil {
			reply = api.MapPostNotificationReply(events.NotificationNotSent)
		}
		reply = api.MapPostNotificationReply(events.NotificationSent)
	default:
		reply = api.MapPostNotificationReply(events.UnknownReply)

		if reply.Type != events.UnknownReply {
			_ = handler.replyPublisher.Publish(reply)
		}

	}

}
