package repository

import (
	"errors"
	"gorm.io/gorm"
	"stroycity/pkg/model"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) AdminSignUp(admin model.Admin) error {
	return r.db.Model(&model.Admin{}).Create(&admin).Error
}

func (r *AdminRepository) AdminSignIn(login, password string) (model.Admin, error) {
	var admin model.Admin
	if err := r.db.Model(&model.Admin{}).Where("login = ?", login).First(&admin).Error; err != nil {
		return admin, errors.New("Пользователя с такой почтой не существует!")
	}

	if err := r.db.Model(&model.Admin{}).Where("login = ? AND password = ?", login, password).First(&admin).Error; err != nil {
		return admin, errors.New("Неверный пароль!")
	}

	return admin, nil
}
