package middleware

import (
	"groove-app/pkg/jwtutil"
	"groove-app/pkg/resp"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	tk := c.Request.Header.Get("Authorization")
	pl, msg, ok := jwtutil.Parse(tk)
	if !ok {
		resp.UnAuth(c, msg)
		c.Abort()
		return
	}

	c.Set("token", pl)
	c.Next()
}
