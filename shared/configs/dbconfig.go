package configs

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfig struct {
	DatabaseName string `envconfig:"DB_NAME"`
	User         string `envconfig:"DB_USER"`
	Password     string `envconfig:"DB_PASSWORD"`
	IpAddress    string `envconfig:"DB_IP"`
	Port         string `envconfig:"DB_PORT"`
	Charset      string `envconfig:"DB_CHARSET"`
	ParseTime    string `envconfig:"DB_PARSETIME"`
	Loc          string `envconfig:"DB_LOC"`
}

func CreateDbConfig() (*DatabaseConfig, error) {
	dbConfig := &DatabaseConfig{}
	err := envconfig.Process("", dbConfig)
	if err != nil {
		return nil, err
	}

	return dbConfig, nil
}

func (dbConfig DatabaseConfig) GetDsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", dbConfig.User, dbConfig.Password,
		dbConfig.IpAddress, dbConfig.Port, dbConfig.DatabaseName, dbConfig.Charset, dbConfig.ParseTime, dbConfig.Loc)

	return dsn
}
