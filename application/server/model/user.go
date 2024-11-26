package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
	FullName string
	Role     *bool // 0为游客，1为亚运主办方
}
