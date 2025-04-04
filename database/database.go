package database

import (
	"fmt"
	"time"

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

	db, err := sqlx.Open(database.Driver, dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(2 * time.Minute)

	return db
}
