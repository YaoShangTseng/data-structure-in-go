package abstract_factory

type LuxuaryCar struct{}

func (l *LuxuaryCar) GetDoors() int {
	return 4
}

func (l *LuxuaryCar) GetWheels() int {
	return 4
}

func (l *LuxuaryCar) GetSeats() int {
	return 5
}
