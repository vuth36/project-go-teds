// Trong thư mục internal/app

package webbiz

import (
	"yourapp/internal/models"
	"yourapp/internal/storage"
)

type UserUseCase struct {
	userRepository *storage.UserRepository
}

func NewUserUseCase(userRepo *storage.UserRepository) *UserUseCase {
	return &UserUseCase{userRepository: userRepo}
}

func (uc *UserUseCase) GetUserByID(id int) (*models.User, error) {
	return uc.userRepository.GetUserByID(id)
}

// Các use case khác
