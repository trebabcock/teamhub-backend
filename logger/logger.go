package logger

import (
	"fmt"
	"time"
)

// Log logs the message
func Log(message string) {
	output := fmt.Sprintf("%s %s", genDate(), message)
	fmt.Println(output)
}

func genDate() string {
	month := "0"
	if time.Now().Month() < 10 {
		month = fmt.Sprintf("%d%d", 0, time.Now().Month())
	} else {
		month = fmt.Sprintf("%d", time.Now().Month())
	}
	day := "0"
	if time.Now().Day() < 10 {
		day = fmt.Sprintf("%d%d", 0, time.Now().Day())
	} else {
		day = fmt.Sprintf("%d", time.Now().Day())
	}
	year := time.Now().Year()
	hour := "0"
	if time.Now().Hour() < 10 {
		hour = fmt.Sprintf("%d%d", 0, time.Now().Hour())
	} else {
		hour = fmt.Sprintf("%d", time.Now().Hour())
	}
	minute := "0"
	if time.Now().Minute() < 10 {
		minute = fmt.Sprintf("%d%d", 0, time.Now().Minute())
	} else {
		minute = fmt.Sprintf("%d", time.Now().Minute())
	}
	second := "0"
	if time.Now().Second() < 10 {
		second = fmt.Sprintf("%d%d", 0, time.Now().Second())
	} else {
		second = fmt.Sprintf("%d", time.Now().Second())
	}
	return fmt.Sprintf("%s/%s/%d %s:%s:%s", month, day, year, hour, minute, second)
}
