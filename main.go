package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PauloHPMKT/cronometro/router"
)

type Timer interface {
	Start()
}

func main() {
	router.Router()

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
