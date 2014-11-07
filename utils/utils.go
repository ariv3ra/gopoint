package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ariv3ra/gopoint/users"
	"os"
)

type Utils struct {
}

func (u Utils) GetLoginCreds(file string) (string, string, string, string) {

	var usr = users.User{}

	// Open the creds file
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&usr); err != nil {
		fmt.Println("parsing config file", err.Error())
	}

	return usr.BaseURL, usr.Email, usr.Password, usr.APIkey

}
