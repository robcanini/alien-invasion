package aliens

import (
	"errors"
	"fmt"
	"github.com/robcanini/alien-invasion/internal/grid"
	"math/rand"
)

var aliens []*Alien

func SpreadOn(grid []*grid.City, number int) (error, []*Alien) {
	if number > len(grid) {
		return errors.New(
			fmt.Sprintf("aliens number (%d) must be lower or equal than the cities number (%d) in the map",
				number, len(grid))), nil
	}
	shuffleSlice(&grid)
	aliensSlice := make([]*Alien, number)
	for index := 0; index < number; index++ {
		aliensSlice[index] = createAlien(grid[index])
	}
	aliens = aliensSlice
	return nil, aliensSlice
}

func shuffleSlice(slicePtr *[]*grid.City) {
	slice := *slicePtr
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func FindInvaderOf(city *grid.City) *Alien {
	for _, alien := range aliens {
		if alien.TargetCity == city {
			return alien
		}
	}
	return nil
}
