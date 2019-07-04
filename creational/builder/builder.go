package creational

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProdcut
}

// Director
// 有一個builder變數，滿足BuilderProcess interface
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

// 	不管是什麼東西，只要有implement BuilderProcess 的東西就可以當作Builder丟進去
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Product
type VehicleProdcut struct {
	Wheels    int
	Seats     int
	Structure string
}

// A Builder of type car
type CarBuilder struct {
	v VehicleProdcut
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProdcut {
	return c.v
}

// A Builder of type motorbike
type BikeBuilder struct {
	v VehicleProdcut
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProdcut {
	return b.v
}

// A Builder of type bus
type BusBuilder struct {
	v VehicleProdcut
}

func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 8
	return b
}

func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}

func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

func (b *BusBuilder) GetVehicle() VehicleProdcut {
	return b.v
}
