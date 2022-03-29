package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	pageSize := 200
	for i:=0; i <= len(ids); i = i+pageSize {

		var idsString = ""
		for j:=i; j<i+pageSize; j++{
			if j>=len(ids){
				break
			}
			idsString = idsString + ids[j] + ","
		}
		idsString = idsString[0:(len(idsString)-1)]


		url := "https://api.honey-badgers.omg.pub/product/products/" + idsString
		// Create a new request using http
		req, err := http.NewRequest("DELETE", url, nil)
		if req == nil || req.Header == nil {
			return
		}
		if err != nil {
			fmt.Println("This is the list that we will delete:\n", idsString)
		}

		// Create a Bearer string by appending string access token
		bearer := `Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjdGl2ZV9ncm91cCI6eyJpZCI6IjEiLCJuYW1lIjoicm9vdCJ9LCJhc3N1bWVkX3VzZXIiOnsiZW1haWwiOlt7ImFkZHJlc3MiOiJyb290QGhvbmV5LWJhZGdlcnMub21nLnB1YiJ9XSwiZmlyc3RfbmFtZSI6IlJvb3QiLCJoYXNNdWx0aXBsZUdyb3VwcyI6dHJ1ZSwiaWQiOiIxIiwiaXNDb250cmFjdG9yIjpmYWxzZSwiaXNPbWciOnRydWUsImxhc3RfbmFtZSI6IlJvb3QifSwiYXV0aGVudGljYXRlZF91c2VyIjp7ImVtYWlsIjpbeyJhZGRyZXNzIjoicm9vdEBob25leS1iYWRnZXJzLm9tZy5wdWIifV0sImZpcnN0X25hbWUiOiJSb290IiwiaGFzTXVsdGlwbGVHcm91cHMiOnRydWUsImlkIjoiMSIsImlzQ29udHJhY3RvciI6ZmFsc2UsImlzT21nIjp0cnVlLCJsYXN0X25hbWUiOiJSb290In0sInRva2VuIjoiU0c2SGRpUHNHRkJWM2ExYUdoR25VQUktbWFXU2FBbktjUzB5R1lSOSJ9LCJleHAiOjE1OTQxNzIxNDZ9.g2KW5tiFNb3MiaYHsblFh3iGyHglbfHc3C7-4V5XA14K4roMJdxNTublle9pZbCmsbbIKRlIm749h1kXwMXF6nzjMTfVWO__pE5zdNFAXuARsq_mPxZtcN0wFFHWo5e57iE7Bb-TNh7ObePkAiGXfFgR-tcrO_Cx01PxA5DcBYom9NQZGAlm5FnLprtefs8iTBxcJrY0rQDtpi6HjIVVTaTzTqTitKo8yG7gpsUEJLUtSUPMZ0hzrdwL7Z34wgTu0MQr7klWwBAAUK9E35u1sJh9AnFGK77t5gzCLJLm1AXAWaqztIJHaz_FiS6d2jBytKto_WeDSks_y_bm4gPRag`
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error on response.\n[ERRO] -", err)
		}
		if resp != nil && resp.Body != nil {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("🌐", resp.Status, string([]byte(body)))
		}

		r := rand.Intn(30)
		time.Sleep(time.Duration(r) * time.Second)
		fmt.Println( "continue after a small nap", r)
	}
}

var ids = []string{`03981890-634e-4628-7191-776dd1003366`,
	`0ad5eedb-cbe5-49c2-7ea7-52a631d1a87e`,
	`0ecfcaa2-f6fb-4db9-4988-dcb6526544ee`,
	`2e4b4ef0-54cb-4e6f-7caf-59dc910966df`,
	`37373f23-26d9-4788-5f2a-c1576e50b01b`,
	`82eeb832-45a3-4245-5de0-90eead67f732`,
	`85feed47-7a57-4497-530d-775fb870bd5f`,
	`8a9467c5-3abe-47ad-74ba-a69efa8bc2c8`,
	`8bcdb524-e3fc-4544-6780-24ec14c09cc2`,
	`90ad392c-25f1-48cf-5c05-c4596e321c23`,
	`9ab620ed-cdd4-4ef3-5155-d6e7f9a3e571`,
	`bc75a69b-68f2-483d-5b31-144273cc4b07`,
	`cc676081-3256-4a8c-5ec8-09ffd0c11551`,
	`ce671b5b-dcce-432f-57e7-b5253a4e7073`,
	`df582d84-c610-4e7c-46a4-80c71aecb89c`,
	`e596a4ad-2543-4769-79a3-ffabda87106a`,
	`f65c12a9-e5fe-4223-7855-4d9879a00093`,
}