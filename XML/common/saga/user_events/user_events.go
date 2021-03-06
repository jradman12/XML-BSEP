package user_events

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserId    uuid.UUID
	Username  string
	Email     string
	FirstName string
	LastName  string
}
type PostUser struct {
	Id        primitive.ObjectID
	UserId    uuid.UUID
	Username  string
	Email     string
	FirstName string
	LastName  string
}

type UserCommandType int8

const (
	UpdateUser UserCommandType = iota
	CreateUser
	DeleteUser
	ActivateUser
	ChangeProfileStatus
	UpdateUserProfessionalDetails
	ChangeEmail
	RollbackUser
	UnknownCommand
)

type UserReplyType int8

const (
	UserUpdated UserReplyType = iota
	UserRolledBack
	UserDeleted
	UserCreated
	UserActivated
	ChangedEmail
	ChangedEmailRolledBack
	ProfileStatusChanged
	UnknownReply
)

type UserCommand struct {
	User User
	Type UserCommandType
}

type UserReply struct {
	Type     UserReplyType
	PostUser PostUser
}

type CreateUserCommand struct {
	User User
	Type UserCommandType
}

type CreateUserReply struct {
	Type     UserReplyType
	PostUser PostUser
}

type ActivateUserCommand struct {
	Type UserCommandType
	Id   uuid.UUID
}

type ActivateUserReply struct {
	Type UserCommandType
}

type ConnectionUser struct {
	UserUID       string
	Username      string
	FirstName     string
	LastName      string
	ProfileStatus string
	Interests     []string
	Skills        []string
	Educations    []string
	Experiences   []string
}

type ConnectionUserCommand struct {
	User ConnectionUser
	Type UserCommandType
}

type UserConnectionReply struct {
	Type UserReplyType
}
