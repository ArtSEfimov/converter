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

	return NewDistance(title, microMeter(floatNumber)), nil

}
