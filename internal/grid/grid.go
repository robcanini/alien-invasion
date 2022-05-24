package grid

import (
	"github.com/robcanini/alien-invasion/internal/utils"
	"sync"
)

type City struct {
	Name        string
	Roads       []*Road
	InvaderFlag *sync.Mutex
}

func (city *City) IsInvaded() bool {
	return city.InvaderFlag != nil && utils.IsMutexLocked(city.InvaderFlag)
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
