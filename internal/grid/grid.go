package grid

import (
	"fmt"
	"github.com/robcanini/alien-invasion/internal/utils"
	"sync"
)

type City struct {
	Name        string
	Roads       []*Road
	invaderFlag *sync.Mutex
}

func CreateCity(name string) *City {
	return &City{Name: name, invaderFlag: &sync.Mutex{}}
}

func (city *City) IsInvaded() bool {
	return utils.IsMutexLocked(city.invaderFlag)
}

func (city *City) Free() {
	fmt.Printf("City of %s has been released\n", city.Name)
	city.invaderFlag.Unlock()
}

func (city *City) Invade() {
	fmt.Printf("City of %s has been invaded\n", city.Name)
	city.invaderFlag.Lock()
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
}
