package server

import "net/http"

func StartServer() {
	http.HandleFunc("/randomPass", generateRandomPasswordHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
