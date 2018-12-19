package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	//todo: set port in argv
	port := 8080
	portstring := strconv.Itoa(port)

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(HomeHandler))

	// Start listing on a given port with these routes on this server.
	// (I think the server name can be set here too , i.e. "foo.org:8080")
	log.Print("Listening on port " + portstring + " ... ")
	err := http.ListenAndServe(":"+portstring, mux)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

// HomeHandler is the hello world entry point.
// could be tested whether server is alive or down.
func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world! request path is %s\n", req.URL.Path)
}
