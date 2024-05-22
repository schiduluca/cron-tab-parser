package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PartType int

const (
	Minute = PartType(iota)
	Hour
	DayOfMonth
	Month
	DayOfWeek
)

var partValues = map[PartType][]int{
	Minute:     {0, 59},
	Hour:       {0, 23},
	DayOfMonth: {1, 31},
	Month:      {1, 12},
	DayOfWeek:  {1, 7},
}

var outputValues = []string{
	Minute:     "minute        ",
	Hour:       "hour          ",
	DayOfMonth: "day of month  ",
	Month:      "month         ",
	DayOfWeek:  "day of week   ",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No input provided.")
		return
	}

	split := strings.Split(os.Args[1], " ")
	if len(split) != 6 {
		fmt.Println("Invalid input provided.")
		return
	}

	err := parseCronExpression(split[:5])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("command      ", split[5])
}

func parseCronExpression(parts []string) error {
	for i, part := range parts {
		result, err := parsePart(PartType(i), part)
		if err != nil {
			return err
		}
		printOutput(outputValues[i], result)
	}

	return nil
}

func parsePart(t PartType, part string) ([]int, error) {
	if part == "*" {
		return makeRangeWithStep(partValues[t][0], partValues[t][1], 1), nil
	}

	num, err := strconv.Atoi(part)
	if err == nil {
		return []int{num}, nil
	}

	var result []int

	if strings.Contains(part, ",") {
		parts := strings.Split(part, ",")
		for _, p := range parts {
			tempResult, err := parsePart(t, p)
			if err != nil {
				return nil, err
			}
			result = append(result, tempResult...)
		}
		return result, nil
	}

	if strings.Contains(part, "/") {
		parts := strings.Split(part, "/")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid cron expression: %s", part)
		}
		step, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid cron expression: %s", parts[1])
		}

		if strings.Contains(parts[0], "-") {
			low, high, err := getRangeValues(parts[0])
			if err != nil {
				return nil, err
			}
			result = append(result, makeRangeWithStep(low, high, step)...)
		} else {
			if parts[0] == "*" {
				result = append(result, makeRangeWithStep(partValues[t][0], partValues[t][1], step)...)
				return result, nil
			}
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("invalid cron expression: %s", parts[1])
			}
			result = append(result, makeRangeWithStep(start, partValues[t][1], step)...)
		}

		return result, nil
	}

	if strings.Contains(part, "-") {
		low, high, err := getRangeValues(part)
		if err != nil {
			return nil, err
		}
		return makeRangeWithStep(low, high, 1), nil
	}

	return result, nil
}

func getRangeValues(part string) (int, int, error) {
	ranges := strings.Split(part, "-")
	low, err := strconv.Atoi(ranges[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid range value")
	}
	high, err := strconv.Atoi(ranges[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid range value")
	}
	return low, high, nil
}

func makeRangeWithStep(min, max, step int) []int {
	var a []int
	for i := min; i <= max; i += step {
		a = append(a, i)
	}
	return a
}

func printOutput(column string, s []int) {
	fmt.Print(column)
	for _, value := range s {
		fmt.Printf("%v ", value)
	}
	fmt.Println()
}
