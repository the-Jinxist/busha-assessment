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

//Make sure to connect both redis(redis) and postgres(busha-assessment) containers in the same network
//docker network create busha-network
//docker network connect busha-network redis
//docker network connect busha-network assessment-image

//The access these services via their registered IP address on the network
//The ensuing command will look like this:

//busha-assessment % docker run --name neo-swapi --network busha-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@assessment-image:5432/comments_db?sslmode=disable" -e REDIS_ADDRESS="redis:6379" busha-assessment:latest
//in this command assessment-image and redis are the corresponding IP address and we're changing the environmental variables so Viper can retrived the updated
//values
