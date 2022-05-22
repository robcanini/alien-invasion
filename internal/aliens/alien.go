package aliens

import (
	petname "github.com/dustinkirkland/golang-petname"
	"github.com/robcanini/alien-invasion/internal/grid"
	"math/rand"
	"time"
)

type Alien struct {
	Name   string
	Target *grid.City
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func generateName() string {
	rand.Seed(time.Now().UnixNano())
	return petname.Generate(2, " ")
}
