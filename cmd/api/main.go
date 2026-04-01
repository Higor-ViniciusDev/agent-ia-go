package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello mundo")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"okay"}`))
	})

	http.ListenAndServe(":8000", nil)
}
