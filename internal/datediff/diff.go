package datediff

import (
	"time"
)

func CalculateDateDiff(start, end time.Time, unit string) int {
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
