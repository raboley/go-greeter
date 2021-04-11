package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello! Greeter has started!, and is listening on port 80 for /hello or /greet/<host name>")
	http.HandleFunc("/hello/", HelloServer)
	http.HandleFunc("/greet/", GreetOtherServer)
	http.HandleFunc("/", Healthy)
	http.ListenAndServe(":80", nil)
}

func Healthy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I'm healthy!")

}

func GreetOtherServer(w http.ResponseWriter, r *http.Request) {
	serverName := r.URL.Path[7:]

	client := &http.Client{}
	resp, err := client.Get(fmt.Sprintf("http://%s/", serverName))
	if err != nil {
		fmt.Fprintf(w, "server  %s returned an error: %v", serverName, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Fprintf(w, "server %s says: %s", serverName, body)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[7:])
}
