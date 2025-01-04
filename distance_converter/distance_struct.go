package distance_converter

import "fmt"

type microMeter float64

const (
	Micrometer microMeter = 1
	Millimeter microMeter = 1_000
	Centimeter microMeter = 10_000
	Decimeter  microMeter = 100_000
	Meter      microMeter = 1_000_000
	Kilometer  microMeter = 1_000_000_000
)

type Distance struct {
	title string
	value microMeter
}

func NewDistance(title string, value microMeter) *Distance {
	switch title {
	case "mm":
		value *= Millimeter
	case "cm":
		value *= Centimeter
	case "dm":
		value *= Decimeter
	case "m":
		value *= Meter
	case "km":
		value *= Kilometer

	}
	return &Distance{title: title, value: value}
}

func validateDistanceTitle(title string) error {
	if _, ok := abbreviations[title]; !ok {
		return fmt.Errorf("bad type: %v", title)
	}
	return nil
}

var abbreviations = map[string]struct{}{
	"mk": struct{}{},
	"mm": struct{}{},
	"cm": struct{}{},
	"dm": struct{}{},
	"m":  struct{}{},
	"km": struct{}{},
}
