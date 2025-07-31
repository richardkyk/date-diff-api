package main

import (
	"log"
	"net/http"
	"strconv"

	"date-diff-api/internal/datediff"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		start := req.URL.Query().Get("start")
		end := req.URL.Query().Get("end")
		units := req.URL.Query().Get("units")

		input, err := datediff.ParseInput(start, end, units)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		diff := datediff.CalculateDateDiff(input)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		if _, err := w.Write([]byte(strconv.Itoa(diff))); err != nil {
			log.Println("Error writing response:", err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
