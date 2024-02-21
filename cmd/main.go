package main

import (
	"github.com/guilhermeFrBarros/gopportunities/config"
	"github.com/guilhermeFrBarros/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	//initialize Configs
	err := config.Init()
	if err != nil {
		logger.ErrF("config initialization erro: %v", err)
		return
	}

	router.Initialize()
}
