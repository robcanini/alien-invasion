package grid

import (
	"fmt"
)

var state State

type State struct {
	data []*City
}

func Init(fetcher Fetcher) (error, []*City) {
	err, cities := fetcher.FetchGrid()
	if err != nil {
		return err, nil
	}
	state = State{data: cities}
	PrintState()
	return nil, cities
}

func PrintState() {
	fmt.Println(state.data)
	for _, city := range state.data {
		fmt.Println(*city)
		for _, road := range city.Roads {
			fmt.Println(*road)
		}
		fmt.Println()
	}
}
