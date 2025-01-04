package distance_converter

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Parse(valueTitle string) (*Distance, error) {
	var value string
	var title string

	valueBuilder := strings.Builder{}
	titleBuilder := strings.Builder{}

	var lastDigitIndexPointer *int
	var firstNoDigitIndexPointer *int

	for i, char := range valueTitle {
		if unicode.IsDigit(char) || char == '.' {
			valueBuilder.WriteRune(char)
			lastDigitIndexPointer = &i
		} else {
			titleBuilder.WriteRune(char)
			if firstNoDigitIndexPointer == nil {
				firstNoDigitIndexPointer = &i
			}
		}
	}

	if lastDigitIndexPointer == nil || firstNoDigitIndexPointer == nil {
		return nil, fmt.Errorf("failed to parse value: %s", valueTitle)
	}
	if *lastDigitIndexPointer > *firstNoDigitIndexPointer {
		return nil, fmt.Errorf("failed to parse value: %s", valueTitle)
	}

	title = titleBuilder.String()
	value = valueBuilder.String()

	title = strings.TrimSpace(title)
	value = strings.TrimSpace(value)

	if err := validateDistanceTitle(title); err != nil {
		return nil, fmt.Errorf("type validation error: %w", err)
	}

	floatNumber, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("bad convertation to number: %w", err)
	}

	return NewDistance(title, micrometer(floatNumber)), nil

}

func (value micrometer) Equal(anotherValue any) bool {
	switch assertionValue := anotherValue.(type) {
	case micrometer:
		return value == assertionValue
	case *Distance:
		return value == assertionValue.value
	default:
		return false
	}
}

func (distance *Distance) Equal(anotherValue any) bool {
	switch assertionValue := anotherValue.(type) {
	case micrometer:
		return distance.value == assertionValue
	case *Distance:
		return distance.value == assertionValue.value
	default:
		return false
	}
}

func (value micrometer) Convert(anotherValue any) *Distance {
	switch assertionValue := anotherValue.(type) {
	case string:
		switch assertionValue {
		case "mm":
			return NewDistance("mm", value/Millimeter)
		case "cm":
			return NewDistance("cm", value/Centimeter)
		case "m":
			return NewDistance("m", value/Meter)
		case "km":
			return NewDistance("km", value/Kilometer)
		default:
			return NewDistance("mc", value)
		}
	case micrometer:
		switch assertionValue {
		case Millimeter:
			return NewDistance("mm", value/Millimeter)
		case Centimeter:
			return NewDistance("cm", value/Centimeter)
		case Meter:
			return NewDistance("m", value/Meter)
		case Kilometer:
			return NewDistance("km", value/Kilometer)
		default:
			return NewDistance("mc", value)
		}
	case *Distance:
		return NewDistance(assertionValue.title, value)
	}
	return nil
}

func (distance *Distance) Convert(anotherValue any) *Distance {
	switch assertionValue := anotherValue.(type) {
	case string:
		switch assertionValue {
		case "mm":
			return NewDistance("mm", distance.value/Millimeter)
		case "cm":
			return NewDistance("cm", distance.value/Centimeter)
		case "m":
			return NewDistance("m", distance.value/Meter)
		case "km":
			return NewDistance("km", distance.value/Kilometer)
		default:
			return NewDistance("mc", distance.value)
		}
	case micrometer:
		switch assertionValue {
		case Millimeter:
			return NewDistance("mm", distance.value/Millimeter)
		case Centimeter:
			return NewDistance("cm", distance.value/Centimeter)
		case Meter:
			return NewDistance("m", distance.value/Meter)
		case Kilometer:
			return NewDistance("km", distance.value/Kilometer)
		default:
			return NewDistance("mc", distance.value)
		}
	case *Distance:
		return NewDistance(assertionValue.title, distance.value)
	}
	return nil
}
