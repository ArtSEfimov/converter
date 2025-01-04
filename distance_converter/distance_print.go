package distance_converter

import "fmt"

func (distance *Distance) String() string {
	switch distance.title {
	case "mm":
		return fmt.Sprintf("%v%s", distance.value/Millimeter, distance.title)
	case "cm":
		return fmt.Sprintf("%v%s", distance.value/Centimeter, distance.title)
	case "dm":
		return fmt.Sprintf("%v%s", distance.value/Decimeter, distance.title)
	case "m":
		return fmt.Sprintf("%v%s", distance.value/Meter, distance.title)
	case "km":
		return fmt.Sprintf("%v%s", distance.value/Kilometer, distance.title)
	default:
		return fmt.Sprintf("%v%s", distance.value, distance.title)
	}
}

func (m micrometer) ToString(formatValue any) string {
	switch formatValue.(type) {
	case micrometer:
		switch formatValue {
		case Micrometer:
			return NewDistance("mk", m).String()
		case Millimeter:
			return NewDistance("mm", m/Millimeter).String()
		case Centimeter:
			return NewDistance("cm", m/Centimeter).String()
		case Decimeter:
			return NewDistance("dm", m/Decimeter).String()
		case Meter:
			return NewDistance("m", m/Meter).String()
		case Kilometer:
			return NewDistance("km", m/Kilometer).String()
		default:
			return "unknown format"
		}
	case string:
		switch formatValue {
		case "mk":
			return NewDistance("mk", m).String()
		case "mm":
			return NewDistance("mm", m/Millimeter).String()
		case "cm":
			return NewDistance("cm", m/Centimeter).String()
		case "dm":
			return NewDistance("dm", m/Decimeter).String()
		case "m":
			return NewDistance("m", m/Meter).String()
		case "km":
			return NewDistance("km", m/Kilometer).String()
		default:
			return "unknown format"
		}
	}
	return "unknown type"
}
