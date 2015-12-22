package examples

import (
	"fmt"
	"log"
	"net/http"
)

func StartWebServer() {
	http.HandleFunc("/hello", handleHello)
	fmt.Println("Server on http://localhost:7777/hello")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "<html>Hello, <b>User</b>:)<html")
}
