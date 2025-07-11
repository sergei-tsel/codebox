package main

import "net/http"

func Pong(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("pong"))
}
