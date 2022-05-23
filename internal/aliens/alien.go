package aliens

import (
	petname "github.com/dustinkirkland/golang-petname"
	"math/rand"
	"time"
)

type Alien struct {
	Name string
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Create() *Alien {
	alienName := generateName()
	return &Alien{Name: alienName}
}

func generateName() string {
	rand.Seed(time.Now().UnixNano())
	return petname.Generate(2, "-")
}
