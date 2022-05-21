package grid

import (
	"fmt"
)

// first file entry must be city name
// next entries must be in the format x=y
// there should be at least one direction otherwise city cannot be reached or become instantly an alien trap
// city directions cannot contain a self city reference
// city name cannot contains '='

var state State

type State struct {
	data []*City
}

func InitState(fetcher *Fetcher, evChan chan *City) error {
	err, cities := (*fetcher).fetchGrid()
	if err != nil {
		return err
	}
	state = State{data: cities}
	listenToUpdateEvents(evChan)
	return nil
}

func listenToUpdateEvents(evChan chan *City) {
	for city := range evChan {
		fmt.Println(city)
	}
}
