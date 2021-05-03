package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", greet)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func isAuthorized(r *http.Request) bool {
	raw := r.Header.Get("Authorization")

	var token string

	// Authorization: Bearer opuy028qfaujawowhgawwtadaewayawarat
	_, err := fmt.Sscanf(raw, "SecretKey %s", &token)
	if err != nil {
		return false
	}

	return token == "secret"
}

func greet(w http.ResponseWriter, r *http.Request) {
	if !isAuthorized(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	_, err := w.Write([]byte("Hola Marcos"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
