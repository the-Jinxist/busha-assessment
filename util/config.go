package util

import "os"

// Config stores all the configurations of an application
// The values are read by viper from a config file or environment variable

// We're storing our three environmental variables
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	RedisAddress  string `mapstructure:"REDIS_ADDRESS"`
}

// This function will load environmental variables from file or environmental variables
func LoadConfig(path string) (config Config, err error) {
	config = Config{}

	config.DBDriver = "postgres"
	config.DBSource = os.Getenv("DATABASE_URL")
	config.ServerAddress = os.Getenv("PORT")
	config.RedisAddress = os.Getenv("REDIS_URL")

	return
}
