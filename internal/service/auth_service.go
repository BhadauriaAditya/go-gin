package service

import (
	"context"
	"errors"
	"go-crm/internal/domain/models"
	"go-crm/internal/domain/repository"
	"go-crm/internal/infrastructure/kafka"
	"go-crm/pkg/utils"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo     *repository.UserRepository
	redis    *redis.Client
	producer *kafka.Producer
}

func NewAuthService(repo *repository.UserRepository, redis *redis.Client, producer *kafka.Producer) *AuthService {
	return &AuthService{repo, redis, producer}
}

func (s *AuthService) Register(email, password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &models.User{Email: email, Password: string(hashed)}

	if err := s.repo.Create(user); err != nil {
		return err
	}

	// Send user created event to Kafka
	ctx := context.Background()
	if err := s.producer.SendMessage(ctx, kafka.TopicUserCreated, user.Email, user); err != nil {
		// Log the error but don't fail the registration
		// TODO: Add proper logging
		return nil
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid credentials")
	}
	token, _ := utils.GenerateJWT(user.ID)
	s.redis.Set(ctx, token, user.ID, 0)
	return token, nil
}

func (s *AuthService) VerifyToken(ctx context.Context, token string) (uint, error) {
	return utils.ParseJWT(token)
}

func (s *AuthService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}
