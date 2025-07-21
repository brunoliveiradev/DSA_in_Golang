package warmups

import "strconv"

func TimeConversion(s string) string {
	// Check if the input string is in the correct format
	hour := s[:2]
	minute := s[3:5]
	second := s[6:8]
	period := s[8:]

	if period == "PM" && hour != "12" {
		convertedHour, _ := strconv.Atoi(hour)
		convertedHour += 12
		hour = strconv.Itoa(convertedHour)
	} else if period == "AM" && hour == "12" {
		hour = "00"
	}

	return hour + ":" + minute + ":" + second
}
