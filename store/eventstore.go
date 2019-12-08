package store

import (
	"sync"
	"sync/atomic"

	model "../model"
)

var initialized uint32
var mu sync.Mutex

/*Singleton containing array of model event*/
type Singleton struct {
	EventDataList []model.Event
}

var instance *Singleton

/*GetInstance of singelton holding the array of events*/
func GetInstance() *Singleton {

	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &Singleton{}
		instance.EventDataList = []model.Event{
			model.Event{Name: "Hello World!", Type: "CATA"},
			model.Event{Name: "Hello World again!", Type: "CATB"},
		}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

/*AddEvent to singelton list*/
func AddEvent(newevent model.Event) {

	if initialized == 0 {
		instance = &Singleton{}
		instance.EventDataList = []model.Event{
			model.Event{Name: "Hello World!", Type: "CATA"},
			model.Event{Name: "Hello World again!", Type: "CATB"},
		}
		atomic.StoreUint32(&initialized, 1)
	}
	instance.EventDataList = append(instance.EventDataList, newevent)
}
