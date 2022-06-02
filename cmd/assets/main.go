package assets

import (
	"fmt"
	"github.com/PaluMacil/wasm-component/web"
	"io/fs"
	"log"
	"os"
)

func list() {
	fs.WalkDir(web.Templates, "/", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
}

func main() {
	args := os.Args
	if len(args) > 2 {
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
