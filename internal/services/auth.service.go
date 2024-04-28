package services

import (
	"fmt"
	"gin-skeleton/internal/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func handleLogin(c *gin.Context) {
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

func authenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Verify the token
		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func handleProtected(c *gin.Context) {
	// Protected resource handler
	// ...
}
