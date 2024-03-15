package config

import (
	"os"

	"github.com/guilhermeFrBarros/gopportunities/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// checa se existe o arquivo
	// _  => descarta o valor
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		// Create the file exists
		err = os.MkdirAll("./db", os.ModePerm) // cria um diretorio no path e da permição para qualque usuario modificar
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath) // cria o arquivo
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// create DB and connect
	//                                                  // config padrão para
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{}) // abre uma conexão com banco de dados
	if err != nil {
		logger.ErrF("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate((&model.Opening{})) // pega a estrutura montado no schema e tenta migra o BD para ela, se não conseguir retorna um erro
	if err != nil {
		logger.ErrF("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
