package main

import (
	"api/internal/config"
	"api/internal/router"
	"fmt"
	"log"
	"net/http"
)

/*
	func init() {
		chave := make([]byte, 64)

		if _, err := rand.Read(chave); err != nil {
			log.Fatal(err)
		}

		strinBase64 := base64.StdEncoding.EncodeToString(chave)

		fmt.Println(strinBase64)
	}
*/
func main() {
	config.Carregar()

	fmt.Println("Start run application port :5000")
	r := router.Router()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
