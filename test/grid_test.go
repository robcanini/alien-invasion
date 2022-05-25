package test

import (
	"github.com/robcanini/alien-invasion/internal/grid"
	"testing"
)

func TestShouldNotBeDestroyedCreateCity(t *testing.T) {
	city := grid.CreateCity("Foo")
	if city.Destroyed {
		t.Fatal("city should not be destroyed yet")
	}
}

func TestShouldNotBeInvadedCreateCity(t *testing.T) {
	city := grid.CreateCity("Foo")
	if city.IsInvaded() {
		t.Fatal("city should not be invaded yet")
	}
}

func TestShouldBeInvadedAfterInvade(t *testing.T) {
	city := grid.CreateCity("Foo")
	city.Invade("Bar")
	if !city.IsInvaded() {
		t.Fatal("city should be invaded")
	}
}

func TestShouldBeFreeAfterInvadeAndFree(t *testing.T) {
	city := grid.CreateCity("Foo")
	city.Invade("Bar")
	city.Free()
	if city.IsInvaded() {
		t.Fatal("city should not be invaded")
	}
}

func TestShouldBeIdempotentFree(t *testing.T) {
	city := grid.CreateCity("Foo")
	city.Invade("Bar")
	city.Free()
	city.Free()
	if city.IsInvaded() {
		t.Fatal("city should not be invaded")
	}
}
