package aliens

import (
	"errors"
	"fmt"
	"github.com/robcanini/alien-invasion/internal/grid"
	"math/rand"
)

func SpreadOn(grid []*grid.City, number int) (error, []*Alien) {
	if number > len(grid) {
		return errors.New(fmt.Sprintf("aliens number (%d) must be lower or equal than the cities number (%d) in the map", number, len(grid))), nil
	}
	shuffleGridSlice(&grid)
	aliens := make([]*Alien, number)
	for index := 0; index < number; index++ {
		aliens[index] = createAlien(grid[index])
	}
	return nil, aliens
}

func shuffleGridSlice(slicePtr *[]*grid.City) {
	slice := *slicePtr
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
