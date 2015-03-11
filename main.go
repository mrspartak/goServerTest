package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const resp = `<html>
    <head>
        <title>Simple Web App</title>
    </head>
    <body>
        <h1>Simple Web App</h1>
        <p>Hello World!</p>
    </body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("server", "Go")
	fmt.Fprintf(w, resp)
	fmt.Fprintf(w, r.URL.String())
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
