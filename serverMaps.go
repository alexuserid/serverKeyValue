package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var (
	data    = make(map[string]string)
	validKV = regexp.MustCompile("/([0-9]{1,10}):([a-zA-Z]{1,15})*/")
)

func handler(w http.ResponseWriter, r *http.Request) {

	KVarr := validKV.FindStringSubmatch(r.URL.Path)

	if r.URL.Path == "/" {
		w.Write([]byte("Enter key and value in form: sitename.com/key:value/. \nA key must contain 1-10 integers, a value must contain 1-15 strings. \nIf a pair key:value is already exis, you can get the value by key, enter sitename.com/key:/"))

	} else if KVarr != nil {
		if KVarr[2] != "" {
			nums, ok := data[KVarr[1]] //maps return two values. If key exist second value will be true
			data[KVarr[1]] = KVarr[2]  //in any case new value will be written with this key
			if ok {
				fmt.Fprintf(w, "The value for key %s was rewrited", nums) //if key already exist
			} else {
				w.Write([]byte("The pair was added"))
			}

		} else if KVarr[2] == "" {
			_, ok := data[KVarr[1]]
			if ok {
				fmt.Fprintf(w, "The value for key %s is %s", KVarr[1], data[KVarr[1]])
			} else {
				w.Write([]byte("There's no value for this key")) //text if value isnt present
			}
		}

	} else {
		w.Write([]byte("The request is incorrect")) //text if request is incorrect
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
