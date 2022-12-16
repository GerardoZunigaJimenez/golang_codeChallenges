package endpoint

import (
	"fmt"
	"net/http"
)

func StarshipsAvailable(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hello")
}
