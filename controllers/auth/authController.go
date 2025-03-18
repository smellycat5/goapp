package auth

import (
	"Go/dto"
	"Go/models"
	"Go/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	UserService *services.UserService
}

func NewAuthController(userService *services.UserService) *AuthController {
	return &AuthController{UserService: userService}
}

func (ac *AuthController) Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userDto := dto.CreateUserDTO{
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
	}

	err := ac.UserService.Create(userDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginCredentials dto.LoginRequest

	if err := c.ShouldBind(&loginCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	jwtToken, loginErr := ac.UserService.Login(loginCredentials)

	if loginErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": loginErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in", "data": jwtToken})
}

func Logout(c *gin.Context) {}
