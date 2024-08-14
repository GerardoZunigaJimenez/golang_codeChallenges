package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	handlerRoot := http.HandlerFunc(getRoot)
	handlerGetHello := http.HandlerFunc(getHello)

	http.Handle("/", middleware(handlerRoot))
	http.Handle("/hello", middleware(handlerGetHello))

	for {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, HTTP!\n")
}

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("url %s, and took time: %v", r.URL.Path, time.Since(now))
	})
}
