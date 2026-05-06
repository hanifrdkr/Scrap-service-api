package config

import (
	"github.com/spf13/viper"
)

type ConfigApp struct {
	App        App       `mapstructure:"app"`
	MysqlDB    Database  `mapstructure:"mysql_db"`
	PostgresDB Database  `mapstructure:"postgres_db"`
	MongoDB    Database  `mapstructure:"mongo_db"`
	Log        LogConfig `mapstructure:"log"`
	FTP        FTPConfig `mapstructure:"ftp"`
}

type App struct {
	Debug     bool   `mapstructure:"debug"`
	Port      string `mapstructure:"port"`
	Timeout   int    `mapstructure:"timeout"`
	SecretKey string `mapstructure:"secret_key"`
}

type Database struct {
	HostName              string `mapstructure:"hostname"`
	Port                  string `mapstructure:"port"`
	Username              string `mapstructure:"username"`
	Password              string `mapstructure:"password"`
	DatabaseName          string `mapstructure:"database_name"`
	MaxIdleConnection     int    `mapstructure:"max_idle_connection"`
	MaxOpenConnection     int    `mapstructure:"max_open_connection"`
	MaxLifetimeConnection int    `mapstructure:"max_lifetime_connection"`
}

type LogConfig struct {
	Dir       string `mapstructure:"dir"`
	FileName  string `mapstructure:"file_name"`
	MaxSize   int    `mapstructure:"max_size"`
	LocalTime bool   `mapstructure:"local_time"`
	Compress  bool   `mapstructure:"compress"`
}

type FTPConfig struct {
	HostName  string `mapstructure:"hostname"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Directory string `mapstructure:"directory"`
}

type Recaptcha struct {
	SiteVerifyURL string `mapstructure:"site_verify_url"`
	Secret        string `mapstructure:"secret"`
}

func LoadConfig(pathFile string) (*ConfigApp, error) {
	var config ConfigApp
	viper.SetConfigType("yaml")
	viper.SetConfigFile(pathFile)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
