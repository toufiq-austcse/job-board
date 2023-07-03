package ent

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/toufiq-austcse/go-api-boilerplate/config"
	"github.com/toufiq-austcse/go-api-boilerplate/ent"
	"github.com/toufiq-austcse/go-api-boilerplate/enums/db_driver"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/db/providers/mysql"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/db/providers/postgresql"
)

func New() (*ent.Client, error) {
	dbDialect, db, err := getDbProvider()
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dbDialect, db)
	client := ent.NewClient(ent.Driver(drv))
	if config.AppConfig.DB_CONFIG.DEBUG_ENABLED == "true" {
		return client.Debug(), nil
	}
	return client, nil

}

func getDbProvider() (dbDialect string, db *sql.DB, err error) {
	if config.AppConfig.DB_DRIVER_NAME == string(db_drivers_enum.POSTGRESQL) {
		db, err := postgresql.OpenPgDbConnection()
		return dialect.Postgres, db, err
	}
	if config.AppConfig.DB_DRIVER_NAME == string(db_drivers_enum.MYSQL) {
		db, err := mysql.OpenMySqlDbConnection()
		return dialect.MySQL, db, err
	}
	return "", nil, errors.New("UnSupported DB Driver")
}
