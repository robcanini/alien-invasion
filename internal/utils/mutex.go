package utils

import (
	"reflect"
	"sync"
)

const mutexLocked = 1

func IsMutexLocked(m *sync.Mutex) bool {
	state := reflect.ValueOf(m).Elem().FieldByName("state")
	return state.Int()&mutexLocked == mutexLocked
}
