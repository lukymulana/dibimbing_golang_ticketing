package service

import (
	"dibimbing_golang_ticketing/dto"
	"dibimbing_golang_ticketing/entity"
	"dibimbing_golang_ticketing/repository"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type AuthService interface {
	Register(dto.RegisterDTO) (*entity.User, error)
	Login(dto.LoginDTO) (*entity.User, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(input dto.RegisterDTO) (*entity.User, error) {
	_, err := s.repo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("username already exists")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	user := &entity.User{
		Username: input.Username,
		Password: string(hash),
		Role:     "user",
	}
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) Login(input dto.LoginDTO) (*entity.User, error) {
	user, err := s.repo.GetUserByUsername(input.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid password")
	}
	return user, nil
}
