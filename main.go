package main

import (
	"encoding/json"
	"fmt"
	"github.com/ariv3ra/gopoint/users"
	"os"
)


// app level sessionKey

var sessionKey string = ""

func main() {

	fmt.Println("Starting the Program")

	// Create instance of new User to hold the cred values from the pio-creds.json
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
	sk, fname, lname := user.GetSessionKey(usr.BaseURL, usr.Email, usr.Password, usr.APIkey)

	// Assign the sessionkey to the global variable
	sessionKey = sk

	msg := fmt.Sprintf("Auth for user: %s \nSession Key: %s\nFname: %s\nLName: %s", usr.Email, sk, fname, lname)
	fmt.Println(msg)
}
