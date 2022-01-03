package application

import (
	"log"
	"net/http"
	"os"

	httpclient "github.com/wlanboy/gowebfrontend/v2/httpclient"
	model "github.com/wlanboy/gowebfrontend/v2/model"
	store "github.com/wlanboy/gowebfrontend/v2/store"
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
