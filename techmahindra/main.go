/*
I want to build a server to create and read information for products: A product have these properties

* Name
* Description
* Price
* Tags

- Complete the implementation to save and retrieve products.
- Make sure there is a store layer implemented it can be file or in-memory.
- Use the provided source code to get it started.

*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Tags        []string
}

var database []Product = []Product{
	{
		Name:        "Botella",
		Description: "Almacenar agua",
		Price:       10,
		Tags:        []string{"1 lt", "transparente"},
	},
	{
		Name:        "contenedor",
		Description: "Almacenar agua",
		Price:       15,
		Tags:        []string{"oscirp"},
	},
}

func main() {
	s := server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		switch r.Method {
		case http.MethodGet:
			s.getProducts(w, r)
			return
		case http.MethodPost:
			s.createProduct(w, r)
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}

	})
	log.Print("starting server")
	log.Fatal(http.ListenAndServe(":8081", nil))

}

type server struct{}

func (s *server) getProducts(w http.ResponseWriter, r *http.Request) {
	// implement code to get products

	payloads, err := json.Marshal(database)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(payloads)
}

func (s *server) createProduct(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var p Product
	err = json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println(err)
	}

	database = append(database, p)
}
