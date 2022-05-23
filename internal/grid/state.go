package grid

import (
	"errors"
	"fmt"
	"github.com/robcanini/alien-invasion/internal/aliens"
	"math/rand"
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

func SpreadAliens(aliens []*aliens.Alien) error {
	if len(aliens) > len(grid) {
		return errors.New(fmt.Sprintf("aliens number (%d) must be lower or equal than the cities number (%d) in the map", len(aliens), len(grid)))
	}
	shuffleGridSlice(&grid)
	for index := 0; index < len(aliens); index++ {
		invadeCity(grid[index], aliens[index])
	}
	return nil
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

func shuffleGridSlice(slicePtr *[]*City) {
	slice := *slicePtr
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func invadeCity(city *City, alien *aliens.Alien) {
	city.Invader = alien
	fmt.Printf("Alien %s is about to invade %s\n", alien.Name, city.Name)
}
