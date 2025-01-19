package model

import (
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
	FirstName           string    `json:"first_name" db:"first_name"`                       //用户名字
	LastName            string    `json:"last_name" db:"last_name"`                         //用户姓氏
	ResetToken          string    `json:"reset_token" db:"reset_token"`                     //用户找回密码时生成的临时令牌
	ResetTokenExpiry    time.Time `json:"reset_token_expiry" db:"reset_token_expiry"`       //重置令牌的有效期
	LastLogin           time.Time `json:"last_login" db:"last_login"`                       //用户上次登录时间
	FailedLoginAttempts int       `json:"failed_login_attempts" db:"failed_login_attempts"` //登录失败次数
	Status              string    `json:"status" db:"status"`                               //用户状态
	Role                string    `json:"role" db:"role"`                                   //用户角色
	IsVerified          bool      `json:"is_verified" db:"is_verified"`                     //标记用户是否通过邮箱/手机验证
	Address             string    `json:"address" db:"address"`                             //用户默认的送货地址
	City                string    `json:"city" db:"city"`                                   //用户所在城市
	Province            string    `json:"province" db:"province"`                           //用户所在省份
	Country             string    `json:"country" db:"country"`                             //用户所在国家
	PostalCode          string    `json:"postal_code" db:"postal_code"`                     //邮政编码
	CreatedAt           time.Time `json:"created_at" db:"created_at"`                       //用户账户创建时间
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`                       //用户信息最后更新时间
	Bio                 string    `json:"bio" db:"bio"`                                     //用户个人简介
}

func (table *User) TableName() string {
	return "user"
}
