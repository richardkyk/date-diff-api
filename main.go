package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func CalculateDiff(start, end time.Time, unit string) int {
	switch unit {
	case "days":
		duration := end.Sub(start)
		return int(duration.Hours() / 24)
	case "weeks":
		duration := end.Sub(start)
		return int(duration.Hours() / 24 / 7)
	case "months":
		duration := end.Sub(start)
		return int(duration.Hours() / 24 / 30)
	default: // "years" or any other value defaults to years
		diff := end.Year() - start.Year()
		if end.YearDay() < start.YearDay() {
			diff--
		}
		return diff
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		start := req.URL.Query().Get("start")
		end := req.URL.Query().Get("end")
		unit := req.URL.Query().Get("unit")

		if start == "" {
			http.Error(w, "Missing start date", http.StatusBadRequest)
			return
		}

		// Parse startTime
		startTime, err := time.Parse("2006-01-02", start)
		if err != nil {
			http.Error(w, "Invalid date format, use YYYY-MM-DD", http.StatusBadRequest)
			return
		}

		// Parse endTime or default to today
		var endTime time.Time
		if end == "" {
			endTime = time.Now()
		} else {
			endTime, err = time.Parse("2006-01-02", end)
			if err != nil {
				http.Error(w, "Invalid end date format, use YYYY-MM-DD", http.StatusBadRequest)
				return
			}
		}

		diff := CalculateDiff(startTime, endTime, unit)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		if _, err := w.Write([]byte(strconv.Itoa(diff))); err != nil {
			log.Println("Error writing response:", err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
