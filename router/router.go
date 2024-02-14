package router // todas as pasta do go são sub package
// Por convensão se eu tenho uma pasta router eu tenho um arquivo router.go, arquivo com o mesmo nome da pasta

import "github.com/gin-gonic/gin" // no import se vc não dar um nome, pega se o ultimo nome

func Initialize() {
	router := gin.Default()                      // muito comum usar ns linguagem so a primeira letra
	router.GET("/ping", func(ctx *gin.Context) { // Context seria a cmbinação do request e do response
		ctx.JSON(200, gin.H{ // um mapa  para criar uma estruta chave valor similar aos objetos do JS
			"message": "pong",
		})
	})

	router.Run(":8080") // padrao 8080
}
