package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BodyCode int

const (
	successCode     BodyCode = 0
	unAuthCode      BodyCode = -401
	badRequestCode  BodyCode = -400
	internalErrCode BodyCode = -500
)

type ResponseBody struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, ResponseBody{
		Code: int(successCode),
		Data: data,
	})
}

func Fail(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponseBody{
		Code:    code,
		Message: message,
	})
}

func UnAuth(c *gin.Context, message string) {
	c.JSON(http.StatusOK, ResponseBody{
		Code:    int(unAuthCode),
		Message: message,
	})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, ResponseBody{
		Code:    int(badRequestCode),
		Message: message,
	})
}

func InternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, ResponseBody{
		Code:    int(internalErrCode),
		Message: "服务端出错",
	})
}
