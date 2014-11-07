package users

import (
	"fmt"
	"github.com/jmcvetta/napping"
)

// User Object
type User struct {
	BaseURL  string
	Email    string
	Password string
	APIkey   string
}

type result struct{}

func (u User) Authenticate(baseurl, email, password, apikey string) string {
	//   TODO Fill this out later maybe
	return ""
}

func (u User) GetSessionKey(baseurl, email, password, apiKey string) (string, string, string) {
	// Here we request the point.io sesionkey

	// SessionKey Variable
	var sessionKey string

	// Assign the payload variable
	payload := User{baseurl, email, password, apiKey}

	// Create the URL request
	url := fmt.Sprintf("%s/auth?email=%s&password=%s&apikey=%s", baseurl, email, password, apiKey)

	// Create Data structs

	//Create the ouptput type from the ppoint.io api
	type Result struct {
		SESSIONKEY string
		FNAME      string
		LNAME      string
	}

	// Create the Response wrapper for the data response
	type Response struct {
		RESULT  Result
		ERROR   float32
		MESSAGE string
	}

	// New instance of the Response Struct
	var res = Response{}

	// Create & send the napping request & get assign the response to the resp var
	resp, err := napping.Post(url, &payload, &res, nil)

	if err != nil {
		panic(err)
	}
	if resp.Status() == 200 {

		if res.ERROR > 0 {
			sessionKey = "Error: " + res.MESSAGE
		} else {
			// Serve Up the SessionKey
			sessionKey = res.RESULT.SESSIONKEY
		}
	}
	return sessionKey, res.RESULT.FNAME, res.RESULT.LNAME
}
