package volume_converter

type microLiter float64

const (
	MicroLiter microLiter = 1
	Milliliter microLiter = 1_000
	Liter      microLiter = 1_000_000
	Barrel     microLiter = Liter / 158.987295
	OilBarrel  microLiter = Liter / 119.240471
	Gallon     microLiter = Liter / 3.785411784
)
