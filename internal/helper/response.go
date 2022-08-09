package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK,
		gin.H{
			"data":  data,
			"error": nil,
		})
}

func ErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest,
		gin.H{
			"data":  nil,
			"error": err.Error(),
		})
}

func BadRequestResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest,
		gin.H{
			"data":  nil,
			"error": err.Error(),
		})
}

func InternalServerResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError,
		gin.H{
			"data":  nil,
			"error": "Internal server error",
		})
}

func UnauthorizedResponse(c *gin.Context) {
	c.JSON(http.StatusUnauthorized,
		gin.H{
			"data":  nil,
			"error": "Unauthorized response",
		})
}
