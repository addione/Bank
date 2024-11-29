package dependencies

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type commonMysql struct {
	uri string
}

func newMysql() *commonMysql {
	godotenv.Load(".env")
	uri := os.Getenv("mysqluri")

	return &commonMysql{
		uri: uri,
	}
}

func (m *commonMysql) getMysqlClient(dbName string) *sql.DB {
	db, err := sql.Open("mysql", m.uri+"/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
