package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type KV struct {
	Key   string
	Value string
}

var data = []KV{}

func handler(w http.ResponseWriter, r *http.Request) {

	validKV := regexp.MustCompile("/[0-9]{1,10}:[a-zA-Z]{1,15}/")
	reKey := regexp.MustCompile("/[0-9]{1,10}:")
	reValue := regexp.MustCompile("[a-zA-Z]{1,15}")
	reKeyNums := regexp.MustCompile("[0-9]{1,10}")

	key := reKeyNums.FindString(r.URL.Path)

	if r.URL.Path[1:] == "" {
		//sending text when address bar is like site.com/
		fmt.Fprintf(w, "Enter key and value in form: sitename.com/key:value/. \nA key must compile 1-10 integers, a value must compile 1-15 strings. \nIf a pair key:value is already exis, you can get the value by key, enter sitename.com/key:/")

	} else if validKV.FindString(r.URL.Path) != "" { //if key and value are both correct
		for i, v := range data { //ranging data with already exesting pairs
			if v.Key == key { //if match is present
				data[i].Value = reValue.FindString(r.URL.Path)           //rewriting value to existing key
				fmt.Fprintf(w, "The value for key %s was rewrited", key) //text after rewriting
			}
		}
		data = append(data, KV{key, reValue.FindString(r.URL.Path)}) //adding pair to data
		fmt.Fprintf(w, "The pair was added")                         //text after addind

	} else if reKey.FindString(r.URL.Path) != "" { //if path include only key like /ints:
		for _, v := range data { //ranging data with already exesting pairs
			if v.Key == key { //if match is present
				fmt.Fprintf(w, "The value for key %s is %s", key, v.Value) //text with value
			}
		}
		fmt.Fprintf(w, "Theres no value for this key") //text if value isnt present

	} else {
		fmt.Fprintf(w, "The request is incorrect") //text if request is incorrect
	}
}

func main() {
	data = append(data, KV{"", ""})

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
