package main

import (
	"net/http"
)

type Client struct {
	url         string
	contentType string
}

var (
	url      = "http://localhost"
	Frontend = "http://140.124.183.37/home"
	Backend  = "http://140.124.183.37/"
	port     = "80"
)

func init() {}

func main() {
	//handle
	//&LAT=%f&LNG=%f&KEYWORD=%S", APIKey, Lat, Lng, keyword,
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", http.StripPrefix("/lib/", http.FileServer(http.Dir("templates"))))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("Connect Fail:" + err.Error())
	}

}
