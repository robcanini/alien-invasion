package test

import (
	"github.com/robcanini/alien-invasion/internal/aliens"
	"github.com/robcanini/alien-invasion/internal/grid"
	"testing"
)

func TestShouldNotBeDeadCreateAlien(t *testing.T) {
	city := grid.CreateCity("Foo")
	alien := aliens.CreateAlien(city)
	if alien.Dead {
		t.Fatal("alien should not be dead")
	}
}

func TestShouldNotBeInIdleCreateAlien(t *testing.T) {
	city := grid.CreateCity("Foo")
	alien := aliens.CreateAlien(city)
	if alien.Idle {
		t.Fatal("alien should not be in idle state")
	}
}

func TestShouldNameBeEvaluatedCreateAlien(t *testing.T) {
	city := grid.CreateCity("Foo")
	alien := aliens.CreateAlien(city)
	if len(alien.Name) == 0 {
		t.Fatal("alien name should be evalauated")
	}
}
