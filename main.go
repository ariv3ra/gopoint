package main

import (
	"fmt"
	"github.com/ariv3ra/gopoint/users"
	"github.com/ariv3ra/gopoint/utils"
)

var baseURL, email, password, apikey, sessionKey string
var util = utils.Utils{}

func main() {

	fmt.Println("Starting the Program")

	// Create instance of new User to hold the cred values from the pio-creds.json
	var usr = users.User{}

	baseURL, email, password, apikey = util.GetLoginCreds("pio-creds.json")

	sk, fname, lname := usr.GetSessionKey(baseURL, email, password, apikey)

	// Assign the sessionkey to the global variable
	sessionKey = sk

	msg := fmt.Sprintf("Auth for user: %s \nSession Key: %s\nFname: %s\nLName: %s", email, sk, fname, lname)
	fmt.Println(msg)
}
