package database

import (
	"database/sql"
	"net"
	"net/url"
	"os"
)

func getConnectionURL() url.URL {
	params := url.Values{}
	params.Set("database", os.Getenv("DB_DATABASE"))
	return url.URL{
		Scheme: "sqlserver",
		User: url.UserPassword(
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
		),
		Host: net.JoinHostPort(
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
		),
		RawQuery: params.Encode(),
	}
}

func Connect() *sql.DB {
	connectionURL := getConnectionURL()
	db, err := sql.Open("sqlserver", connectionURL.String())
	if err != nil {
		panic("Erro ao preparar conexão " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic("Falha ao se conectar ao banco de dados " + err.Error())
	}
	return db
}
