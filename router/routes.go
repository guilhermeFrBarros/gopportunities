package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermeFrBarros/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitalizeHandler() // iniciliza as vairvaies globais do handler

	rV1 := router.Group("/api/v1")
	{
		rV1.GET("/opening", handler.GetOpeningHandler)

		rV1.POST("/opening", handler.CreateOpeningHandler)

		rV1.DELETE("/opening", handler.DeleteOpeningHandler)
		//DELETE("/opening/:id", func)"  com param

		rV1.PUT("/opening", handler.UpdateOpeningHandler)

		rV1.GET("/openingAll", handler.ListOpeningsHandler)
	}

}
