package service

import (
	"errors"
	"go-blog-app/internal/domain"
	"go-blog-app/internal/dto"
	"go-blog-app/internal/helper"
	"go-blog-app/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func NewUserService(repo repository.UserRepository, auth helper.Auth) UserService {
	return UserService{
		Repo: repo,
		Auth: auth,
	}
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {

	hPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.Repo.CreateUser(domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hPassword,
	})

	return s.Auth.GenerateToken(user.ID, user.Email, user.Role)
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUser(email)

	return &user, err
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.findUserByEmail(email)

	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}

	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.Role)
}

func (s UserService) GetProfile(id uint) (domain.User, error) {
	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
