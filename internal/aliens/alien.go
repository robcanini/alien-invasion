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
	multiplexor := sync.Mutex{}
	for alien.Steps < MaxIterations && !alien.dead && !alien.idle {
		multiplexor.Lock()
		city := alien.aimNextCity()
		alien.leaveCurrentCity()
		// no directions available, alien trapped
		if city == nil {
			alien.trapped()
			multiplexor.Unlock()
			break
		}
		alien.invade(city)
		multiplexor.Unlock()
	}
	alien.sync.Done()
}

func (alien *Alien) leaveCurrentCity() {
	alien.increaseStepsCounter()
	if alien.TargetCity != nil {
		alien.TargetCity.Free()
		alien.TargetCity = nil
	}
}

func (alien *Alien) invade(targetCity *grid.City) {
	if targetCity.IsInvaded() {
		fightOccupant(alien, targetCity)
	} else {
		conquerCity(alien, targetCity)
	}
}

func (alien *Alien) aimNextCity() *grid.City {
	if alien.Steps == 0 {
		return alien.TargetCity
	}
	if alien.TargetCity == nil {
		return nil
	}
	return alien.TargetCity.RandomDirection()
}

func (alien *Alien) die() {
	alien.dead = true
}

func (alien *Alien) trapped() {
	alien.idle = true
	fmt.Printf("Alien %s is trapped\n", alien.Name)
}

func conquerCity(attacker *Alien, targetCity *grid.City) {
	targetCity.Invade(attacker.Name)
	attacker.TargetCity = targetCity
}

func fightOccupant(attacker *Alien, targetCity *grid.City) {
	occupant := FindInvaderOf(targetCity)
	if occupant != nil {
		// battle begins!
		finalClash(attacker, occupant, targetCity)
	}
}

func finalClash(attacker *Alien, defensor *Alien, city *grid.City) {
	grid.DestroyCity(city)
	defensor.die()
	attacker.die()
	fmt.Printf("%s has been destroyed by alien %s and alien %s!\n", city.Name, attacker.Name, defensor.Name)
}
