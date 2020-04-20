package main

import "fmt"

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek [7]int
}

type Employee struct {
	Id int
	FirstName string
	LastName string
}

type Weekday int

const(
	Sunday Weekday = iota
	Monday Weekday = iota
	Tuesday Weekday = iota
	Wednesday Weekday = iota
	Thursday Weekday = iota
	Friday Weekday = iota
	Saturday Weekday = iota
)

func (d *Developer) HoursWorked() int{
	total := 0
	for _, v := range d.WorkWeek{
		total += v
	}
	return total
}

func (d *Developer) LogHours(day Weekday, hours int){
	d.WorkWeek[day] = hours
}

func main(){
	d := Developer{
		Individual:Employee{1,
			"Tony", "Stark"},
	HourlyRate:10}
	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)
	fmt.Println("Hours worked on Monday:", d.WorkWeek[Monday])
	fmt.Println("Hours worked on Tuesday:", d.WorkWeek[Tuesday])
	fmt.Println("Hours worked on week:", d.HoursWorked())

}