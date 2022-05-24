package aliens

import (
	"fmt"
	petname "github.com/dustinkirkland/golang-petname"
	"github.com/robcanini/alien-invasion/internal/grid"
	"math/rand"
	"sync"
	"time"
)

type Alien struct {
	Name       string
	Steps      int
	TargetCity *grid.City
	sync       *sync.WaitGroup
}

func (alien *Alien) increaseStepsCounter() {
	alien.Steps++
}

func (alien *Alien) Startup(wg *sync.WaitGroup) {
	alien.sync = wg
	alien.invade(alien.TargetCity)
}

func (alien *Alien) invade(targetCity *grid.City) {
	fmt.Printf("Alien %s started invading %s\n", alien.Name, targetCity.Name)
	defer alien.increaseStepsCounter()

	// aliens should fight
	if targetCity.IsInvaded() {

	}
}

func (alien *Alien) die() {
	fmt.Printf("Alien %s is dead in %s\n", alien.Name, alien.TargetCity.Name)
	alien.sync.Done()
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func createAlien(targetCity *grid.City) *Alien {
	alienName := generateName()
	return &Alien{Name: alienName, Steps: 0, TargetCity: targetCity}
}

func generateName() string {
	rand.Seed(time.Now().UnixNano())
	return petname.Generate(2, "-")
}
