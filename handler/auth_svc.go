package handler

import (
	"fmt"
	"groove-app/db"
	"groove-app/db/model"
	"groove-app/pkg/jwtutil"
	"groove-app/pkg/pwdutil"
	"os/user"

	"gorm.io/gorm"
)

type authSvc struct{}

func (authSvc) Login(username, password string) (ok bool, msg, token string, u *model.User, err error) {
	u = new(model.User)
	err = db.Client.Model(u).Take(u, "username = ?", username).Error
	if err == gorm.ErrRecordNotFound {
		return false, "用户不存在", "", nil, nil
	}
	if err != nil {
		err = fmt.Errorf("failed to query user: %w", err)
		return
	}

	if !pwdutil.VerifySaltHashPwd(password, u.Salt, u.PwdHash) {
		return false, "密码错误", "", nil, nil
	}

	token, err = jwtutil.Gen(jwtutil.TokenPayload{
		UserID:   int(u.ID),
		Username: u.Username,
	})
	if err != nil {
		return false, "", "", nil, fmt.Errorf("failed to gen token: %w", err)
	}

	db.Client.Model(u).Preload("Roles.Permissions").Take(u)

	return true, "登录成功", token, u, nil
}

func (authSvc) ChangePassword(userID int, newPass string) error {
	pwd, salt := pwdutil.SaltHashPwd(newPass, 8)
	u := new(user.User)
	err := db.Client.Model(&u).Where("id = ?", userID).Updates(map[string]any{"pwd_hash": pwd, "salt": salt}).Error
	if err != nil {
		err = fmt.Errorf("update failed: %w", err)
	}
	return err
}
