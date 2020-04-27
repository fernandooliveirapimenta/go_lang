package main

import (
	"errors"
	"fmt"
)

var(
	ErrInvalidLastName = errors.New("invalid last name")
	ErrInvaliRoutingNumber = errors.New("invalid routing number")
)

type directDeposit struct {
	firstName string
	lastName string
	bankName string
	routingNumber int
	accountNumber int
}

func (d *directDeposit) validateRoutingNumber() {
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	if(d.routingNumber < 100) {
		panic(ErrInvaliRoutingNumber)
	}
}

func (d *directDeposit) validateLastName() {
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	if(d.lastName == "") {
		panic(ErrInvalidLastName)
	}
}

func main() {

	d := directDeposit{
		firstName: "Abe",
		bankName: "XYZ Inc",
		routingNumber: 17,
		accountNumber: 1809,

	}

    d.validateRoutingNumber()
	d.validateLastName()

	fmt.Println("********************************************")

	fmt.Printf("Last Name: %v\nFirst Name: %v\n" +
		"Bank: %v\nRouting: %v\nAccount: %v\n",
		d.lastName, d.firstName, d.bankName, d.routingNumber, d.accountNumber)

	//fmt.Println(ErrInvalidLastName)
	//fmt.Println(ErrInvaliRoutingNumber)

}
