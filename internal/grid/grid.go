package grid

type City struct {
	Name  string
	Roads []*Road
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
