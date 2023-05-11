package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/meisterunused/badgeify"
)

func printUsage() {
	usage := `Usage:
  $ badgeify # lookup coverage automatically
  $ badgeify --label "coverage" --value "80%" > coverage-master.svg
  $ badgeify --color cdcdcd --label "works" --value "ok" > coverage-master.svg`
	fmt.Println(usage)
}

func main() {
	help := flag.Bool("help", false, "show help")
	version := flag.Bool("version", false, "show version")
	label := flag.String("label", "coverage", "label (default: coverage)")
	value := flag.String("value", "", "value")
	color := flag.String("color", "", "a custom color in hex format w/o leading #")

	flag.Parse()

	if *version {
		fmt.Printf("badgeify v%s\n", badgeify.Version)
		os.Exit(0)
	}
	if *help {
		printUsage()
		os.Exit(0)
	}

	if *value == "" {
		cov, _ := badgeify.LookupCoverage()
		autoValue := fmt.Sprintf("%d%%", int(cov))
		value = &autoValue
	}

	badge := badgeify.New(*label, *value)
	if *color != "" {
		badge.CustomColor = *color
	}
	if err := badge.Print(os.Stdout); err != nil {
		panic(err)
	}
}
