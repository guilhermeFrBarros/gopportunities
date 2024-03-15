package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeFrBarros/gopportunities/model"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := model.UpdateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil { // analisa o dados da requisição e preenche o objeto dentro dos parentes
		logger.ErrF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.ErrF("Validation: error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		print("Upda")
		return
	}

	opening := model.Opening{}
	if err := db.First(&opening, request.Id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	//Opdate opening
	opening.Updated(&request)
	if err := db.Save(&opening).Error; err != nil {
		logger.ErrF("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error update opening in DB")
		return
	}

	sendSuccess(ctx, "Update opening sucess", opening, http.StatusOK)
}
