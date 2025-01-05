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
		case MicroLiter:
			return NewVolume("mkl", m).String()
		case Milliliter:
			return NewVolume("ml", m/Milliliter).String()
		case Liter:
			return NewVolume("l", m/Liter).String()
		case Barrel:
			return NewVolume("bl", m/Barrel).String()
		case OilBarrel:
			return NewVolume("bbl", m/OilBarrel).String()
		case Gallon:
			return NewVolume("km", m/Gallon).String()
		default:
			return "unknown format"
		}
	case string:
		switch formatValue {
		case "mkl":
			return NewVolume("mkl", m).String()
		case "ml":
			return NewVolume("ml", m/Milliliter).String()
		case "l":
			return NewVolume("l", m/Liter).String()
		case "bl":
			return NewVolume("bl", m/Barrel).String()
		case "bbl":
			return NewVolume("bbl", m/OilBarrel).String()
		case "gal":
			return NewVolume("gal", m/Gallon).String()
		default:
			return "unknown format"
		}
	}
	return "unknown type"
}
