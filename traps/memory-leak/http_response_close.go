// COPY FROM: https://husobee.github.io/golang/memory/leak/2016/02/11/go-mem-leak.html
package memory_leak

import (
	"io/ioutil"
	"log"
	"net/http"
)

var done = make(chan bool)

func WrongHttpBodyClose() {
	defer func() { done <- true }()
	// perform a web request
	resp, err := http.Get("https://cip.cc")
	if err != nil {
		log.Fatal(err)
	}
	// check the status code of the response
	// if it returns ok, in our example,
	// we dont care about the body, but if
	// not okay, then we need to read the body
	if resp.StatusCode != http.StatusOK {
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
	}
}

func GoodHttpBodyClose() {
	// signal we are done doing something
	defer func() { done <- true }()
	// perform a web request
	resp, err := http.Get("https://husobee.github.io/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// close it on defer
	// check the status code of the response
	// if it returns ok, in our example,
	// we dont care about the body, but if
	// not okay, then we need to read the body
	if resp.StatusCode != http.StatusOK {
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
