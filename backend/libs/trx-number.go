package libs

import (
	"fmt"
	"time"
)

func GenerateTrxNumber() string {
	time := time.Now()
	y := time.Year()
	m := int(time.Month())
	d := time.Day()
	h := time.Hour()
	n := time.Minute()
	s := time.Second()

	number := y + m + d + h + n + s

	return fmt.Sprintf("%d", number)
}
