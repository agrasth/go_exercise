package main

import (
	"http-server-project/api"
	"net/http"
)

func main(){
	server := api.NewServer()
	http.ListenAndServe(":8080", server)
}