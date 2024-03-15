package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeFrBarros/gopportunities/config"
	"gorm.io/gorm"
)

var ( // declarando as variaveis globais do handler
	logger *config.Logger
	db     *gorm.DB
)

func InitalizeHandler() { // inicia as variaveis golbais do pacote handler
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}, status int) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(status, gin.H{
		"message": fmt.Sprintf("operation froim handler: %s suaccessfull", op),
		"data":    data,
	})
}
