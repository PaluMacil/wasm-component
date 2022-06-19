package main

import (
	"fmt"
	"github.com/PaluMacil/wasm-component/example"
	"io/fs"
	"log"
	"os"
)

func list() {
	fs.WalkDir(example.Templates, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if d.IsDir() {
			return nil
		}
		fmt.Println(path)
		return nil
	})
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("no subcommand given")
	}
	subCommand := args[1]
	switch subCommand {
	case "list":
		list()
	default:
		log.Printf("unknown subcommand: %s", subCommand)
	}
}
