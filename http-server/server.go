package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	port := ":80"

	http.HandleFunc("/", rootHandler)
	http.Handle("/about/", aboutHandler{})
	http.HandleFunc("/post/", postHandler)

	fmt.Printf("Listening on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

type aboutHandler struct {
}

func (h aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World!</h1>")
	printRequest(r)
}

type jsonData struct {
	Name string
	Age  int
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	printRequest(r)
	fmt.Println("Received to /post/")
	var data jsonData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Fprintln(w, "Invalid json data")
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Fprintln(w, "Json data received")
	fmt.Printf("Json data: %v", data)
}

func printRequest(r *http.Request) {
	fmt.Printf("URL:  %s\n", r.URL)
	fmt.Printf("Host:  %s\n", r.Host)
	fmt.Printf("Method:  %s\n", r.Method)
	fmt.Printf("Header: %v\n", r.Header)
	fmt.Printf("Body: %v\n", r.Body)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response: This is the root\n")
	printRequest(r)
}
