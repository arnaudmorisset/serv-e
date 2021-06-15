package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/record", createRecord)
	http.HandleFunc("/records", getRecords)

	fmt.Println("server listening on http://localhost:80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Fprintf(os.Stderr, "server stopped for reason: %v", err)
		os.Exit(1)
	}
}

func createRecord(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

func getRecords(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte("<h1>Logger</h1>"))
}
