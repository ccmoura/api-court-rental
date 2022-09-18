package main
import (
	"log"
	"github.com/spf13/viper"
)
type Config struct {
	Port             string `mapstructure:"PORT"`
	PgConnectionString string `mapstructure:"PG_CONNECTION_STRING"`
}

var AppConfig *Config

func LoadAppConfig(){
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
