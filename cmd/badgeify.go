package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/meisterunused/badgeify"
)

func printUsage() {
	usage := `Usage:
  $ badgeify "coverage" "80%" > coverage-master.svg
  $ badgeify --color cdcdcd "works" "ok" > coverage-master.svg`
	fmt.Println(usage)
}

func main() {
	customColor := flag.String("color", "", "a custom color in hex format w/o leading #")
	flag.Parse()

	if len(os.Args) < 3 {
		printUsage()
		os.Exit(-1)
	}

	badge := badgeify.New(os.Args[len(os.Args)-2], os.Args[len(os.Args)-1])
	if *customColor != "" {
		badge.CustomColor = *customColor
	}
	if err := badge.Print(os.Stdout); err != nil {
		panic(err)
	}
}
