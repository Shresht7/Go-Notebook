package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	//	Parse flags
	method := flag.String("method", "GET", "HTTP Method")
	showHeaders := flag.Bool("show-headers", false, "Show HTTP headers")
	body := flag.String("body", "", "HTTP request body")
	contentType := flag.String("content-type", "application/json", "HTTP Header Content-Type")
	flag.Parse()

	//	Exit if no URL was provided
	if len(flag.Args()) < 1 {
		log.Fatalln("Please provide a URL")
		os.Exit(1)
	}

	//	Get URL from parsed arguments
	url := flag.Arg(0)

	switch strings.ToUpper(*method) {

	case "GET":
		//	Get the URL
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		showResponse(*response, url, *method, *showHeaders)

	case "POST":

		//	Create JSON
		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(*body), &jsonMap)
		postBody, _ := json.Marshal(jsonMap)

		//	Post to URL
		response, err := http.Post(url, *contentType, bytes.NewBuffer(postBody))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		showResponse(*response, url, *method, *showHeaders)
	}
}

func showResponse(response http.Response, url string, method string, showHeaders bool) {

	//	Read response body
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	responseBody := string(bodyBytes)

	//	Print output
	fmt.Println("\n", method, url, response.Status)
	if showHeaders {
		fmt.Println()
		fmt.Println(response.Header)
		fmt.Println()
	}
	fmt.Println("\n" + responseBody + "\n")

}
