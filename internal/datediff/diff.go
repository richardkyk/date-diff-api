package datediff

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

type Input struct {
	Start time.Time
	End   time.Time
	Units string
}

// ParseInput takes raw strings and returns parsed and validated Input
func ParseInput(startStr, endStr, unitsStr string) (Input, error) {
	if startStr == "" {
		return Input{}, errors.New("start date must be provided")
	}

	start, err1 := time.Parse(dateLayout, startStr)
	if err1 != nil {
		return Input{}, fmt.Errorf("invalid start date format (expected YYYY-MM-DD)")
	}

	var end time.Time
	if endStr == "" {
		end = time.Now()
	} else {
		endParsed, err2 := time.Parse(dateLayout, endStr)
		if err2 != nil {
			return Input{}, fmt.Errorf("invalid end date format (expected YYYY-MM-DD)")
		}
		end = endParsed
	}

	units := strings.ToLower(unitsStr)

	return Input{
		Start: start,
		End:   end,
		Units: units,
	}, nil
}

func CalculateDateDiff(input Input) int {
	start, end, units := input.Start, input.End, input.Units

	switch units {
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
