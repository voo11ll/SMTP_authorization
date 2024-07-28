package utils

import (
	"auth/auth_back/pkg/globalvars"
	"fmt"
	"time"
)

func ReturnDateString(d time.Time, sep string, useTime bool) string {
	day := fmt.Sprint(d.Day())
	month := fmt.Sprint(d.Month())
	year := fmt.Sprint(d.Year())

	hour := fmt.Sprint(d.Hour())
	_minuteInt := d.Minute()
	_secondsInt := d.Second()

	if _monthInt := globalvars.Months[month]; _monthInt < 10 {
		month = "0" + fmt.Sprint(_monthInt)
	} else {
		month = fmt.Sprint(_monthInt)
	}

	minute := fmt.Sprint(_minuteInt)

	if _minuteInt < 10 {
		minute = "0" + fmt.Sprint(_minuteInt)
	}

	seconds := fmt.Sprint(_secondsInt)

	if _secondsInt < 10 {
		seconds = "0" + fmt.Sprint(_secondsInt)
	}

	result := day + sep + month + sep + year

	if useTime {
		result = result + " " + hour + ":" + minute + ":" + seconds
	}

	return result
}
