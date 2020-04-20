package main

type Developer2 struct {
	Individual Employee2
	HourlyRate int
	Workeek [7]int
}
type Employee2 struct {
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

func main() {
}

func Te(n string) string {
	return n
}

func (d *Developer2) LogHours(day Weekday, hours int){
	d.Workeek[day] = hours
}

func (d *Developer2) HoursWorked() int{
	total := 0
	for _,v := range d.Workeek {
		total += v
	}
	return total
}

func (d *Developer2) PayDay() (int, bool) {
	if d.HoursWorked() > 40 {
		ho := d.HoursWorked() - 40
		ot := ho * 2
		rp := d.HoursWorked() * d.HourlyRate
		return rp + ot, true
	}
	
	return d.HoursWorked() * d.HourlyRate, false
}

func nonLoggedHours() func(int) int{
	total := 0
	
	return func(i int) int {
		total += i
		return total
	}
}

func (d * Developer2) PayDetails(){
}