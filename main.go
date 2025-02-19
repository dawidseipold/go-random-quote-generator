package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/quote", handleQuote)

	err := http.ListenAndServe(":8080", nil)
	handleError(err)
}
