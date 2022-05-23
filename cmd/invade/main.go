package main

import (
	"flag"
	"fmt"
	"github.com/robcanini/alien-invasion/internal/invasion"
	"os"
)

var (
	filePath     = flag.String("filePath", "grid_data.db", "The grid data file")
	aliensNumber = flag.Int("aliens", 5, "The number of aliens invading the planet")
)

func main() {
	flag.Parse()

	err, _ := invasion.Run(invasion.Spec{
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
