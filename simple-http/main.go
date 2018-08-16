package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

var count int64
var counterActive = false
var startTime = time.Now().In(time.Local)
var buildTime time.Time

func hello(w http.ResponseWriter, r *http.Request) {

	count++
	hostname, _ := os.Hostname()

	fmt.Fprintf(w,
		`<html><pre>
       Hostname:          %s
       Requests:          %d
       Server Build Time: %s
       Server Start Time: %s
       Current Time:      %s`,
		hostname, count, formatTime(buildTime), formatTime(startTime), formatTime(time.Now().In(time.Local)))

	if counterActive {
		response, err := http.Get("http://localhost:8181/inc")
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			fmt.Fprintf(w, "\n       Global Requests:   %s", string(contents))
		}
	}
}

func throwAway(w http.ResponseWriter, r *http.Request) {
	// Do nothing
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/favicon.ico", throwAway)
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":8080", mux)
}

func formatTime(t time.Time) string {
	return fmt.Sprintf("%d/%d/%d %.02d:%.02d:%.02d", t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute(), t.Second())
}

func init() {
	s, _ := os.Stat(os.Args[0])
	buildTime = s.ModTime().In(time.Local)

	timeout := time.Duration(1 * time.Second)
	_, err := net.DialTimeout("tcp", "localhost:8181", timeout)
	if err == nil {
		counterActive = true
	}
}
