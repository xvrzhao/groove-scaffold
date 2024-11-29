package handler

import (
	"groove-app/db/model"
	"groove-app/pkg/page"
	"groove-app/pkg/resp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userCtrl struct {
	userSvc
}

func NewUserCtrl() userCtrl {
	return userCtrl{
		userSvc: NewUserSvc(),
	}
}

func (ctrl userCtrl) Page(c *gin.Context) {
	offset, limit := page.GetPageSize(c, 20, 999)

	users, total, err := ctrl.userSvc.Page(offset, limit)
	if err != nil {
		resp.Fail(c, 500, err.Error())
	} else {
		resp.Success(c, gin.H{
			"users": users,
			"total": total,
		})
	}
}

func (ctrl userCtrl) Create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}

	if err := ctrl.userSvc.Create(&user); err != nil {
		resp.Fail(c, 500, err.Error())
	} else {
		resp.Success(c, user)
	}
}

func (ctrl userCtrl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.BadRequest(c, "invalid path param: id")
		return
	}

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}

	if err := ctrl.userSvc.Update(id, &user); err != nil {
		resp.Fail(c, 500, err.Error())
	} else {
		resp.Success(c, user)
	}
}

func (ctrl userCtrl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.BadRequest(c, "invalid path param: id")
		return
	}

	if err := ctrl.userSvc.Delete(id); err != nil {
		resp.Fail(c, 500, err.Error())
	} else {
		resp.Success(c, nil)
	}
}
