package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var appHash = "Application " + generateRandomHash(6)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%s. Request %s.", appHash, generateRandomHash(6))))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func generateRandomHash(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"        // Characters to use for the hash
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))] // Pick a random character from the charset
	}
	return string(b)
}
