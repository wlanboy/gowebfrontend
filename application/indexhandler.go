package application

import (
	"log"
	"net/http"

	store "../store"
)

/*GetAll GET method*/
func (goservice *GoService) GetAll(w http.ResponseWriter, r *http.Request) {

	datastore := store.GetInstance()

	err := goservice.Indexlayout.ExecuteTemplate(w, "layout", datastore)
	if err != nil {
		log.Println("Model error")
		log.Println(err)
	}
}
