package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	SQLHost     = "mysqldb"
	SQLPort     = "3306"
	SQLUser     = "user"
	SQLPass     = "pass"
	SQLDatabase = "bitbybit"
)

//Connect: Inicializa a conexão com a base de dados e retorna uma struct do tipo DB, que tem os valores para manipulação do banco de dados
func Connect_SQL() (*sql.DB, error) {

	var (
		dataSourceName = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			SQLUser,
			SQLPass,
			SQLHost,
			SQLPort,
			SQLDatabase,
		)
	)

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
