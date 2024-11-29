package handler

import (
	"fmt"

	"github.com/xvrzhao/groove-scaffold/pkg/jwtutil"
	"github.com/xvrzhao/groove-scaffold/pkg/logger"
	"github.com/xvrzhao/groove-scaffold/pkg/resp"

	"github.com/gin-gonic/gin"
)

type authCtrl struct {
	authSvc
}

func NewAuthCtrl() *authCtrl {
	return new(authCtrl)
}

func (ctrl authCtrl) Login(c *gin.Context) {
	type req struct {
		Username string `binding:"required"`
		Password string `binding:"required"`
	}
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}

	ok, msg, tk, u, err := ctrl.authSvc.Login(r.Username, r.Password)
	if err != nil {
		logger.Error(fmt.Errorf("userSvc.Login failed: %w", err).Error())
		resp.InternalErr(c)
		return
	}

	if ok {
		resp.Success(c, gin.H{
			"token":    tk,
			"userInfo": u,
		})
	} else {
		resp.Fail(c, -1, msg)
	}
}

func (ctrl authCtrl) ChangePassword(c *gin.Context) {
	type req struct {
		Password string `binding:"required,alphanum,gte=8,lte=20"`
		Confirm  string `binding:"required,eqfield=Password"`
	}
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}

	v, _ := c.Get("token")
	token := v.(*jwtutil.TokenPayload)

	if err := ctrl.authSvc.ChangePassword(token.UserID, r.Password); err != nil {
		logger.Error(fmt.Errorf("userSvc.ChangePassword failed: %w", err).Error())
		resp.InternalErr(c)
		return
	}

	resp.Success(c, nil)
}
