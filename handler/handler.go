package handler

import (
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
