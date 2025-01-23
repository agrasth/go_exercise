package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)


func main() {
	app := cli.NewApp()
	app.Name = "Calculator"
	app.Usage = "A simple calculator"
	app.Version = "1.0.0"
	app.Description = "This is a simple calculator"
	app.Authors = []*cli.Author{
		{Name: "Agrasth naman", Email: "agrasthn@jfrog.com"},
	}

	// app.Flags = []cli.Flag{
	// 	&cli.StringFlag{
	// 		Destination: &operation,
	// 		Name: "operation",
	// 		Value: "add",
	// 		Usage: "operation to perform add/sub/mul/div",
	// 	},
	// 	&cli.IntFlag{
	// 		Destination: &a,
	// 		Name: "a",
	// 		Value: 0,
	// 		Usage: "first operand",
	// 	},
	// 	&cli.IntFlag{
	// 		Destination: &b,
	// 		Name: "b",
	// 		Value: 0,
	// 		Usage: "second operand",
	// 	},
	// }

	app.Action = calculator

	app.Commands = []*cli.Command{
		addCommand(),
		subCommand(),
		mulCommand(),
		divCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func divCommand() *cli.Command {
	return &cli.Command{
		Name: "div",
		Usage: "div two numbers",
		Aliases: []string{"d"},
		Action: func(ctx *cli.Context) error {
			n := ctx.NArg()
			if n == 0 {
				return errors.New("No arguments provided")
			}
			fmt.Printf("Number of arguments: %d\n", n) 
			val := ctx.Args().Get(0)
			res, _ := strconv.Atoi(val)
			fmt.Printf("%v", res)
			for i := 1; i < n; i++ {
				val = ctx.Args().Get(i)
				op, _ := strconv.Atoi(val)
				res /= op
				fmt.Printf(" / %v", op)
			}
			fmt.Printf(" = %v\n", res)
			return nil
		},
	}
}

func addCommand() *cli.Command {
	return &cli.Command{
		Name: "add",
		Usage: "add two numbers",
		Aliases: []string{"a"},
		Action: func(ctx *cli.Context) error {
			n := ctx.NArg()
			if n == 0 {
				return errors.New("No arguments provided")
			}
			fmt.Printf("Number of arguments: %d\n", n) 
			val := ctx.Args().Get(0)
			res, _ := strconv.Atoi(val)
			fmt.Printf("%v", res)
			for i := 1; i < n; i++ {
				val = ctx.Args().Get(i)
				op, _ := strconv.Atoi(val)
				res += op
				fmt.Printf(" + %v", op)
			}
			fmt.Printf(" = %v\n", res)
			return nil
		},
	}
}

func mulCommand() *cli.Command {
	return &cli.Command{
		Name: "mul",
		Usage: "mul two numbers",
		Aliases: []string{"m"},
		Action: func(ctx *cli.Context) error {
			n := ctx.NArg()
			if n == 0 {
				return errors.New("No arguments provided")
			}
			fmt.Printf("Number of arguments: %d\n", n) 
			val := ctx.Args().Get(0)
			res, _ := strconv.Atoi(val)
			fmt.Printf("%v", res)
			for i := 1; i < n; i++ {
				val = ctx.Args().Get(i)
				op, _ := strconv.Atoi(val)
				res *= op
				fmt.Printf(" * %v", op)
			}
			fmt.Printf(" = %v\n", res)
			return nil
		},
	}
}

func subCommand() *cli.Command {
	return &cli.Command{
		Name: "sub",
		Usage: "sub two numbers",
		Aliases: []string{"s"},
		Action: func(ctx *cli.Context) error {
			n := ctx.NArg()
			if n == 0 {
				return errors.New("No arguments provided")
			}
			fmt.Printf("Number of arguments: %d\n", n) 
			val := ctx.Args().Get(0)
			res, _ := strconv.Atoi(val)
			fmt.Printf("%v", res)
			for i := 1; i < n; i++ {
				val = ctx.Args().Get(i)
				op, _ := strconv.Atoi(val)
				res -= op
				fmt.Printf(" - %v", op)
			}
			fmt.Printf(" = %v\n", res)
			return nil
		},
	}
}

func calculator(ctx *cli.Context) error {

	// switch operation {
	// case "add":
	// 	res := a + b
	// 	fmt.Printf("Result: %d\n", res)
	// case "sub":
	// 	res := a - b
	// 	fmt.Printf("Result: %d\n", res)
	// case "mul":
	// 	res := a * b
	// 	fmt.Printf("Result: %d\n", res)
	// case "div":
	// 	if b == 0 {
	// 	return errors.New("Division by zero")
	// 	}
	// 	res := a / b
	// 	fmt.Printf("Result: %d\n", res)

	// }
	ctx.App.Command("help").Run(ctx)
	return nil
}
