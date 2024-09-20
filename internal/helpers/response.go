package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorMessage struct {
	statusCode int
	Message    string
}

func NewResponseMessage(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, ErrorMessage{statusCode, message})
}
