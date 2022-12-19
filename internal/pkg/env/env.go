package env

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	ENV                   string `mapstructure:"ENV"`
	ServiceName           string `mapstructure:"SERVICE_NAME"`
	ServicePort           string `mapstructure:"SERVICE_PORT"`
	ServiceVersion        string `mapstructure:"SERVICE_VERSION"`
	ServiceURL            string `mapstructure:"SERVICE_URL"`
	DBUsername            string `mapstructure:"DB_USERNAME"`
	DBPassword            string `mapstructure:"DB_PASSWORD"`
	DBHost                string `mapstructure:"DB_HOST"`
	DBName                string `mapstructure:"DB_NAME"`
	DBPort                string `mapstructure:"DB_PORT"`
	DBMaxIdleConn         int    `mapstructure:"DB_MAX_IDLE_CONN"`
	DBMaxOpenConn         int    `mapstructure:"DB_MAX_OPEN_CONN"`
	DBMaxLifeTimeConn     int    `mapstructure:"DB_MAX_TTL_CONN"`
	HubServiceURL         string `mapstructure:"HUB_SERVICE_URL"`
	TransactionServiceURL string `mapstructure:"TRANSACTION_SERVICE_URL"`
	JWTSecret             string `mapstructure:"JWT_SECRET"`
	StaticToken           string `mapstructure:"STATIC_TOKEN"`
}

var (
	cfg  *Config
	once sync.Once
)

func Get() *Config {
	if strings.HasSuffix(os.Args[0], ".test") || flag.Lookup("test.v") != nil {
		return &Config{}
	}

	return cfg
}

func Load() {
	once.Do(func() {
		v := viper.New()
		v.AutomaticEnv()

		v.AddConfigPath(".")
		v.SetConfigType("env")
		v.SetConfigName(".env")

		err := v.ReadInConfig()
		if err != nil {
			fmt.Println("config file not found: ", err)
		}

		config := new(Config)
		err = v.Unmarshal(config)
		if err != nil {
			panic(err)
		}

		cfg = config
	})
}
