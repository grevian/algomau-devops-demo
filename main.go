package main

import (
	"context"
	"fmt"
	"html"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	srv := runWebserver()

	// Setting up signal capturing, when the process is killed the webserver shuts down
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// When we receive a KILL signal, shutdown the webserver and exit
	<-stop
	_ = srv.Shutdown(context.TODO())
}

func runWebserver() *http.Server {

	// Create an http server and a router to map requests to functions to fulfil them
	server := &http.Server{Addr: ":8080"}
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", reverseString(html.EscapeString(r.URL.Path)))
	})

	router.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	server.Handler = router

	// Run the server in a goroutine
	go func() {
		_ = server.ListenAndServe()
	}()

	return server
}

// Reverse any passed in string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
