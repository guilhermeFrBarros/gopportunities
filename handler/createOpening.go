package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeFrBarros/gopportunities/model"
)

func CreateOpeningHandler(ctx *gin.Context) {

	request := model.CreateOpeningDTO{} // {} instanciação

	ctx.BindJSON(&request) // analisa o dados da requisição e preenche o objeto dentro dos parentes

	if err := request.Validate(); err != nil {
		logger.ErrF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := model.Opening{}
	//func que passa os valores para da req para a opening
	opening.NewOpening(&request)
	// opening := model.Opening{
	// 	Role:     request.Role,
	// 	Company:  request.Company,
	// 	Location: request.Location,
	// 	Remote:   *request.Remote,
	// 	Link:     request.Link,
	// 	Salary:   request.Salary,
	// }

	if err := db.Create((&opening)).Error; err != nil {
		logger.ErrF("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on data")
		return
	}

	sendSuccess(ctx, "create-opening", opening, http.StatusCreated) // devido ao ponteiro oi gorm atualiza as informações da opening
}

/*
struct {
		Role string `json: "role"` // role com a 1 letra minuscula  não exportado para o console, ou seja, vc não consegue imprimir o valor
	}{} // {} apos é uma instaciação
*/
