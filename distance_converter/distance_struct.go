package distance_converter

import "fmt"

type micrometer float64

const (
	Micrometer micrometer = 1
	Millimeter micrometer = 1_000
	Centimeter micrometer = 10_000
	Decimeter  micrometer = 100_000
	Meter      micrometer = 1_000_000
	Kilometer  micrometer = 1_000_000_000
)

type Distance struct {
	title string
	value micrometer
}

func NewDistance(title string, value micrometer) *Distance {
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
