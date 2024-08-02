package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	DebugMode    string = "debug"
	TestMode     string = "test"
	ReleasebMode string = "release"
)

type GeneralConfig struct {
	AppName     string
	Environment string
	Version     string
	HTTPPort    string
	HTTPScheme  string
	SignInKey   string
}

type PgConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

type RedisConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DbIndex  int
}

type Config struct {
	GeneralConfig GeneralConfig
	PgConfig      PgConfig
	RedisConfig   RedisConfig
}

func NewConfig() Config {
	return Config{GeneralConfig: GeneralConfig{AppName: "my_blog_backend"}}
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("env file is not found !")
	}

	var config = NewConfig()

	// general config
	config.GeneralConfig.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.GeneralConfig.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.8"))
	config.GeneralConfig.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
	config.GeneralConfig.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))
	config.GeneralConfig.SignInKey = cast.ToString(getOrReturnDefaultValue("SIGN_IN_KEY", "wedrfgtyhujikol"))

	// postgres config
	config.PgConfig.Username = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "superuser"))
	config.PgConfig.Password = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "superuser"))
	config.PgConfig.Host = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PgConfig.Port = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PgConfig.Database = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", config.GeneralConfig.AppName))

	// redis config
	config.RedisConfig.Host = cast.ToString(getOrReturnDefaultValue("REDIS_HOST", "localhost"))
	config.RedisConfig.Port = cast.ToInt(getOrReturnDefaultValue("REDIS_PORT", 6379))
	config.RedisConfig.DbIndex = cast.ToInt(getOrReturnDefaultValue("REDIS_DB_INDEX", 0))

	return config

}

// env fileni ishga tushirish
func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
