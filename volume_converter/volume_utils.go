package volume_converter

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Parse(valueTitle string) (*Volume, error) {
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
	value = strings.Trim(title, ".")

	if err := validateVolumeTitle(title); err != nil {
		return nil, fmt.Errorf("type validation error: %w", err)
	}

	floatNumber, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("bad convertation to number: %w", err)
	}

	return NewVolume(title, microLiter(floatNumber)), nil

}

func (m microLiter) Equal(anotherValue any) bool {
	switch assertionValue := anotherValue.(type) {
	case microLiter:
		return m == assertionValue
	case *Volume:
		return m == assertionValue.value
	default:
		return false
	}
}

func (volume *Volume) Equal(anotherValue any) bool {
	switch assertionValue := anotherValue.(type) {
	case microLiter:
		return volume.value == assertionValue
	case *Volume:
		return volume.value == assertionValue.value
	default:
		return false
	}
}

func (m microLiter) Convert(anotherValue any) *Volume {
	switch assertionValue := anotherValue.(type) {
	case string:
		switch assertionValue {
		case "ml":
			return NewVolume("ml", m/Milliliter)
		case "l":
			return NewVolume("l", m/Liter)
		case "bl":
			return NewVolume("bl", m/Barrel)
		case "bbl":
			return NewVolume("bbl", m/OilBarrel)
		case "gal":
			return NewVolume("gal", m/Gallon)

		default:
			return NewVolume("mkl", m)
		}
	case microLiter:
		switch assertionValue {
		case Milliliter:
			return NewVolume("ml", m/Milliliter)
		case Liter:
			return NewVolume("l", m/Liter)
		case Barrel:
			return NewVolume("bl", m/Barrel)
		case OilBarrel:
			return NewVolume("bbl", m/OilBarrel)
		case Gallon:
			return NewVolume("gal", m/Gallon)
		default:
			return NewVolume("mkl", m)
		}
	case *Volume:
		return NewVolume(assertionValue.title, m)
	}
	return nil
}

func (volume *Volume) Convert(anotherValue any) *Volume {
	switch assertionValue := anotherValue.(type) {
	case string:
		switch assertionValue {
		case "ml":
			return NewVolume("ml", volume.value/Milliliter)
		case "l":
			return NewVolume("l", volume.value/Liter)
		case "bl":
			return NewVolume("bl", volume.value/Barrel)
		case "bbl":
			return NewVolume("bbl", volume.value/OilBarrel)
		case "gal":
			return NewVolume("gal", volume.value/Gallon)
		default:
			return NewVolume("mkl", volume.value)
		}
	case microLiter:
		switch assertionValue {
		case Milliliter:
			return NewVolume("ml", volume.value/Milliliter)
		case Liter:
			return NewVolume("l", volume.value/Liter)
		case Barrel:
			return NewVolume("bl", volume.value/Barrel)
		case OilBarrel:
			return NewVolume("bbl", volume.value/OilBarrel)
		case Gallon:
			return NewVolume("gal", volume.value/Gallon)
		default:
			return NewVolume("mkl", volume.value)
		}
	case *Volume:
		return NewVolume(assertionValue.title, volume.value)
	}
	return nil
}
