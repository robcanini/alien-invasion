package grid

import (
	"fmt"
)

var state State

type State struct {
	data []*City
}

func InitState(fetcher Fetcher, evChan chan *City) error {
	err, cities := fetcher.FetchGrid()
	if err != nil {
		return err
	}
	fmt.Println(cities)
	for _, city := range cities {
		fmt.Println(*city)
		for _, road := range city.Roads {
			fmt.Println(*road)
		}
		fmt.Println()
	}
	state = State{data: cities}
	// listenToUpdateEvents(evChan)
	return nil
}

func listenToUpdateEvents(evChan chan *City) {
	for city := range evChan {
		fmt.Println(city)
	}
}
