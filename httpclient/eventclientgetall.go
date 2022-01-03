package httpclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	model "github.com/wlanboy/gowebfrontend/v2/model"
)

/*LoadEvents using event service*/
func LoadEvents() *model.SingletonEventStore {

	var netClient = ConfigureHTTPClient()

	var datastore model.SingletonEventStore
	var eventlist []model.Event

	response, _ := netClient.Get("http://127.0.0.1:8000/api/v1/event")

	body, errio := ioutil.ReadAll(response.Body)
	if errio != nil {
		log.Println("Cloud Config read error")
		log.Fatal(errio)
	} else {
		defer response.Body.Close()
		errjson := json.Unmarshal(body, &eventlist)

		if errjson != nil {
			fmt.Println(errjson)
			log.Fatal("Cloud Config json unmarshal error")
		}
	}

	datastore.EventDataList = eventlist

	return &datastore
}
