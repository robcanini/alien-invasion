package main

import (
	"fmt"
	"github.com/robcanini/alien-invasion/internal/grid"
	"github.com/robcanini/alien-invasion/internal/io"
)

func main() {
	fmt.Println("Planet X invasion started")

	ch := make(chan *grid.City)

	err := grid.InitState(&io.FileFetcher{FilePath: "/tmp/grid_data.db"}, ch)
	if err != nil {
		fmt.Println(err)
	}
}
