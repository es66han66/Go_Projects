package main

import (
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func makeGetRequest(client *http.Client) error {
	req, err := http.NewRequest("GET", "http://example.com?q=one", nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	//=================================================
	// CLOSE THE RESPONSE BODY
	//=================================================
	if resp != nil {
		defer resp.Body.Close() // MUST CLOSED THIS
	}
	if err != nil {
		return err
	}
	//=================================================
	// READ THE BODY EVEN THE DATA IS NOT IMPORTANT
	// THIS MUST TO DO, TO AVOID MEMORY LEAK WHEN REUSING HTTP
	// CONNECTION
	//=================================================
	_, err = io.Copy(ioutil.Discard, resp.Body) // WE READ THE BODY
	if err != nil {
		return err
	}
	return nil
}

func main() {
	keepAliveTimeout := 600 * time.Second
	timeout := 2 * time.Second
	defaultTransport := &http.Transport{
		Dial: (&net.Dialer{
			KeepAlive: keepAliveTimeout,
		}).Dial,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}
	client := &http.Client{
		Transport: defaultTransport,
		Timeout:   timeout,
	}
	makeGetRequest(client)
}
