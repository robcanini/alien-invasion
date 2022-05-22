package main

import (
	"flag"
	"fmt"
	"github.com/robcanini/alien-invasion/internal/aliens"
	"github.com/robcanini/alien-invasion/internal/grid"
	"github.com/robcanini/alien-invasion/internal/io"
	"os"
)

var (
	filePath     = flag.String("filePath", "grid_data.db", "The grid data file")
	aliensNumber = flag.Int("aliens", 5, "The number of aliens invading the planet")
)

func main() {
	flag.Parse()
	fmt.Println("Planet X invasion started")

	err, cities := grid.Init(&io.FileFetcher{FilePath: *filePath})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = aliens.SpawnInGrid(cities, *aliensNumber)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
