package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func generatePassword(length int) string {
	rand.Seed(time.Now().UnixNano())

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}|:<>?-=[];,."
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}

func passwordGeneratorHandler(w http.ResponseWriter, r *http.Request) {
	length := 12 // Длина пароля по умолчанию
	if len(r.URL.Query()["length"]) > 0 {
		length = 12 // Можно задать длину пароля через параметр запроса в URL
	}

	password := generatePassword(length)
	fmt.Fprintf(w, "Сгенерированный пароль: %s", password)
}

func main() {
	http.HandleFunc("/generate-password", passwordGeneratorHandler)
	fmt.Println("Сервер запущен на http://localhost:8080/generate-password")
	http.ListenAndServe(":8080", nil)
}
