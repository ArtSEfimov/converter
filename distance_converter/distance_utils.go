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
		if unicode.IsDigit(char) {
			valueBuilder.WriteRune(char)
			lastDigitIndexPointer = &i
		} else if char == '.' || char == ',' {
			valueBuilder.WriteRune('.')
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

	title = strings.Trim(title, ".")
	value = strings.Trim(value, ".")

	title = strings.ToLower(title)

	if err := validateDistanceTitle(title); err != nil {
		return nil, fmt.Errorf("type validation error: %w", err)
	}

	floatNumber, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("bad convertation to number: %w", err)
	}

	return NewDistance(title, microMeter(floatNumber)), nil

}

func (m microMeter) Equal(anotherValue any) bool {
	switch assertionValue := anotherValue.(type) {
	case microMeter:
		return m == assertionValue
	case *Distance:
		return m == assertionValue.value
	default:
		return false
	}
}

func (distance *Distance) Equal(anotherValue any) bool {
	switch assertionValue := anotherValue.(type) {
	case microMeter:
		return distance.value == assertionValue
	case *Distance:
		return distance.value == assertionValue.value
	default:
		return false
	}
}

func (m microMeter) Convert(anotherValue any) *Distance {
	switch assertionValue := anotherValue.(type) {
	case string:
		assertionValue = strings.ToLower(assertionValue)
		switch assertionValue {
		case "mk":
			return NewDistance("mk", m)
		case "mm":
			return NewDistance("mm", m/Millimeter)
		case "cm":
			return NewDistance("cm", m/Centimeter)
		case "dm":
			return NewDistance("dm", m/Decimeter)
		case "m":
			return NewDistance("m", m/Meter)
		case "km":
			return NewDistance("km", m/Kilometer)
		default:
			return nil
		}
	case microMeter:
		switch assertionValue {
		case Millimeter:
			return NewDistance("mm", m/Millimeter)
		case Centimeter:
			return NewDistance("cm", m/Centimeter)
		case Decimeter:
			return NewDistance("dm", m/Decimeter)
		case Meter:
			return NewDistance("m", m/Meter)
		case Kilometer:
			return NewDistance("km", m/Kilometer)
		default:
			return NewDistance("mk", m)
		}
	case *Distance:
		return NewDistance(assertionValue.title, m)
	}
	return nil
}

func (distance *Distance) Convert(anotherValue any) *Distance {
	switch assertionValue := anotherValue.(type) {
	case string:
		assertionValue = strings.ToLower(assertionValue)
		switch assertionValue {
		case "mk":
			return NewDistance("mk", distance.value)
		case "mm":
			return NewDistance("mm", distance.value/Millimeter)
		case "cm":
			return NewDistance("cm", distance.value/Centimeter)
		case "dm":
			return NewDistance("dm", distance.value/Decimeter)
		case "m":
			return NewDistance("m", distance.value/Meter)
		case "km":
			return NewDistance("km", distance.value/Kilometer)
		default:
			return nil
		}
	case microMeter:
		switch assertionValue {
		case Millimeter:
			return NewDistance("mm", distance.value/Millimeter)
		case Centimeter:
			return NewDistance("cm", distance.value/Centimeter)
		case Decimeter:
			return NewDistance("dm", distance.value/Decimeter)
		case Meter:
			return NewDistance("m", distance.value/Meter)
		case Kilometer:
			return NewDistance("km", distance.value/Kilometer)
		default:
			return NewDistance("mk", distance.value)
		}
	case *Distance:
		return NewDistance(assertionValue.title, distance.value)
	}
	return nil
}
