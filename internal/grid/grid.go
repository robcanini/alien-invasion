package grid

type City struct {
	name  string
	roads []*Road
}

type Direction string

const (
	North Direction = "north"
	South           = "south"
	East            = "east"
	West            = "west"
)

type Road struct {
	direction   Direction
	destination *City
}
