package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/partnerhub24/konticket-serverless-libs/environment"
)

func InitDatabase(database environment.DatabaseEnvironment) *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		database.User,
		database.Pass,
		database.Host,
		database.Port,
		database.Name,
	)

	fmt.Println("dsn --->", dsn)

	db, err := sqlx.Open(database.Driver, dsn)
	if err != nil {
		panic(err)
	}

	return db
}
