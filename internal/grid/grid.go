package grid

import "github.com/robcanini/alien-invasion/internal/aliens"

type City struct {
	Name    string
	Roads   []*Road
	Invader *aliens.Alien
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
