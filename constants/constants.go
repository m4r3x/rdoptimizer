package constants

import (
	"fmt"
	"os"
	"strconv"
)

// RPS is Requests per second used to emulate traffic.
func RPS() int {
	if len(os.Args) < 2 {
		fmt.Println("Input warning: Could not read RPS from argv!")
		return 1
	}

	i, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Input warning: Could not read RPS from argv!")
		return 1
	}

	fmt.Printf("RPS: %v\n", i)

	return i
}

// MODE is value of that defines type of benchmark ran.
func MODE() int {
	if len(os.Args) < 3 {
		fmt.Println("Input warning: Could not read MODE from argv!")
		return 1
	}

	i, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Input warning: Could not read MODE from argv!")
		return 1
	}

	switch i {
	case 1:
		fmt.Println("Events scenario, case study #1 - raw Data")
	case 2:
		fmt.Println("Events scenario, case study #2 - map data")
	case 3:
		fmt.Println("Events scenario, case study #3 - raw Data + ipv6")
	case 4:
		fmt.Println("Events scenario, case study #4 - map data + ipv6 in list")
	case 5:
		fmt.Println("Raw text scenario, case study #1 - raw data.")
	case 6:
		fmt.Println("Raw text scenario, case study #2 - proto encoded data.")
	}
	fmt.Println()

	return i
}
