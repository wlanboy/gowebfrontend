package httpclient

import (
	"net"
	"net/http"
	"time"
)

/*ConfigureHTTPClient setting timeouts*/
func ConfigureHTTPClient() *http.Client {

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	return netClient
}
