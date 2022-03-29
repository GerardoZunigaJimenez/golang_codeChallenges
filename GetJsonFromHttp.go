package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type people struct {
	Number int `json:"number"`
}

func main() {

	url := "https://storage.googleapis.com/product-exports/product_data.json"

	spaceClient := http.Client{
		Timeout: time.Minute * 3, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	err = ioutil.WriteFile("big_marhsall.json", body, os.ModePerm)
	if err != nil{
		log.Fatal(err)
	}

}