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
	data    []KV
	validKV = regexp.MustCompile("/([0-9]{1,10}):([a-zA-Z]{1,15})*/")
)

func handler(w http.ResponseWriter, r *http.Request) {

	KVarr := validKV.FindStringSubmatch(r.URL.Path)

	if r.URL.Path == "/" {
		//sending text when url in address bar is like site.com/
		w.Write([]byte("Enter key and value in form: sitename.com/key:value/. \nA key must contain 1-10 integers, a value must contain 1-15 strings. \nIf a pair key:value is already exis, you can get the value by key, enter sitename.com/key:/"))

	} else if KVarr != nil { //if key and value are both correct
		if KVarr[2] != "" {
			for i := range data { //ranging data with already exesting pairs
				if data[i].Key == KVarr[1] { //if match is present
					data[i].Value = KVarr[2] //rewriting value to existing key
					fmt.Fprintf(w, "The value for key %s was rewrited", data[i].Key)
					return
				}
			}
			data = append(data, KV{KVarr[1], KVarr[2]}) //adding pair to data
			w.Write([]byte("The pair was added"))
			return

		} else if KVarr[2] == "" {
			for i := range data { //ranging data with already exesting pairs
				if data[i].Key == KVarr[1] { //if match is present
					fmt.Fprintf(w, "The value for key %s is %s", data[i].Key, data[i].Value)
					return
				}
			}
			w.Write([]byte("There's no value for this key")) //text if value isnt present
			return
		}

	} else {
		w.Write([]byte("The request is incorrect")) //text if request is incorrect
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
