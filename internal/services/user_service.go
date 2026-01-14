package services

import (
	"errors"
	"task-manager/internal/models"
	"task-manager/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: &repositories.UserRepository{},
	}
}

func (s *UserService) Register(name , email , password string) (*models.User, error) {
	if name == "" || email == "" || password == "" {
		return nil , errors.New("all fields are required")
	}


	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil , err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(user); err != nil {
		return nil , err
	}
      return user , nil
}



//login

func (s *UserService) Login(email , password string) (string , error) {
	user , err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "" , errors.New("invalid email or password")
	}

	jwtService := NewJWTService()
	return jwtService.GenerateToken(user.ID , user.Email) 

}

