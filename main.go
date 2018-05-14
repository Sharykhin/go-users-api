package main

import (
	"log"
	"github.com/Sharykhin/go-users-api/http"
)

func main() {
	log.Fatal(http.ListenAnsServe())
}
