package student

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gustavocoimbradev/rm-cli/internal/terminal"
)

func valueOrDash(value sql.NullString) string {
	if !value.Valid || value.String == "" {
		return "-"
	}
	return value.String
}

func Get(db *sql.DB, ra string) {
	rows, err := db.Query(`
		SELECT TOP 1
			SALUNO.RA AS ra,
			PPESSOA.NOME AS name,
			PPESSOA.CPF AS cpf,
			CONVERT(VARCHAR(10), PPESSOA.DTNASCIMENTO, 23) AS birth,
			PPESSOA.EMAIL AS email,

			PPESSOA.TELEFONE1 AS phone1,
			PPESSOA.TELEFONE2 AS phone2,
			PPESSOA.TELEFONE3 AS phone3,

			PPESSOA.RUA AS street,
			PPESSOA.NUMERO AS number,
			PPESSOA.COMPLEMENTO AS complement,
			PPESSOA.BAIRRO AS district,
			PPESSOA.ESTADO AS state,
			PPESSOA.CIDADE AS city,
			PPESSOA.CEP AS zip_code,
			PPESSOA.PAIS AS country

		FROM SALUNO
		INNER JOIN PPESSOA
			ON PPESSOA.CODIGO = SALUNO.CODPESSOA

		WHERE SALUNO.RA = @p1
	`, ra)
	if err != nil {
		panic("Erro ao consultar alunos: " + err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var studentRA string
		var name string

		var cpf sql.NullString
		var birth sql.NullString
		var email sql.NullString

		var phone1 sql.NullString
		var phone2 sql.NullString
		var phone3 sql.NullString

		var street sql.NullString
		var number sql.NullString
		var complement sql.NullString
		var district sql.NullString
		var state sql.NullString
		var city sql.NullString
		var zipCode sql.NullString
		var country sql.NullString

		err := rows.Scan(
			&studentRA,
			&name,
			&cpf,
			&birth,
			&email,
			&phone1,
			&phone2,
			&phone3,
			&street,
			&number,
			&complement,
			&district,
			&state,
			&city,
			&zipCode,
			&country,
		)

		phones := make([]string, 0, 3)

		for _, phone := range []sql.NullString{phone1, phone2, phone3} {
			if phone.Valid && strings.TrimSpace(phone.String) != "" {
				phones = append(phones, phone.String)
			}
		}

		if err != nil {
			panic("Erro ao ler aluno: " + err.Error())
		}
		fmt.Printf(
			terminal.Yellow+"\nDADOS DO ALUNO\n\n"+terminal.Reset+
				terminal.Gray+"RA:"+terminal.BrightWhite+" %s\n"+terminal.Reset+
				terminal.Gray+"Nome:"+terminal.BrightWhite+" %s\n"+terminal.Reset+
				terminal.Gray+"CPF:"+terminal.BrightWhite+" %s\n"+terminal.Reset+
				terminal.Gray+"Nascimento:"+terminal.BrightWhite+" %s\n"+terminal.Reset+
				terminal.Gray+"E-mail:"+terminal.BrightWhite+" %s\n"+terminal.Reset+
				terminal.Gray+"Telefones:"+terminal.BrightWhite+" %s\n\n"+terminal.Reset+

				terminal.Yellow+"ENDEREÇO\n\n"+terminal.Reset+

				terminal.Gray+"Rua:"+terminal.BrightWhite+" %s\n"+terminal.Reset+
				terminal.Gray+"Número:"+terminal.BrightWhite+" %s\n"+terminal.Reset+
				terminal.Gray+"Complemento:"+terminal.BrightWhite+" %s\n"+terminal.Reset+
				terminal.Gray+"Bairro:"+terminal.BrightWhite+" %s\n"+terminal.Reset+
				terminal.Gray+"Cidade:"+terminal.BrightWhite+" %s\n"+terminal.Reset+
				terminal.Gray+"Estado:"+terminal.BrightWhite+" %s\n"+terminal.Reset+
				terminal.Gray+"CEP:"+terminal.BrightWhite+" %s\n"+terminal.Reset+
				terminal.Gray+"País:"+terminal.BrightWhite+" %s\n\n"+terminal.Reset,
			studentRA,
			name,
			valueOrDash(cpf),
			valueOrDash(birth),
			valueOrDash(email),
			strings.Join(phones, ", "),
			valueOrDash(street),
			valueOrDash(number),
			valueOrDash(complement),
			valueOrDash(district),
			valueOrDash(city),
			valueOrDash(state),
			valueOrDash(zipCode),
			valueOrDash(country),
		)
	}
}
