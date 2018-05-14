package main

import (
	"fmt"
	"log"
	"github.com/Sharykhin/go-users-api/http"
)

func main() {

	fmt.Println("Hello World %d")
	log.Fatal(http.ListenAnsServe())
}
