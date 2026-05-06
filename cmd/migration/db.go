// Package migration
package migration

import (
	"helicopter-hr/config"
	"strconv"
	"time"

	"helicopter-hr/pkg/databasex"
)

func MigrateDatabase() {
	cfg, _ := config.LoadConfig("./config.yaml")

	port, _ := strconv.Atoi(cfg.PostgresDB.Port)

	databasex.DatabaseMigration(&databasex.Config{
		Driver:       "postgres",
		Host:         cfg.PostgresDB.HostName,
		Port:         port,
		Name:         cfg.PostgresDB.DatabaseName,
		User:         cfg.PostgresDB.Username,
		Password:     cfg.PostgresDB.Password,
		MaxIdleConns: cfg.PostgresDB.MaxIdleConnection,
		MaxOpenConns: cfg.PostgresDB.MaxOpenConnection,
		MaxLifetime:  time.Duration(cfg.PostgresDB.MaxLifetimeConnection) * time.Millisecond,
		Charset:      "UTF-8",
		Timeout:      10000000000,
		TimeZone:     "Asia/Jakarta",
	})
}
