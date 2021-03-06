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
	Dead       bool
	Idle       bool
}

const MaxIterations = 10000

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func CreateAlien(targetCity *grid.City) *Alien {
	alienName := generateName()
	return &Alien{
		Name:       alienName,
		Steps:      0,
		TargetCity: targetCity,
		Dead:       false,
		Idle:       false,
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
	for alien.Steps < MaxIterations && !alien.Dead && !alien.Idle {
		city := alien.aimNextCity()
		alien.leaveCurrentCity()
		// no directions available, alien trapped
		if city == nil {
			alien.trapped()
			break
		}
		alien.invade(city)
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
	alien.Dead = true
}

func (alien *Alien) trapped() {
	alien.Idle = true
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
