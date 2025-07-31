package main

import (
	"flag"
	"fmt"
	"os"

	"date-diff-api/internal/datediff"
)

func main() {
	// Define flags
	endFlag := flag.String("e", "", "End date (YYYY-MM-DD). Defaults to today.")
	unitsFlag := flag.String("u", "years", "Units (days, weeks, months, years). Defaults to 'years'.")

	flag.Usage = func() {
		fmt.Println("Usage: cli [flags] <start-date>")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Get positional arguments (after flags)
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	start := args[0]

	// Validate units flag
	validUnits := map[string]bool{
		"days":   true,
		"weeks":  true,
		"months": true,
		"years":  true,
	}

	if !validUnits[*unitsFlag] {
		fmt.Printf("Error: invalid units '%s'. Must be one of: days, weeks, months, years\n", *unitsFlag)
		os.Exit(1)
	}

	input, err := datediff.ParseInput(start, *endFlag, *unitsFlag)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	diff := datediff.CalculateDateDiff(input)

	fmt.Printf(
		"The difference between %s and %s is %.1d %s.\n",
		input.Start.Format("Jan 2, 2006"),
		input.End.Format("Jan 2, 2006"),
		diff,
		input.Units,
	)
}
