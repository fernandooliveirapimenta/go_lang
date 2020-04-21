package car

type Car struct {
	Name string
	Color string
	Ff string
}

func (c Car) Start() string{
	return c.Name + " has been started..."
}