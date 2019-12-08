package application

import (
	"fmt"
	"html/template"
	"log"

	"github.com/gorilla/mux"
)

/*Initialize app router and configuration*/
func (goservice *GoService) Initialize() {
	goservice.Router = mux.NewRouter()

	templates, err := template.ParseGlob("./templates/*.tmpl")
	if err != nil {
		log.Println(fmt.Sprintf("%q", err))
	}
	goservice.Indexlayout = templates
	//initTemplate(template.New("./templates/layout"))
	//initTemplate(goservice.Indexlayout.New("./templates/event"))

	goservice.Router.HandleFunc("/event", goservice.PostCreate).Methods("POST")
	goservice.Router.HandleFunc("/", goservice.GetAll).Methods("GET")

	goservice.Router.Use(loggingMiddleware)

	var appconfig ConfigParameters = handleConfiguration()
	goservice.Config = &appconfig
}

func handleConfiguration() ConfigParameters {
	var appconfig ConfigParameters

	appconfig.ServiceURL = "http://127.0.0.1:8080"
	return appconfig
}
