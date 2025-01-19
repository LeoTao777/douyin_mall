package model

import (
	"douyin_mall/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID                  int       `json:"id" db:"id"`                                       //用户ID
	Username            string    `json:"username" db:"username"`                           //用户名
	Email               string    `json:"email" db:"email"`                                 //用户邮箱
	PhoneNumber         string    `json:"phone_number" db:"phone_number"`                   //用户手机号
	PasswordHash        string    `json:"password_hash" db:"password_hash"`                 //加密的用户密码
	Salt                string    `json:"salt" db:"salt"`                                   //用于密码哈希的盐值
	ResetToken          string    `json:"reset_token" db:"reset_token"`                     //用户找回密码时生成的临时令牌
	ResetTokenExpiry    time.Time `json:"reset_token_expiry" db:"reset_token_expiry"`       //重置令牌的有效期
	LastLogin           time.Time `json:"last_login" db:"last_login"`                       //用户上次登录时间
	FailedLoginAttempts int       `json:"failed_login_attempts" db:"failed_login_attempts"` //登录失败次数
	Status              int8      `json:"status" db:"default:1"`                            //用户状态
	Role                string    `json:"role" db:"role"`                                   //用户角色
	IsVerified          bool      `json:"is_verified" db:"is_verified"`                     //标记用户是否通过邮箱/手机验证
	Address             string    `json:"address" db:"address"`                             //用户默认的送货地址
	CreatedAt           time.Time `json:"created_at" db:"created_at"`                       //用户账户创建时间
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`                       //用户信息最后更新时间
	Bio                 string    `json:"bio" db:"bio"`                                     //用户个人简介
}

func (table *User) TableName() string {
	return "user"
}

func FindUserByName(name string) User {
	user := User{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func CreateUser(user User) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user User) *gorm.DB {
	return utils.DB.Delete(&user)
}
