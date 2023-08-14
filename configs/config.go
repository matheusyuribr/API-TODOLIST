package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIconfig
	DB  DBconfig
}

type APIconfig struct {
	Port string
}

type DBconfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")

}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIconfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBconfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Pass:     viper.GetString("db.pass"),
		Database: viper.GetString("db.database"),
	}
	return nil
}

func GetDB() DBconfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
