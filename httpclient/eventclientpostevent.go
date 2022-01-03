package httpclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	model "github.com/wlanboy/gowebfrontend/v2/model"
)

/*AddEvent using event service*/
func AddEvent(newevent model.Event) {

	var netClient = ConfigureHTTPClient()

	jsonValue, _ := json.Marshal(newevent)

	req, reqerr := http.NewRequest("POST", "http://127.0.0.1:8000/api/v1/event", bytes.NewBuffer(jsonValue))

	if reqerr != nil {
		log.Println("HTTP Request error")
		log.Fatal(reqerr)
	} else {
		req.Header.Set("Content-Type", "application/json")

		response, _ := netClient.Do(req)
		log.Println(response.StatusCode)
	}
}
