package utils

import (
	"fmt"
	"math"
)

var (
	unitlist = [4]string{"", "K", "M", "B"}
)

// AmountConverter converts a number to a string with a unit.
func AmountConverter(number float64) string {
	var sign = math.Copysign(-1, number)
	var unit = 0

	for math.Abs(number) > 1000 {
		unit = unit + 1
		number = math.Floor(math.Abs(number)/100) / 10
	}
	return fmt.Sprint(sign*math.Abs(number)) + unitlist[unit]
}
