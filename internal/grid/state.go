package grid

import (
	"fmt"
)

var grid []*City

func Load(fetcher Fetcher) error {
	err, cities := fetcher.FetchGrid()
	if err != nil {
		return err
	}
	grid = cities
	return nil
}

// GetGrid todo: encapsulation required (?)
func GetGrid() []*City {
	return grid
}

func PrintGrid() {
	fmt.Println(grid)
	for _, city := range grid {
		fmt.Println(*city)
		for _, road := range city.Roads {
			fmt.Println(*road)
		}
	}
}
