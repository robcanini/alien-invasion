package grid

import (
	"fmt"
)

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
