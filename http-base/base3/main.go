package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("into index")
		fmt.Printf("URL.Path = %v\n", r.URL.Path)
	})
	engine.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("into hello")
		for k, v := range r.Header {
			fmt.Printf("Header[%v] = %v \n", k, v)
			fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
		}
	})

	engine.Run("9999")
}
