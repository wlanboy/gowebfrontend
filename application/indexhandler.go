package application

import (
	"log"
	"net/http"
	"os"

	httpclient "../httpclient"
	model "../model"
	store "../store"
)

/*GetAll GET method*/
func (goservice *GoService) GetAll(w http.ResponseWriter, r *http.Request) {

	var datastore *model.SingletonEventStore

	test := os.Getenv("TEST")
	if test == "" {
		datastore = httpclient.LoadEvents()
	} else {
		datastore = store.GetInstance()
	}

	err := goservice.Indexlayout.ExecuteTemplate(w, "layout", datastore)
	if err != nil {
		log.Println("Model error")
		log.Println(err)
	}
}
