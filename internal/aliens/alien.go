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
	dead       bool
	idle       bool
}

const MaxIterations = 1000

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func createAlien(targetCity *grid.City) *Alien {
	alienName := generateName()
	return &Alien{
		Name:       alienName,
		Steps:      0,
		TargetCity: targetCity,
		dead:       false,
		idle:       false,
	}
}

func generateName() string {
	rand.Seed(time.Now().UnixNano())
	return petname.Generate(2, "-")
}

func (alien *Alien) increaseStepsCounter() {
	alien.Steps++
}

func (alien *Alien) Startup(wg *sync.WaitGroup) {
	alien.sync = wg
	for alien.Steps < MaxIterations && !alien.dead {
		city := alien.aimNextCity()
		// no directions available, alien trapped
		if city == nil {
			alien.trapped()
			break
		}
		alien.invade(city)
	}
}

func (alien *Alien) invade(targetCity *grid.City) {
	alien.increaseStepsCounter()
	// free previous occupied city
	if alien.TargetCity != targetCity {
		alien.TargetCity.Free()
	}
	if targetCity.IsInvaded() {
		multiplexedFight(targetCity, func() { fightOccupant(alien, targetCity) })
	} else {
		conquerCity(alien, targetCity)
	}
}

func multiplexedFight(city *grid.City, fightFn func()) {
	city.FightLock.Lock()
	defer city.FightLock.Unlock()
	if !city.Destroyed {
		fightFn()
	}
}

func (alien *Alien) aimNextCity() *grid.City {
	if alien.Steps == 0 {
		return alien.TargetCity
	}
	return alien.TargetCity.RandomDirection()
}

func (alien *Alien) die() {
	alien.TargetCity.Free()
	alien.dead = true
	if !alien.idle {
		alien.sync.Done()
	}
}

func (alien *Alien) trapped() {
	alien.idle = true
	alien.sync.Done()
}

func conquerCity(attacker *Alien, targetCity *grid.City) {
	targetCity.Invade(attacker.Name)
	attacker.TargetCity = targetCity
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
	finalClash(attacker, occupant, targetCity)
}

func finalClash(attacker *Alien, defensor *Alien, city *grid.City) {
	fmt.Printf("%s has been destroyed by alien %s and alien %s!\n", city.Name, attacker.Name, defensor.Name)
	grid.DestroyCity(city)
	attacker.die()
	defensor.die()
}
