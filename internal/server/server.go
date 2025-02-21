package server

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/randomPass", generateRandomPasswordHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
