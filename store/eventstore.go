package store

import (
	"sync"
	"sync/atomic"

	model "github.com/wlanboy/gowebfrontend/v2/model"
)

var initialized uint32
var mu sync.Mutex

var instance *model.SingletonEventStore

/*GetInstance of singelton holding the array of events*/
func GetInstance() *model.SingletonEventStore {

	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &model.SingletonEventStore{}
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
		instance = &model.SingletonEventStore{}
		instance.EventDataList = []model.Event{
			model.Event{Name: "Hello World!", Type: "CATA"},
			model.Event{Name: "Hello World again!", Type: "CATB"},
		}
		atomic.StoreUint32(&initialized, 1)
	}
	instance.EventDataList = append(instance.EventDataList, newevent)
}
