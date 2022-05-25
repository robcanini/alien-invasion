package test

import (
	"github.com/robcanini/alien-invasion/internal/utils"
	"sync"
	"testing"
)

var mut = &sync.Mutex{}

func TestShouldBeUnlockedMutex(t *testing.T) {
	if utils.IsMutexLocked(mut) {
		t.Fatal("shouldBeUnlockedMutex")
	}
}

func TestShouldBeLockedMutex(t *testing.T) {
	mut.Lock()
	defer mut.Unlock()
	if !utils.IsMutexLocked(mut) {
		t.Fatal("shouldBeLockedMutex")
	}
}
