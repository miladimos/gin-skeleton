package services

import "github.com/miladimos/sanjabi/app/repository"

type AuthService interface {
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthRepository(authRepository repository.AuthRepository) AuthService {
	return &authRepository{
		authRepository: authRepository,
	}
}
