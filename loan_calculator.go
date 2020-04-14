package main

import (
	"errors"
	"fmt"
)

const (
	goodScore = 450
	lowScoreRatio = 10
	goodScoreRatio = 20
)

var (
	errCreditScore = errors.New("invalid credit score")
	errIncome = errors.New("income invalid")
	errLoanAmount = errors.New("loan amount invalid")
	errLoanTerm = errors.New("loan term not a multiple of 12")
)

func main() {

	fmt.Println("Applicant 1")
	fmt.Println("-----------")
	err := checkLoan(500, 1000, 1000, 24)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Applicant 2")
	fmt.Println("-----------")
	err2 := checkLoan(350, 1000, 1000, 12)

	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
}

func checkLoan(creditScore int,
	income float64,
	loanAmount float64,
	loanTerm float64) error {

	interest := 20.0

	if creditScore >= goodScore {
		interest = 15.0
	}

	if creditScore < 1 {
		return errCreditScore
	}

	if income < 1 {
		return errIncome
	}

	if loanAmount < 1 {
		return errLoanAmount
	}

	if loanTerm < 1 || int(loanTerm) % 12 != 0 {
		return errLoanTerm
	}

	rate := interest /100

	payment := ((loanAmount * rate) / loanTerm) + (loanAmount / loanTerm)

	totalInterest := (payment * loanTerm) - loanAmount

	approved := false

	if income > payment {
		ratio := (payment / income) * 100
		if creditScore >= goodScore && ratio < goodScoreRatio {
			approved = true
		} else if ratio < lowScoreRatio {
			approved = true
		}
	}

	fmt.Println("Credit Score :", creditScore)
	fmt.Println("Income :", income)
	fmt.Println("Loan Amount :", loanAmount)
	fmt.Println("Loan Term :", loanTerm)
	fmt.Println("Monthly Payment :", payment)
	fmt.Println("Rate :", interest)
	fmt.Println("Total Cost :", totalInterest)
	fmt.Println("Approved :", approved)
	fmt.Println("")

	return nil

}
