package main

import (
	"fmt"
	"net/url"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			return
		}
	}()

	/********** Parsing URL **********/

	rawUrl := "https://www.google.com:8080/path?queryName=TestName&queryValue=TestValue#id"

	parsedUrl, err := url.Parse(rawUrl)

	if err != nil {
		panic(fmt.Sprintf("Error parsing URL: %v", err))
	}
	fmt.Println("URL Scheme:", parsedUrl.Scheme)
	fmt.Println("URL Host:", parsedUrl.Host)
	fmt.Println("URL Port:", parsedUrl.Port())
	fmt.Println("URL Hostname:", parsedUrl.Hostname())
	fmt.Println("URL Path:", parsedUrl.Path)
	fmt.Println("URL Raw Query:", parsedUrl.RawQuery)
	fmt.Println("URL Fragmwnt:", parsedUrl.Fragment)

	query := parsedUrl.Query()

	fmt.Println(query.Get("queryName"))

	/********** Building URL **********/

	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "www.google.com",
		Path:   "new-path",
	}

	baseUrlQuery := baseUrl.Query()
	baseUrlQuery.Set("name", "Test")

	baseUrl.RawQuery = baseUrlQuery.Encode()
	baseUrl.Fragment = "frag"
	fmt.Println(baseUrl) //OUTPUT: https://www.google.com/new-path?name=Test#frag

	/********** Encoding Query **********/
	queryValues := url.Values{}

	//Adding key value
	queryValues.Add("name", "Rohit")
	queryValues.Add("age", "28")
	queryValues.Add("unusedVal", "28")

	//Checking is value exists for specific key and replacing value
	if queryValues.Has("age") {
		queryValues.Set("age", "29")
	}

	//Deleting value
	queryValues.Del("unusedVal")

	//Encoding Value
	fmt.Println(queryValues.Encode()) // OUTPUT: age=29&name=Rohit

}
