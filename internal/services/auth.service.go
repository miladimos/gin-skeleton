package services

import "gin-skeleton/internal/repository"

type AuthService interface {
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthRepository(authRepository repository.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}
