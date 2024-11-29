package handler

import (
	"fmt"
	"groove-app/db"
	"groove-app/db/model"
	"groove-app/pkg/pwdutil"

	"gorm.io/gorm/clause"
)

type userSvc struct {
	defaultUserPassword string
}

func NewUserSvc() userSvc {
	return userSvc{
		defaultUserPassword: "abcd1234",
	}
}

func (userSvc) Page(offset, limit int) (users []model.User, total int64, err error) {
	if err := db.Client.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	users = make([]model.User, 0, limit)
	if err := db.Client.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (svc userSvc) Create(user *model.User) error {
	pw, salt := pwdutil.SaltHashPwd(svc.defaultUserPassword, 8)
	user.PwdHash, user.Salt = pw, salt

	if err := db.Client.Create(user).Error; err != nil {
		return fmt.Errorf("failed create user: %w", err)
	}

	if err := db.Client.Preload(clause.Associations).First(user).Error; err != nil {
		return fmt.Errorf("failed to load user: %w", err)
	}
	return nil
}

func (userSvc) Update(id int, user *model.User) error {
	if err := db.Client.Where("id = ?", id).Updates(user).Error; err != nil {
		return fmt.Errorf("failed create user: %w", err)
	}
	user.ID = uint(id)

	if err := db.Client.Preload(clause.Associations).First(user).Error; err != nil {
		return fmt.Errorf("failed to load user: %w", err)
	}
	return nil
}

func (userSvc) Delete(id int) error {
	return db.Client.Delete(&model.User{}, id).Error
}
