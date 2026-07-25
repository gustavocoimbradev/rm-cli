package app

import (
	"database/sql"
	"fmt"

	"github.com/gustavocoimbradev/rm-cli/internal/config"
	"github.com/gustavocoimbradev/rm-cli/internal/database"
)

type App struct {
	DB *sql.DB
}

func New() *App {
	config.LoadEnv()

	db := database.Connect()

	return &App{
		DB: db,
	}
}

func (app *App) Close() {
	if app.DB == nil {
		return
	}

	if err := app.DB.Close(); err != nil {
		fmt.Println("Erro ao fechar conexão com o banco:", err)
	}
}
