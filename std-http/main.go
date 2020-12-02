package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := `{"healthy": true}`+"\n"
		_, err := w.Write([]byte(res))
		if err != nil {
			log.Fatal("could not write the response", err)
		}

		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("could not read response body", err)
		}
		fmt.Println("response body:", string(bs))
	})
	fmt.Println("listening for clients")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("server error", err)
	}
}
