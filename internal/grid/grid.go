package grid

import (
	"github.com/robcanini/alien-invasion/internal/utils"
	"math/rand"
	"sync"
)

type City struct {
	Name        string
	Roads       []*Road
	invaderFlag *sync.Mutex
	invaderName string
	Destroyed   bool
}

func CreateCity(name string) *City {
	return &City{
		Name:        name,
		invaderFlag: &sync.Mutex{},
		Destroyed:   false,
	}
}

func (city *City) IsInvaded() bool {
	return utils.IsMutexLocked(city.invaderFlag)
}

func (city *City) Free() {
	if utils.IsMutexLocked(city.invaderFlag) {
		city.invaderFlag.Unlock()
	}
	city.invaderName = ""
}

func (city *City) Invade(invaderName string) {
	city.invaderFlag.Lock()
	city.invaderName = invaderName
}

func (city *City) RandomDirection() *City {
	roads := city.getCrossableRoads()
	if len(roads) == 0 {
		return nil
	}
	randomIndex := rand.Intn(len(roads))
	return roads[randomIndex].Destination
}

func (city *City) getCrossableRoads() []*Road {
	roads := make([]*Road, 0)
	for _, road := range city.Roads {
		if road.crossable {
			roads = append(roads, road)
		}
	}
	return roads
}

type Direction string

const (
	North Direction = "north"
	South           = "south"
	East            = "east"
	West            = "west"
)

type Road struct {
	Direction   Direction
	Destination *City
	crossable   bool
}

func CreateRoad(direction Direction, destination *City) *Road {
	return &Road{
		Direction:   direction,
		Destination: destination,
		crossable:   true,
	}
}
