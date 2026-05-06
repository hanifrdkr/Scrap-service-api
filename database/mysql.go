package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"helicopter-hr/config"
	"time"
)

func InitSqlConnection(cfg *config.ConfigApp) (db *sqlx.DB, err error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.MysqlDB.Username,
		cfg.MysqlDB.Password,
		cfg.MysqlDB.HostName,
		cfg.MysqlDB.Port,
		cfg.MysqlDB.DatabaseName)

	db, err = sqlx.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MysqlDB.MaxIdleConnection)
	db.SetMaxOpenConns(cfg.MysqlDB.MaxOpenConnection)
	db.SetConnMaxLifetime(time.Duration(cfg.MysqlDB.MaxLifetimeConnection) * time.Second)

	return
}
