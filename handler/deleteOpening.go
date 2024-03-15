package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeFrBarros/gopportunities/model"
	svc "github.com/guilhermeFrBarros/gopportunities/services"
)

func DeleteOpeningHandler(ctx *gin.Context) {

	//http://localhost:8080/api/v1/opening?id=333
	id := ctx.Query("id") // ctx.Param
	if id == "" {
		sendError(ctx, http.StatusBadRequest, svc.ErrParamIsRequired("id",
			"queryParameter").Error()) //.Error() rertorna uma string que representa o erro
		return
	}
	opening := model.Opening{}

	// find opening   //pocura a opening no Db com id e carrega os dados na  estrutura opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	//delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	var null interface{}
	sendSuccess(ctx, "deleting-opening", null, http.StatusNoContent)
}
