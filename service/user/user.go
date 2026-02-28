package user_service

import (
	"web-server-101/repository"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type UserService interface {
	HandleCreateUser(c fiber.Ctx) error
}

type userService struct {
	db             *gorm.DB
	userRepository repository.UserRepository
}

func ProvideUserService(db *gorm.DB) UserService {
	return &userService{
		db:             db,
		userRepository: repository.ProvideUserRepository(db),
	}
}
