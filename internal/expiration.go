package internal

import (
	"fmt"
	"time"
)

// calcTime takes an expiration time in unix timestamp and returns a string
// representing the time until expiration
func CalcTime(expiration int) string {
	now := time.Now()
	// convert expiration to seconds
	futureTimestamp := int64(expiration)
	futureTime := time.Unix(futureTimestamp, 0)
	duration := futureTime.Sub(now)

	// Convert to minutes or hours as needed and print
	hours := int(duration.Hours())
	minutes := int(duration.Minutes())

	if minutes > 59 {
		hours = hours + 1
	}
	if hours > 0 {
		result := fmt.Sprintf("%d hours\n", hours)
		return result
	} else if minutes > 0 {
		result := fmt.Sprintf("%d minutes\n", minutes)
		return result
	} else {
		result := "No Expiration"
		return result
	}
}

// calcUnixTime takes a human readable time and returns a unix timestamp as int
// for example 1h or 30m
func CalcUnixTime(expiration string) int64 {
	// Parse the expiration time
	expirationTime, err := time.ParseDuration(expiration)
	if err != nil {
		fmt.Println("Error parsing time")
	}
	// Convert to unix timestamp
	unixTime := time.Now().Add(expirationTime).Unix()
	return int64(unixTime)
}
