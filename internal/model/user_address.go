package model

import (
	"gorm.io/gorm"
	"time"
)

type UserAddress struct {
	gorm.Model
	ID        int       `json:"id" db:"id"`                 //ID
	UserID    int       `json:"user_id" db:"user_id"`       //用户ID
	City      string    `json:"city" db:"city"`             //用户所在城市
	Province  string    `json:"province" db:"province"`     //用户所在省份
	Country   string    `json:"country" db:"country"`       //用户所在国家
	IsDefault int       `json:"is_default" db:"is_default"` //是否默认地址
	CreatedAt time.Time `json:"created_at" db:"created_at"` //用户地址创建时间
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"` //用户地址最后更新时间
}
