package main

import (
	"api/internal/config"
	"api/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	fmt.Println("start run application port :5000")
	r := router.Router()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
