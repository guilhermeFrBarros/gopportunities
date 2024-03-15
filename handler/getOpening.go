package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeFrBarros/gopportunities/model"
	svc "github.com/guilhermeFrBarros/gopportunities/services"
)

func GetOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, svc.ErrParamIsRequired("id", "queryPameter").Error())
		return
	}

	opening := model.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening, 200)
}
