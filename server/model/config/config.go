package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var Config *config

// var cfgFile = "config/app.json"

type config struct {
	HttpPort int      `json:"HttpPort"`
	TcpPort  int      `json:"TcpPort"`
	MySQL    MySQLCfg `json:"MySQL"`
	MongoDB  MongoCfg `json:"MongoDB"`
	Redis    RedisCfg `json:"Redis"`
}

func NewConfig() {
	// Open config.json
	file, err := os.Open("config/app.json")
	if err != nil {
		log.Panic("Cannot read app.json:", err)
	}

	// Read config.json as a byte array
	cfgByte, _ := ioutil.ReadAll(file)

	// Parse json to struct
	json.Unmarshal(cfgByte, &Config)
}

type MySQLCfg struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Database string `json:"Database"`
}

type MongoCfg struct {
	URL      string `json:"URL"`
	Database string `json:"Database"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type RedisCfg struct {
	Size     int    `json:"Size"`
	Network  string `json:"Network"`
	Address  string `json:"Address"`
	Password string `json:"Password"`
}
