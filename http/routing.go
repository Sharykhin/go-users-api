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

func ListenAnsServe() error {
	address := os.Getenv("HTTP_ADDRESS")
	fmt.Printf("Go Users Api Service starts listening on %s", address)
	return http.ListenAndServe(address, router())
}
