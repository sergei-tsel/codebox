package main

import (
	"codebox/db"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	godotenv.Load(".env")

	db.Init()

	r := NewRouter()

	http.ListenAndServe(":8080", r)
}
