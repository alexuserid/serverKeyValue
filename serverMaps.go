package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var data = make(map[string]string)

func handler(w http.ResponseWriter, r *http.Request) {

	validKV := regexp.MustCompile("/[0-9]{1,10}:[a-zA-Z]{1,15}/")
	reKey := regexp.MustCompile("/[0-9]{1,10}:")
	reValue := regexp.MustCompile("[a-zA-Z]{1,15}")
	reKeyNums := regexp.MustCompile("[0-9]{1,10}")

	key := reKeyNums.FindString(r.URL.Path)

	if r.URL.Path[1:] == "" {
		//sending text when url in address bar is like site.com/
		fmt.Fprintf(w, "Enter key and value in form: sitename.com/key:value/. \nA key must contain 1-10 integers, a value must contain 1-15 strings. \nIf a pair key:value is already exis, you can get the value by key, enter sitename.com/key:/")

	} else if validKV.FindString(r.URL.Path) != "" { //if key and value are both correct
		nums, ok := data[key] //maps return two values. If key exist second value will be true
		data[key] = reValue.FindString(r.URL.Path) //in any case new value will be written with this key
		if ok {
			fmt.Fprintf(w, "The value for key %s was rewrited", nums) //if key already exist
		} else {
			fmt.Fprintf(w, "The pair was added")
		}
		
	} else if reKey.FindString(r.URL.Path) != "" { //if path include only key like /ints:
		_, ok := data[key]
		if ok {
			fmt.Fprintf(w, "The value for key %s is %s", key, data[key])
		} else {
			fmt.Fprintf(w, "There's no value for this key") //text if value isnt present
		}
		
	} else {
		fmt.Fprintf(w, "The request is incorrect") //text if request is incorrect
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
