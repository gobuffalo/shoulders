package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gobuffalo/shoulders/shoulders"
)

var flags = struct {
	Write bool
	JSON  bool
	Name  string
}{}

func init() {
	flag.StringVar(&flags.Name, "n", "", "name of the project")
	flag.BoolVar(&flags.Write, "w", false, "write SHOULDERS.md to disk")
	flag.BoolVar(&flags.JSON, "j", false, "print JSON of format of the dep list")
	flag.Parse()
}

func main() {
	flag.Parse()

	view, err := shoulders.New()
	if err != nil {
		log.Fatal(err)
	}

	flags.Name = strings.TrimSpace(flags.Name)
	if len(flags.Name) > 0 {
		view.Name = flags.Name
	}

	if flags.JSON {
		deps, err := view.DepList()
		if err != nil {
			log.Fatal(err)
		}
		if err := json.NewEncoder(os.Stdout).Encode(deps); err != nil {
			log.Fatal(err)
		}
		return
	}
	var w io.Writer = os.Stdout
	if flags.Write {
		w, err = os.Create("SHOULDERS.md")
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := view.Write(w); err != nil {
		log.Fatal(err)
	}
}
