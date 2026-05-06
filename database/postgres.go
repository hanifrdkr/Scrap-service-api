package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"helicopter-hr/config"
	"time"
)

func InitPostgreSqlConnection(cfg *config.ConfigApp) (db *sql.DB, err error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.PostgresDB.HostName,
		cfg.PostgresDB.Port,
		cfg.PostgresDB.Username,
		cfg.PostgresDB.Password,
		cfg.PostgresDB.DatabaseName)

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(cfg.PostgresDB.MaxIdleConnection)
	db.SetMaxOpenConns(cfg.PostgresDB.MaxOpenConnection)
	db.SetConnMaxLifetime(time.Duration(cfg.MysqlDB.MaxLifetimeConnection) * time.Second)

	return
}
