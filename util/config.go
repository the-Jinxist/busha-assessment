package util

import (
	"github.com/spf13/viper"
)

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
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") //The file type could be json, xml and such

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

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
