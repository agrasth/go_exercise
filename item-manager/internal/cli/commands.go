package cli

import (
    "errors"
    "fmt"

    "github.com/urfave/cli/v2"
    "item-manager/internal/data"
)

var Storage *data.Storage

var Commands = []*cli.Command{
    {
        Name:    "add",
        Usage:   "Add a new shopping item",
        Aliases: []string{"a"},
        Action: func(c *cli.Context) error {
            name := c.Args().First()
            if name == "" {
                return errors.New("no item name provided")
            }
            item := data.Item{ID: data.NewUUID(), Name: name}
            Storage.AddItem(item)
            fmt.Printf("Added item: %v\n", item)
            return nil
        },
    },
    {
        Name:    "remove",
        Usage:   "Remove a shopping item by id",
        Aliases: []string{"r"},
        Action: func(c *cli.Context) error {
            id := c.Args().First()
            if id == "" {
                return errors.New("no id provided")
            }
            Storage.RemoveItem(id)
            fmt.Printf("Removed item with ID: %v\n", id)
            return nil
        },
    },
    {
        Name:    "list",
        Usage:   "List all shopping items",
        Aliases: []string{"l"},
        Action: func(c *cli.Context) error {
            items := Storage.GetItems()
            for i, item := range items {
                fmt.Printf("%d: %v\n", i+1, item)
            }
            return nil
        },
    },
}