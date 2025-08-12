package main

import (
	"flag"
	"fmt"
	"os"

	"pearshakes/internal"
)

func main() {
	dir := flag.String("dir", "./text", "directory of .txt files (recursively)")
	flag.Parse()

	app := internal.NewApp(*dir)
	if err := app.BuildIndex(); err != nil {
		fmt.Fprintf(os.Stderr, "build index: %v\n", err)
		os.Exit(1)
	}
	app.RunShell()
}
