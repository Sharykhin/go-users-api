package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"fmt"
)

func router() http.Handler {
	r := mux.NewRouter()

	return r
}

// ListenAnsServ2 starts listening http server on address that is provided through env variable
func ListenAnsServe() error {
	address := os.Getenv("HTTP_ADDRESS")
	fmt.Printf("go-users-api is listening on %s", address)
	return http.ListenAndServe(address, router())
}
