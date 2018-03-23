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

var (
	data []KV

	validKV = regexp.MustCompile("/[0-9]{1,10}:[a-zA-Z]{1,15}/")
	reKey = regexp.MustCompile("/[0-9]{1,10}:")
	reValue = regexp.MustCompile("[a-zA-Z]{1,15}")
	reKeyNums = regexp.MustCompile("[0-9]{1,10}")
)

func handler(w http.ResponseWriter, r *http.Request) {

	key := reKeyNums.FindString(r.URL.Path)

	if r.URL.Path == "/" {
		//sending text when url in address bar is like site.com/
		w.Write([]byte("Enter key and value in form: sitename.com/key:value/. \nA key must contain 1-10 integers, a value must contain 1-15 strings. \nIf a pair key:value is already exis, you can get the value by key, enter sitename.com/key:/"))

	} else if validKV.FindString(r.URL.Path) != "" { //if key and value are both correct
		for i := range data { //ranging data with already exesting pairs
			if data[i].Key == key { //if match is present
				data[i].Value = reValue.FindString(r.URL.Path)           //rewriting value to existing key
				fmt.Fprintf(w, "The value for key %s was rewrited", key)
				return
			}
		}
		data = append(data, KV{key, reValue.FindString(r.URL.Path)}) //adding pair to data
		w.Write([]byte("The pair was added"))
                     
	} else if reKey.FindString(r.URL.Path) != "" { //if path include only key like /ints:
		for i := range data { //ranging data with already exesting pairs
			if data[i].Key == key { //if match is present
				fmt.Fprintf(w, "The value for key %s is %s", key, data[i].Value)
				return
			}
		}
		w.Write([]byte("There's no value for this key")) //text if value isnt present

	} else {
		w.Write([]byte("The request is incorrect")) //text if request is incorrect
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
