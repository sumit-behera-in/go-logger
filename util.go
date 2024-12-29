package goLogger

import (
	"time"
	"fmt"
)

func loadTimeZone(timezone string) (*time.Location, string) {
	var location *time.Location
	var err error

	// Mapping timezone abbreviations to specific time zones
	switch timezone {
	case "ACST":
		location, err = time.LoadLocation("Australia/Adelaide")
	case "AEST":
		location, err = time.LoadLocation("Australia/Sydney")
	case "AKST":
		location, err = time.LoadLocation("America/Anchorage")
	case "AST":
		location, err = time.LoadLocation("Asia/Riyadh")
	case "AWST":
		location, err = time.LoadLocation("Australia/Perth")
	case "BST":
		location, err = time.LoadLocation("Europe/London")
	case "CCT":
		location, err = time.LoadLocation("Asia/Shanghai")
	case "CDT":
		location, err = time.LoadLocation("America/Chicago")
	case "CET":
		location, err = time.LoadLocation("Europe/Paris")
	case "CST":
		location, err = time.LoadLocation("America/Chicago")
	case "EAT":
		location, err = time.LoadLocation("Africa/Nairobi")
	case "EDT":
		location, err = time.LoadLocation("America/New_York")
	case "EET":
		location, err = time.LoadLocation("Europe/Bucharest")
	case "EST":
		location, err = time.LoadLocation("America/New_York")
	case "GMT":
		location, err = time.LoadLocation("Europe/London")
	case "HKT":
		location, err = time.LoadLocation("Asia/Hong_Kong")
	case "HST":
		location, err = time.LoadLocation("Pacific/Honolulu")
	case "IST":
		location, err = time.LoadLocation("Asia/Calcutta")
	case "JST":
		location, err = time.LoadLocation("Asia/Tokyo")
	case "KST":
		location, err = time.LoadLocation("Asia/Seoul")
	case "MDT":
		location, err = time.LoadLocation("America/Denver")
	case "MSK":
		location, err = time.LoadLocation("Europe/Moscow")
	case "MST":
		location, err = time.LoadLocation("America/Denver")
	case "NZST":
		location, err = time.LoadLocation("Pacific/Auckland")
	case "PDT":
		location, err = time.LoadLocation("America/Los_Angeles")
	case "PST":
		location, err = time.LoadLocation("America/Los_Angeles")
	case "SAST":
		location, err = time.LoadLocation("Africa/Johannesburg")
	case "SGT":
		location, err = time.LoadLocation("Asia/Singapore")
	case "UTC":
		location, err = time.LoadLocation("UTC")
	case "WAT":
		location, err = time.LoadLocation("Africa/Lagos")
	default:
		// Fallback to local time zone if invalid timezone is given
		location = time.Local
	}

	// Handle error if timezone is invalid
	if err != nil && location == nil {
		// Default to local time and provide a helpful message
		location = time.Local
		fmt.Printf("Warning: invalid or unknown timezone abbreviation '%s'. Defaulting to local time.\n", timezone)
	}

	// Get the current time in the chosen location and format the timezone string
	currentTime := time.Now().In(location)
	timezone = currentTime.Format("MST") // Returns a formatted timezone string (e.g., "PST", "EST", etc.)

	return location, timezone
}
