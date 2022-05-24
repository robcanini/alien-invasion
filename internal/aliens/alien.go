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
	fmt.Printf("Alien %s march against %s\n", alien.Name, targetCity.Name)
	alien.increaseStepsCounter()
	if targetCity.IsInvaded() {
		fmt.Printf(
			"City of %s was already occupied. Alien %s fight against occupant\n", targetCity.Name, alien.Name)
		fightOccupant(alien, targetCity)
	} else {
		conquerCity(alien, targetCity)
	}

	alien.die()
}

func conquerCity(attacker *Alien, targetCity *grid.City) {
	targetCity.Invade()
	attacker.TargetCity = targetCity
	fmt.Printf("City of %s has been conquered by %s\n", targetCity.Name, attacker.Name)
}

func fightOccupant(attacker *Alien, targetCity *grid.City) {
	occupant := FindInvaderOf(targetCity)
	if occupant == nil {
		fmt.Printf("Flag was raised in the unoccupied %s. This appears to be a deadlock\n", targetCity.Name)
		// force flag lowering
		targetCity.Free()
		// it is not necessary to fight, alien can continue its invasion
		return
	}
	// battle begins!
	finalClash(attacker, occupant)
}

func finalClash(attacker *Alien, defensor *Alien) {
	attacker.die()
	defensor.die()
	fmt.Printf("Attacker %s and defensor %s are died in battle\n", attacker.Name, defensor.Name)
	// city is destroyed
}

func (alien *Alien) die() {
	fmt.Printf("Alien %s is dead in %s\n", alien.Name, alien.TargetCity.Name)
	alien.TargetCity.Free()
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
