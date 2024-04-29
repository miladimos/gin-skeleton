package services

import (
	"gin-skeleton/internal/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

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

func Register(c *gin.Context) {
	//
}

func Login(c *gin.Context) {
	// Perform authentication
	// ...

	// Generate a JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "example@example.com"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and send it as a response
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": t})
}
