package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	DbUserName string `json:"dbUserName"`
	DbPassword string `json:"dbPassword"`
	DbName     string `json:"dbName"`
	DbIP       string `json:"dbIp"`
}

func Parse() Config {
	// JsonConfig, err := os.Open("config/config.json")
	JsonConfig, err := os.Open("config.json")
	config := Config{}
	//if config failed, by default use test mod
	if err != nil {

		config = Config{
			DbUserName: "root",
			DbPassword: "12345678",
			DbName:     "goleader",
			DbIP:       "127.0.0.1:3306",
		}
	} else {
		ByteData, _ := ioutil.ReadAll(JsonConfig)
		err = json.Unmarshal(ByteData, &config)
		if err != nil {
			panic(err)
		}
	}
	return config
}
