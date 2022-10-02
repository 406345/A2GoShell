package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReportError(c *gin.Context, code int, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": err.Error(),
	})
}

func ReportErrorMessage(c *gin.Context, code int, err string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": err,
	})
}

func ReportData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}
