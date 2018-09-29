package config

import (
	"encoding/json"
	"log"
	"os"
)

const (
	BetaAppId_Accounts   string = "f53cc48ad2af7e1c23debfbb825481c1a3c1933534284641c3031876bf58b29c"
	BetaAppId_Chalmersit string = "eb8d1a954c189a88df9134c1dd3ba668dab04fd8ece2bbf51a873c6aac66382f"
	BetaAppId_Hubbit     string = "4e43fe76c887b3bfe4041d53269ceb2a5947d61443af39c6161aa8cece226eaa"
	BetaAppId_Bookit     string = "57d45947cbd1fc3090686cdb1a37e0e3a473e858c10ab54906a7a27a8417511b"
)

var BetaAppSecrets struct {
	Accounts   string `json:"account_client_secret"`
	Chalmersit string `json:"chalmersit_client_secret"`
	Hubbit     string `json:"hubbit_client_secret"`
	Bookit     string `json:"bookit_client_secret"`
}

func LoadSecrets() {
	configFile, err := os.Open("config/secrets.json")
	if err != nil {
		log.Fatal("Can't find secrets file!", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&BetaAppSecrets); err != nil {
		log.Fatal("Can't parse secrets file", err.Error())
	}
}
