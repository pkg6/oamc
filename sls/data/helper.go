package data

import (
	"fmt"
	"math"
	"strings"
)

func FormatBytes(size float64, unit string) string {
	units := map[string]float64{
		"B":  math.Pow(1024, 0),
		"KB": math.Pow(1024, 1),
		"MB": math.Pow(1024, 2),
		"GB": math.Pow(1024, 3),
		"TB": math.Pow(1024, 4),
		"PB": math.Pow(1024, 5),
	}
	if multiplier, ok := units[unit]; ok {
		formatted := size / multiplier
		if math.Floor(formatted) == formatted {
			return fmt.Sprintf("%.0f%s", formatted, unit)
		} else {
			return fmt.Sprintf("%.2f%s", formatted, unit)
		}
	}
	return fmt.Sprintf("%.2fB", size)
}

func FormatSeconds(seconds int, unit string) int {
	switch unit {
	case "hour":
		return seconds / 3600
	case "day":
		return seconds / 86400
	case "millisecond":
		return seconds * 1000
	default:
		return 0
	}
}

func TitleContextLine(title string, length int, symbol, context string) string {
	titleLength := len(title)
	symbolLength := length - titleLength
	halfSymbolLength := symbolLength / 2
	leftSymbols := strings.Repeat(symbol, halfSymbolLength)
	rightSymbols := strings.Repeat(symbol, symbolLength-halfSymbolLength)
	titleLine := leftSymbols + title + rightSymbols
	symbolLine := strings.Repeat(symbol, length)
	return strings.Join([]string{titleLine, context, symbolLine}, "")
}
