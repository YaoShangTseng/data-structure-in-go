package composition

type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

func Swim() {
	println("Swimming!")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

//-----------------------------------------------------------

type Trainer interface {
	Train()
}

type Swimmer interface {
	Swim()
}

type SwimmerImpelemtor struct{}

func (s *SwimmerImpelemtor) Swim() {
	println("Swimming!")
}

//匿名的話不用instance，可以直接呼叫
type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

//-----------------------------------------------------------

type Animal struct{}

func (a *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}
