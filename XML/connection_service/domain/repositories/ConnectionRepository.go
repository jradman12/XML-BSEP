package repositories

import (
	"connection/module/domain/dto"
	connectionModel "connection/module/domain/model"
)

type ConnectionRepository interface {
	CreateConnection(connection *connectionModel.Connection) (*dto.ConnectionResponse, error)
	AcceptConnection(connection *connectionModel.Connection) (*dto.ConnectionResponse, error)
	GetAllConnectionForUser(userUid string) (userNodes []*connectionModel.User, error1 error)
	GetAllConnectionRequestsForUser(userUid string) (userNodes []*connectionModel.User, error1 error)
	ConnectionStatusForUsers(senderId string, receiverId string) (*dto.ConnectionResponse, error)
}
