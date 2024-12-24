package goLogger

import "time"

func loadTimeZone(timezone string) (*time.Location, string) {
	var location *time.Location
	switch timezone {
	case "ACST":
		location, _ = time.LoadLocation("Australia/Adelaide")
	case "AEST":
		location, _ = time.LoadLocation("Australia/Sydney")
	case "AKST":
		location, _ = time.LoadLocation("America/Anchorage")
	case "AST":
		location, _ = time.LoadLocation("Asia/Riyadh")
	case "AWST":
		location, _ = time.LoadLocation("Australia/Perth")
	case "BST":
		location, _ = time.LoadLocation("Europe/London")
	case "CCT":
		location, _ = time.LoadLocation("Asia/Shanghai")
	case "CDT":
		location, _ = time.LoadLocation("America/Chicago")
	case "CET":
		location, _ = time.LoadLocation("Europe/Paris")
	case "CST":
		location, _ = time.LoadLocation("America/Chicago")
	case "EAT":
		location, _ = time.LoadLocation("Africa/Nairobi")
	case "EDT":
		location, _ = time.LoadLocation("America/New_York")
	case "EET":
		location, _ = time.LoadLocation("Europe/Bucharest")
	case "EST":
		location, _ = time.LoadLocation("America/New_York")
	case "GMT":
		location, _ = time.LoadLocation("Europe/London")
	case "HKT":
		location, _ = time.LoadLocation("Asia/Hong_Kong")
	case "HST":
		location, _ = time.LoadLocation("Pacific/Honolulu")
	case "IST":
		location, _ = time.LoadLocation("Asia/Calcutta")
	case "JST":
		location, _ = time.LoadLocation("Asia/Tokyo")
	case "KST":
		location, _ = time.LoadLocation("Asia/Seoul")
	case "MDT":
		location, _ = time.LoadLocation("America/Denver")
	case "MSK":
		location, _ = time.LoadLocation("Europe/Moscow")
	case "MST":
		location, _ = time.LoadLocation("America/Denver")
	case "NZST":
		location, _ = time.LoadLocation("Pacific/Auckland")
	case "PDT":
		location, _ = time.LoadLocation("America/Los_Angeles")
	case "PST":
		location, _ = time.LoadLocation("America/Los_Angeles")
	case "SAST":
		location, _ = time.LoadLocation("Africa/Johannesburg")
	case "SGT":
		location, _ = time.LoadLocation("Asia/Singapore")
	case "UTC":
		location, _ = time.LoadLocation("UTC")
	case "WAT":
		location, _ = time.LoadLocation("Africa/Lagos")
	default:
		location = time.Local
		currentTime := time.Now().In(location)
		timezone = currentTime.Format("IST")
	}
	return location, timezone
}
