package main

import (
	"encoding/json"
	"fmt"
	"github.com/ariv3ra/gopoint/users"
	"os"
)

func main() {

	fmt.Println("Starting the Program ")

	// Create instance of new User
	var usr = users.User{}

	// Open the creds file
	configFile, err := os.Open("pio-creds.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&usr); err != nil {
		fmt.Println("parsing config file", err.Error())
	}

	user := users.User{}
	sk := user.GetSessionKey(usr.BaseURL, usr.Email, usr.Password, usr.APIkey)
	// fmt.Println("This is the returned json from the point.io api " + sk)
	msg := fmt.Sprintf("Auth for user: %s \nSession Key: %s", usr.Email, sk)
	fmt.Println(msg)
}
