package service

import (
	"douyin_mall/internal/model"
	"douyin_mall/internal/model/dto"
	"douyin_mall/pkg/db"
	"douyin_mall/pkg/utils"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register(req *dto.RegisterRequest) error {
	// Check if username exists
	var count int64
	if err := db.DB.Model(&model.User{}).Where("username = ?", req.Username).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exists")
	}

	// Hash password
	salt := fmt.Sprintf("%06d", rand.Int31())
	hashedPassword := utils.MakePassword(req.Password, salt)

	// Create user
	user := model.User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		Email:        req.Email,
		Salt:         salt,
		PhoneNumber:  req.PhoneNumber,
		Role:         "user",
		Status:       1,
		CreatedAt:    time.Now(),
	}
	err := model.CreateUser(user).Error

	return err
}
