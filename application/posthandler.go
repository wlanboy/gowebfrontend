package application

import (
	"log"
	"net/http"
	"os"

	httpclient "github.com/wlanboy/gowebfrontend/v2/httpclient"
	model "github.com/wlanboy/gowebfrontend/v2/model"
	store "github.com/wlanboy/gowebfrontend/v2/store"
)

/*PostCreate POST method*/
func (goservice *GoService) PostCreate(w http.ResponseWriter, r *http.Request) {

	Name := r.FormValue("Name")
	Type := r.FormValue("Type")

	newevent := model.Event{}
	newevent.Name = Name
	newevent.Type = Type

	var datastore *model.SingletonEventStore

	test := os.Getenv("TEST")
	if test == "" {
		httpclient.AddEvent(newevent)
		datastore = httpclient.LoadEvents()
	} else {
		store.AddEvent(newevent)
		datastore = store.GetInstance()
	}

	err := goservice.Indexlayout.ExecuteTemplate(w, "layout", datastore)
	if err != nil {
		log.Println("Model error")
		log.Println(err)
	}
}
