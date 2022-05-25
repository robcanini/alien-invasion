package main

import (
	"flag"
	"fmt"
	"github.com/robcanini/alien-invasion/internal/invasion"
	"os"
)

var (
	filePath     = flag.String("filePath", "not_specified", "The grid data file")
	aliensNumber = flag.Int("aliens", -1, "The number of aliens invading the planet")
)

func main() {
	flag.Parse()
	enforceCliArgs()

	err := invasion.Run(invasion.Spec{
		PlanetName:           "X",
		PlanetGridSourceType: invasion.FileSource,
		PlanetGridSourceUri:  *filePath,
		AliensNumber:         *aliensNumber,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func enforceCliArgs() {
	if *filePath == "not_specified" {
		fmt.Println("You must specify --filePath argument. Check examples/data.db for file sample")
		os.Exit(1)
	}

	if *aliensNumber <= 0 {
		fmt.Println("You must set --aliens argument with a valid number (from 1 to number of cities in grid)")
		os.Exit(1)
	}
}
