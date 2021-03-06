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

func DestroyCity(city *City) {
	city.Free()
	city.Destroyed = true
	removeCityRefs(city)
}

func removeCityRefs(city *City) {
	for _, it := range grid {
		for _, road := range it.Roads {
			if road.Destination == city {
				road.crossable = false
			}
		}
	}
}

func GetGrid() []*City {
	return grid
}

func PrintGrid() {
	fmt.Println()
	for _, city := range grid {
		if city.Destroyed {
			continue
		}
		fmt.Printf("%s", city.Name)
		for _, road := range city.Roads {
			if road.crossable {
				fmt.Printf(" %s=%s", road.Direction, road.Destination.Name)
			}
		}
		fmt.Println()
	}
}
