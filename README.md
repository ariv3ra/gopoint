gopoint
===

Usage of point.io's apiDoc api in Go Lang

##Application configuration

Install the [napping](https://github.com/jmcvetta/napping) go [package](https://godoc.org/github.com/jmcvetta/napping)

`go get github.com/jmcvetta/napping`

Create a new file called:`pio-creds.json`

Add the following text to this file & replace the <> values with your information
```
{
	"email":"<your point.io email>",
	"password":"<your password>",
	"apikey":"<your apikey/appID>"
}
```