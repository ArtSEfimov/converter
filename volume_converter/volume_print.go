package volume_converter

import "fmt"

func (volume *Volume) String() string {
	switch volume.title {
	case "ml":
		return fmt.Sprintf("%v%s", volume.value/Milliliter, volume.title)
	case "l":
		return fmt.Sprintf("%v%s", volume.value/Liter, volume.title)
	case "bl":
		return fmt.Sprintf("%v%s", volume.value/Barrel, volume.title)
	case "bbl":
		return fmt.Sprintf("%v%s", volume.value/OilBarrel, volume.title)
	case "gal":
		return fmt.Sprintf("%v%s", volume.value/Gallon, volume.title)
	default:
		return fmt.Sprintf("%v%s", volume.value, volume.title)
	}
}

func (m microLiter) ToString(formatValue any) string {
	switch formatValue.(type) {
	case microLiter:
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
