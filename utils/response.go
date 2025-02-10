package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context,success bool, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"success": success, "message": message, "data": data})
}

func BadRequestResponse(c *gin.Context,success bool, message string, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{"success": success, "message": message, "data": data})
}

func NotFoundResponse(c *gin.Context,success bool, message string, data interface{}) {
	c.JSON(http.StatusNotFound, gin.H{"success": success, "message": message, "data": data})
}

func UnauthorizedResponse(c *gin.Context,success bool, message string, data interface{}) {
	c.JSON(http.StatusUnauthorized, gin.H{"success": success, "message": message, "data": data})
}

func StatusConflictResponse(c *gin.Context,success bool, message string, data interface{}) {
	c.JSON(http.StatusConflict, gin.H{"success": success, "message": message, "data": data})
}
