package libs

import (
	"fmt"
	"time"
)

func ParseToDateTime(sDate string) *time.Time {
	adate, err := time.Parse(time.DateTime, sDate)
	if err != nil {
		return nil
	}

	return &adate
}

func ParseToDate(sDate string) *time.Time {
	adate, err := time.Parse(time.DateOnly, sDate)
	if err != nil {
		return nil
	}

	return &adate
}

func ParseToTime(sTime string) *time.Time {
	atime, err := time.Parse(time.TimeOnly, sTime)
	if err != nil {
		return nil
	}

	return &atime
}

func ParseToTimeString(sDateTime string) string {
	atime, err := time.Parse(time.TimeOnly, sDateTime)
	if err != nil {
		return time.Now().Format(time.TimeOnly)
	}
	dtime := fmt.Sprintf("%d:%d:%d", atime.Hour(), atime.Minute(), atime.Second())

	return dtime
}

func FormatDate(sDate time.Time) *string {
	adate := sDate.Format(time.DateOnly)

	return &adate
}

func FormatTime(sDate time.Time) string {
	adate := sDate.Format(time.TimeOnly)

	return adate
}

func FormatDateTime(sDate time.Time) *string {
	adate := sDate.Format(time.DateTime)

	return &adate
}
