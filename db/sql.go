package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

const (
	SQLHost     = "172.17.223.65"
	SQLPort     = "1433"
	SQLUser     = "SA"
	SQLPass     = "DockerSql!"
	SQLDatabase = "db_bitbybit"
)

//Connect: Inicializa a conexão com a base de dados e retorna uma struct do tipo DB, que tem os valores para manipulação do banco de dados
func Connect_SQL() error {

	var (
		dataSourceName = fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s",
			SQLHost, SQLPort, SQLUser, SQLPass, SQLDatabase)
	)

	db, err := sql.Open("mssql", dataSourceName)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	DB = db

	return nil
}
