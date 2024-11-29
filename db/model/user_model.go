package model

import (
	"github.com/xvrzhao/groove-scaffold/pkg/basemodel"
)

const TableNameUser = "users"

type User struct {
	basemodel.Model
	Username string `json:"username" binding:"required,max=50"`
	Nickname string `json:"nickname" binding:"required,max=50"`
	Phone    string `json:"phone" binding:"required,max=20"`
	Email    string `json:"email" binding:"required,email"`
	PwdHash  string `json:"-"`
	Salt     string `json:"-"`
}

func (*User) TableName() string {
	return TableNameUser
}
