package repository

import (
	"web-server-101/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ent entity.User) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) CreateUser(ent entity.User) (entity.User, error) {
	result := u.db.Create(&ent)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return ent, nil
}

func ProvideUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
