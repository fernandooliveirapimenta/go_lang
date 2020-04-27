package main

import "fmt"

type Employee struct {
	Id int
	FirstName string
	LastName string
}

type Developer struct {
	Individual Employee
	HourlyRate float64
	HoursWorkedInYear float64
	Review map[string]interface{}
}
func (d Developer) FullName() string{
	fullName := d.Individual.FirstName +
		" " + d.Individual.LastName
	return fullName
}
func (d Developer) Pay() (string, float64) {
	ful := d.FullName()
	return ful, d.HourlyRate * d.HoursWorkedInYear
}

type Manager struct {
	Individual Employee
	Salary float64
	CommisionRate float64
}
func (m Manager) Pay() (string, float64){
	f := m.Individual.FirstName + " " + m.Individual.LastName
	return f, m.Salary + ( m.Salary * m.CommisionRate)
}

type Payer interface {
	Pay() (string, float64)
}

type Sayer interface {SayHello() interface{ String() string }}
type Hello struct {}
func (h *Hello) SayHello() interface{ String() string } {return h}
func (h *Hello) String() string {return "Hello"}

func main() {
	var h interface{} = &Hello{}
	h, _ = h.(Sayer)
	fmt.Printf("%v\n",h)
	employeeReview := map[string]interface{}{
		"WorkQuality": 5,
		"TeamWork":2,
		"Communication": "Poor",
		"Problem-solving":4,
		"Dependability": "Unsatisfactory",
	}
	d := Developer{
		Individual: Employee{
			Id:1,
			FirstName: "Eric",
			LastName: "Davis",
		},
		HourlyRate: 35,
		HoursWorkedInYear:2400,
		Review: employeeReview,
	}

	m := Manager{
		Individual: Employee{
			Id: 2,
			FirstName: "Mr.",
			LastName:"Boss",
		},
		Salary: 150000,
		CommisionRate: .07,
	}

	payDetails(d)
	payDetails(m)

}

func payDetails(p Payer){
	f, y := p.Pay()
	fmt.Printf("%s got paid %.2f for the year\n", f, y)
}

