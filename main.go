package main

import (
	"fmt"
	"os"

	"github.com/gustavocoimbradev/rm-cli/internal/app"
	"github.com/gustavocoimbradev/rm-cli/internal/student"
	"github.com/gustavocoimbradev/rm-cli/internal/terminal"
	_ "github.com/microsoft/go-mssqldb"
)

func main() {

	application := app.New()
	defer application.Close()

	if len(os.Args) < 2 {
		fmt.Println(terminal.Yellow + "LISTA DE COMANDOS DISPONÍVEIS:" + terminal.Reset)
		fmt.Println(terminal.BrightWhite + "./rm-cli student {{ra}}" + terminal.Cyan + " - Retorna os dados básicos do aluno" + terminal.Reset)
		return
	}

	if os.Args[1] == "student" {

		if len(os.Args) < 3 {
			panic("Informe o RA do aluno")
		}

		student.Get(application.DB, os.Args[2])

	}

}
