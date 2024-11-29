package handler

import (
	"{{.module}}/db/model"
	"{{.module}}/pkg/page"
	"{{.module}}/pkg/resp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type {{.lowerCamelModel}}Ctrl struct {
	{{.lowerCamelModel}}Svc
}

func New{{.upperCamelModel}}Ctrl() *{{.lowerCamelModel}}Ctrl {
	return new({{.lowerCamelModel}}Ctrl)
}

func (ctrl {{.lowerCamelModel}}Ctrl) Page(c *gin.Context) {
	offset, limit := page.GetPageSize(c, 20, 999)

	{{.lowerCamelModel}}s, total, err := ctrl.{{.lowerCamelModel}}Svc.Page(offset, limit)
	if err != nil {
		resp.Fail(c, 500, err.Error())
	} else {
		resp.Success(c, gin.H{
			"{{.lowerCamelModel}}s": {{.lowerCamelModel}}s,
			"total": total,
		})
	}
}

func (ctrl {{.lowerCamelModel}}Ctrl) Create(c *gin.Context) {
	var {{.lowerCamelModel}} model.{{.upperCamelModel}}
	if err := c.ShouldBindJSON(&{{.lowerCamelModel}}); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}

	if err := ctrl.{{.lowerCamelModel}}Svc.Create(&{{.lowerCamelModel}}); err != nil {
		resp.Fail(c, 500, err.Error())
	} else {
		resp.Success(c, {{.lowerCamelModel}})
	}
}

func (ctrl {{.lowerCamelModel}}Ctrl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.BadRequest(c, "invalid path param: id")
		return
	}

	var {{.lowerCamelModel}} model.{{.upperCamelModel}}
	if err := c.ShouldBindJSON(&{{.lowerCamelModel}}); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}

	if err := ctrl.{{.lowerCamelModel}}Svc.Update(id, &{{.lowerCamelModel}}); err != nil {
		resp.Fail(c, 500, err.Error())
	} else {
		resp.Success(c, {{.lowerCamelModel}})
	}
}

func (ctrl {{.lowerCamelModel}}Ctrl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.BadRequest(c, "invalid path param: id")
		return
	}

	if err := ctrl.{{.lowerCamelModel}}Svc.Delete(id); err != nil {
		resp.Fail(c, 500, err.Error())
	} else {
		resp.Success(c, nil)
	}
}
