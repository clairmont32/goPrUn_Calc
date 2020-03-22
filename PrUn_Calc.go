package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseHours(hours string) time.Duration {
	dur, parseErr := time.ParseDuration(hours)
	if parseErr != nil {
		log.Println(parseErr)
	}
	return dur
}

func parseDays(entry string) int {
	entrySlice := strings.Split(entry, "d")
	days, convErr := strconv.Atoi(entrySlice[0])
	if convErr != nil {
		log.Fatalln("Could not convert days to int")
	}
	return days
}

// add all input times together
func addTimes(values []string) {
	totalTime := time.Now()
	for _, entry := range values {
		entry = strings.ToLower(entry)

		// if days were entered, process them separately from the hours
		if strings.Contains(entry, "d") {
			entrySlice := strings.Split(entry, "d")
			days := parseDays(entrySlice[0])

			totalTime = totalTime.AddDate(0, 0, days)
			hours := parseHours(entrySlice[1])
			totalTime = totalTime.Add(hours)

		// add hours
		} else {
				duration := parseHours(entry)
				totalTime = totalTime.Add(duration)

		}
	}

	fmt.Println(totalTime.Format(time.RFC822))

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var vals []string

	// read a series of inputs, append them to a slice
	for {
		fmt.Println("Enter the times one at a time in 0h0m format")
		for scanner.Scan() {
			if scanner.Text() == "" {
				break
			} else if scanner.Text() == "q" {
				os.Exit(0)
			}

			vals = append(vals, scanner.Text())
		}

		addTimes(vals)
		vals = []string{} // clear slice to avoid constant time additions
	}
}
