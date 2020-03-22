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

func addTimes(values []string) {
	totalTime := time.Now()
	for _, entry := range values {
		entry = strings.ToLower(entry)
		if strings.Contains(entry, "d") {
			entrySlice := strings.Split(entry, "d")
			days := parseDays(entrySlice[0])
			totalTime = totalTime.AddDate(0, 0, days)
			hours := parseHours(entrySlice[1])
			totalTime = totalTime.Add(hours)
		} else {
			if strings.Contains(entry, "h") {
				duration := parseHours(entry)
				totalTime = totalTime.Add(duration)
			}
		}
	}

	fmt.Println(totalTime.Format(time.RFC822))

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var vals []string
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
		vals = []string{}
	}
}
