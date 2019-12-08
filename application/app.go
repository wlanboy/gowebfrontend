package application

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

/*ConfigParameters for App*/
type ConfigParameters struct {
	ServiceURL string
}

/*GoService containing router and database*/
type GoService struct {
	Router      *mux.Router
	Server      *http.Server
	Config      *ConfigParameters
	Indexlayout *template.Template
}

/*Run app and initialize http server*/
func (goservice *GoService) Run() {

	//load http server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8100"
	}
	fmt.Println(port)

	goservice.Server = &http.Server{
		Handler:      goservice.Router,
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Starting http server...")
		if err := goservice.Server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

}

/*WaitForShutdown application server*/
func (goservice *GoService) WaitForShutdown() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	goservice.Server.Shutdown(ctx)

	log.Println("Shutting down http server.")
	os.Exit(0)
}
