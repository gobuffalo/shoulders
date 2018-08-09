package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gobuffalo/shoulders/shoulders"
)

var flags = struct {
	Write bool
	JSON  bool
	Name  string
}{}

func init() {
	pkg, err := shoulders.CurrentPkg()
	if err != nil {
		log.Fatal(err)
	}
	flag.StringVar(&flags.Name, "n", fmt.Sprintf("`%s`", pkg), "name of the project")
	flag.BoolVar(&flags.Write, "w", false, "write SHOULDERS.md to disk")
	flag.BoolVar(&flags.JSON, "j", false, "print JSON of format of the dep list")
	flag.Parse()
}

func main() {
	view, err := shoulders.New()
	if err != nil {
		log.Fatal(err)
	}
	view.Name = flags.Name

	if flags.JSON {
		if err := json.NewEncoder(os.Stdout).Encode(view.Deps); err != nil {
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
