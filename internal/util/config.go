package config

import (
	"github.com/spf13/viper"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
)

// LoadConfig func
func LoadConfig() m.Config {
	c, _ := getConfigFromFile()
	return c
}

func getConfigFromFile() (m.Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	c := m.Config{}
	if err != nil { // Handle errors reading the config file
		return c, err
	}
	c.Server.ServerAddress = viper.GetString("server.addr")
	c.Server.ServerPort = viper.GetString("server.port")
	c.Server.Secret = viper.GetString("server.secret")
	c.DB.DBAddress = viper.GetString("db.addr")
	c.DB.DBName = viper.GetString("db.database")
	c.OBS.Endpoint = viper.GetString("storage.endpoint")
	c.OBS.AccessKeyID = viper.GetString("storage.accessKeyID")
	c.OBS.SecretAccessKey = viper.GetString("storage.secretAccessKey")
	c.OBS.Secure = viper.GetBool("storage.secure")
	c.OBS.BucketName = viper.GetString("storage.bucketName")
	c.OBS.Region = viper.GetString("storage.region")
	return c, nil
}
