package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-blog-app/config"
	"go-blog-app/internal/domain"
	"go-blog-app/internal/dto"
	"go-blog-app/internal/helper"
	"go-blog-app/internal/repository"
	"go-blog-app/pkg/notification"
	"time"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Redis  redis.Client
	Config config.AppConfig
}

func NewUserService(repo repository.UserRepository, auth helper.Auth, redis redis.Client, config config.AppConfig) UserService {
	return UserService{
		Repo:   repo,
		Auth:   auth,
		Redis:  redis,
		Config: config,
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

	notificationClient := notification.NewNotificationClient(s.Config)
	err = notificationClient.SendEmail(user.Email, "Register Successfully", "Welcome to Blog App")
	if err != nil {
		return "", errors.New("error on sending email")
	}

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

func (s UserService) GetProfile(id uint) (*dto.ProfileInfo, error) {
	cachedUser, err := s.Redis.Get(context.Background(), fmt.Sprintf("user:%d", id)).Result()
	if err == nil {
		var profileInfo dto.ProfileInfo
		if err := json.Unmarshal([]byte(cachedUser), &profileInfo); err != nil {
			return nil, err
		}
		return &profileInfo, nil
	}

	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return nil, err
	}

	profileInfo := dto.ProfileInfo{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	userData, err := json.Marshal(profileInfo)
	if err != nil {
		return nil, err
	}
	if err := s.Redis.Set(context.Background(), fmt.Sprintf("user:%d", id), userData, 24*time.Hour).Err(); err != nil {
		return nil, err
	}

	return &profileInfo, nil
}
