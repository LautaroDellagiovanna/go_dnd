package services

import (
	"database/sql"
	"go_dnd/internal/models"
	"go_dnd/internal/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		userRepo: repositories.NewUserRepository(db),
	}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *UserService) GetUsers(u *models.User) ([]models.User, error) {
	return s.userRepo.Find(u)
}

func (s *UserService) AddUser(u *models.User) error {
	return s.userRepo.Save(u)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) UpdateUserByID(u *models.User) error {
	return s.userRepo.Update(u)
}
