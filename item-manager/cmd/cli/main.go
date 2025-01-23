package main

import (
    "log"
    "os"
    
    cli_app "item-manager/internal/cli"
    "item-manager/internal/data"
    "github.com/urfave/cli/v2"
)

func main() {
    storage := data.NewStorage("items.json")
    cli_app.Storage = storage
    
    app := &cli.App{
        Name:    "Shopping CLI",
        Usage:   "Manage shopping items",
        Version: "1.0.0",
        Authors: []*cli.Author{
            {
                Name:  "Agrasth Naman",
                Email: "agrasthn@jfrog.com",
            },
        },
        Commands: cli_app.Commands,
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}