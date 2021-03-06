package persistance

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"user/module/domain/model"
	"user/module/domain/repositories"
)

type EmailVerificationRepositoryImpl struct {
	db *gorm.DB
}

func NewEmailVerificationRepositoryImpl(db *gorm.DB) repositories.EmailVerificationRepository {
	return &EmailVerificationRepositoryImpl{db: db}
}

func (e EmailVerificationRepositoryImpl) CreateEmailVerification(ver *model.EmailVerification) (*model.EmailVerification, error) {
	result := e.db.Create(&ver)
	fmt.Print(result)
	return ver, result.Error
}
func (e EmailVerificationRepositoryImpl) GetVerificationByUsername(username string) ([]model.EmailVerification, error) {
	var verifications []model.EmailVerification
	records := e.db.Find(&verifications, "username = ?", username)
	if records.Error != nil {
		log.Fatalln(records.Error)
	}
	if records.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return verifications, nil
}
