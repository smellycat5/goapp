package services

import (
	"Go/dto"
	"Go/models"
	"Go/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) Create(createUserDto dto.CreateUserDTO) error {

	var existingUser models.User
	result := s.DB.Where("email =?", createUserDto.Email).Find(&existingUser)
	if result.RowsAffected > 0 {
		return errors.New("email already taken")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(createUserDto.Password), 10)

	user := models.User{
		Name:     createUserDto.Name,
		Email:    createUserDto.Email,
		Password: string(hashedPassword),
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (s *UserService) Login(request dto.LoginRequest) (string, error) {

	var token string
	var existingUser models.User

	result := s.DB.Where("email =?", request.Email).Find(&existingUser)
	if result.RowsAffected == 0 {
		return token, errors.New("invalid username or password")
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(request.Password))
	if err != nil {
		return token, errors.New("invalid username or password")
	}

	token, tokenErr := utils.GenerateJWT(existingUser)
	if tokenErr != nil {
		return token, errors.New("error generating token")
	}

	return token, nil
}

//func (s *UserService) View(id int) (*models.User, error) {
//
//	var user models.User
//	result := s.DB.Where("id = ?", id).Find(&user)
//
//	log.Print(result)
//}
