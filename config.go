package learningdatabase

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	TableIndex = "index"
	TableQuery = "query"
)

type Config struct {
	DBNameIndex string `envconfig:"DB_NAME_INDEX" default:"moviedatabase"`
	DBNameQuery string `envconfig:"DB_NAME_QUERY" default:"ruangbook"`
	DBUser      string `envconfig:"DB_USERNAME" default:""`
	DBPass      string `envconfig:"DB_PASSWORD" default:""`
	DBHost      string `envconfig:"DB_HOST" default:""`
	DBPort      string `envconfig:"DB_PORT" default:"3306"`
}

func GetConfig() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}

func InitDB(table string) (*gorm.DB, error) {
	cfg := GetConfig()
	var db string
	if table == TableIndex {
		db = cfg.DBNameIndex
	} else if table == TableQuery {
		db = cfg.DBNameQuery
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, db)
	fmt.Printf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, db)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return conn, err
}
