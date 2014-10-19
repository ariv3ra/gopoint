package main

import (
	"encoding/json"
	"fmt"
	"os"
	// "net/http"

	// "github.com/franela/goreq"
	"github.com/jmcvetta/napping"
)

var baseURL string = "https://api.point.io/v2"

// User Object
type User struct {
	Email    string
	Password string
	APIkey   string
}

type result struct{}

func (u User) GetSessionKey(email, password, apiKey string) string {
	// Here we request the point.io sesionkey
	payload := User{email, password, apiKey}

	fmt.Println(payload)

	param := fmt.Sprintf("/auth?email=%s&password=%s&apikey=%s", email, password, apiKey)
	baseURL += param

	fmt.Println(baseURL)

	url := baseURL
	res := result{}
	resp, err := napping.Post(url, &payload, &res, nil)

	if err != nil {
		panic(err)
	}
	if resp.Status() == 200 {
		fmt.Println("We have a 200 OK------")
		fmt.Println(resp.RawText())
	}
	return resp.RawText()
}

func main() {

	var usr = User{}

	fmt.Println("Starting the Program ")

	configFile, err := os.Open("pio-creds.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&usr); err != nil {
		fmt.Println("parsing config file", err.Error())
	}

	user := User{}
	sk := user.GetSessionKey(usr.Email, usr.Password, usr.APIkey)
	fmt.Println("This is the returned json from the point.io api " + sk)
}
