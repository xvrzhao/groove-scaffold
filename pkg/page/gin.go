package page

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPageSize(c *gin.Context, defaultSize, maxSize int) (offset, limit int) {
	pageStr := c.Query("page")
	perStr := c.Query("size")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(perStr)
	if err != nil {
		size = defaultSize
	}
	if size > maxSize {
		size = maxSize
	}

	return (page - 1) * size, size
}
