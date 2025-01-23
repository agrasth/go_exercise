package main

import (
    "log"
    "net/http"
    "item-manager/internal/api"
    "item-manager/internal/data"
)

func main() {
    storage := data.NewStorage("items.json")
    server := api.NewServer(storage)
    log.Fatal(http.ListenAndServe(":8080", server))
}