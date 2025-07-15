package main

import (
	"codebox/db"
	"codebox/router"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	godotenv.Load("../.env")

	db.Init()

	r := router.NewRouter()

	http.ListenAndServe(":8080", r)
}
