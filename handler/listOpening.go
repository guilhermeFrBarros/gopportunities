package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeFrBarros/gopportunities/model"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []model.Opening{}

	pageSize, offSet, err := calculatePageSizeAndOffsetAndValidQuery(ctx)
	if err != nil {
		return
	}

	if err := db.Limit(pageSize).Offset(offSet).Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "errot listing openings")
		return
	}

	sendSuccess(ctx, "list-opening", openings, http.StatusOK)

}

func calculatePageSizeAndOffsetAndValidQuery(ctx *gin.Context) (int, int, error) {
	// default
	pageSize := 20
	pageNumber := 1

	//valid parameter pageNumber
	if pN := ctx.Query("pageNumber"); pN != "" {
		var err error
		pageNumber, err = strconv.Atoi(pN)
		if err != nil || pageNumber < 1 {
			sendError(ctx, http.StatusBadRequest, "Invalid page number parameter")
			return 0, 0, fmt.Errorf("erro")
		}
	}

	//valid parameter pageSize
	if pS := ctx.Query("pageSize"); pS != "" {
		var err error
		pageSize, err = strconv.Atoi(pS)
		if err != nil || 1 > pageSize {
			sendError(ctx, http.StatusBadRequest, "Invalid page size parameter")
			return 0, 0, fmt.Errorf("erro")
		}
	}

	offSet := (pageNumber - 1) * pageSize

	return pageSize, offSet, nil
}
