package volume_converter

import "fmt"

type microLiter float64

const (
	MicroLiter microLiter = 1
	Milliliter microLiter = 1_000
	Liter      microLiter = 1_000_000
	Barrel     microLiter = Liter / 158.987295
	OilBarrel  microLiter = Liter / 119.240471
	Gallon     microLiter = Liter / 3.785411784
)

type Volume struct {
	title string
	value microLiter
}

func NewVolume(title string, value microLiter) *Volume {
	switch title {
	case "ml":
		value *= Milliliter
	case "l":
		value *= Liter
	case "bl":
		value *= Barrel
	case "bbl":
		value *= OilBarrel
	case "gal":
		value *= Gallon

	}
	return &Volume{title: title, value: value}

}

func validateVolumeTitle(title string) error {
	if _, ok := abbreviations[title]; !ok {
		return fmt.Errorf("bad type: %v", title)
	}
	return nil
}

var abbreviations = map[string]struct{}{
	"mkl": struct{}{},
	"ml":  struct{}{},
	"l":   struct{}{},
	"bl":  struct{}{},
	"bbl": struct{}{},
	"gal": struct{}{},
}
