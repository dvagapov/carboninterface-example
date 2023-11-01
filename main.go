package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// helm message
var help_message = `
Usage:
  carboninterface-example <apikey> [-url <url>] [-method <method>] [-body <body>] [-header <header>]

Check API-KEY:
  ./carboninterface-example -apikey YOUR_API_KEY -url "https://www.carboninterface.com/api/v1/auth" -method GET

Variables:
  -apikey	The API key for 'Authorization: Bearer API_KEY'
  -url		The URL of the carboninterface API. Default is https://www.carboninterface.com/api/v1/estimates
  -method	The method of the API request. Default is POST
  -body		The body of the API request. Default is empty JSON "{}"
  -ctx		The Content-Type of the API request. Default is "application/json"
  -timeout	The timeout seconds of the API request. Default is 30 seconds
  -help		Show help message. Default is false

Example:
  carboninterface-example -apikey XXX -body '{"type": "flight","passengers":2,"legs":[{"departure_airport":"sfo","destination_airport":"yyz"},{"departure_airport":"yyz","destination_airport":"sfo"}]}'
`
// define the variables for the arguments and flags
var url string
var method string
var body string
var apikey string
var ctx string
var timeout int
var help bool

// init function to parse the arguments and flags
func init() {
	flag.StringVar(&url, "url", "https://www.carboninterface.com/api/v1/estimates", "The URL of the carboninterface API")
	flag.StringVar(&apikey, "apikey", "", "The API key for 'Authorization: Bearer API_KEY'")
	flag.StringVar(&method, "method", "POST", "The method of the API request")
	flag.StringVar(&body, "body", "{}", "The body of the API request")
	flag.StringVar(&ctx, "ctx", "application/json", "The Content-Type of the API request")
	flag.IntVar(&timeout, "timeout", 30, "The timeout seconds of the API request")
	flag.BoolVar(&help, "help", false, "Show help message")
	flag.Parse()
	}

func main() {

	// check URL || help
	if help || apikey == "" || (body == "{}" && method != "GET") {
		// show usage and exit
		fmt.Println(help_message)
		os.Exit(0)
	}

  // create http client
  client := &http.Client {Timeout: time.Second * time.Duration(timeout) }
	
  req, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))

  if err != nil {
    fmt.Println(err)
    return
  }
  
	// set header
	req.Header.Set("Content-Type", ctx)
	// add API_KEY to header if it's not empty
	if apikey != "" {
		req.Header.Set("Authorization", "Bearer " + apikey)
	}

	// send the request
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
	// close the body to free resources
  defer res.Body.Close()

	// read body
  body, err := io.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}