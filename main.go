package main

import (
	"fmt"
)

func calculateInterest(paymentAmount, initialBalance, interestRate float64) (totalInterest float64) {
	totalInterest = 0
	balance := initialBalance

	for balance > 0 {
		// Calculate interest for the current month
		interest := balance * interestRate
		totalInterest += interest

		// Determine the amount paid towards reducing the balance
		amountPaid := paymentAmount
		if amountPaid > balance {
			amountPaid = balance
		}
		balance -= amountPaid

		fmt.Printf("Month: Interest=%.2f, Amount Paid=%.2f, New Balance=%.2f\n", interest, amountPaid, balance)
	}

	return totalInterest
}

func main() {
	initialBalance := 100000.00 // Initial balance
	paymentAmount := 10000.00   // Monthly payment
	interestRate := 0.0273      // Monthly interest rate

	totalInterest := calculateInterest(paymentAmount, initialBalance, interestRate)
	fmt.Printf("Total interest paid: %.2f\n", totalInterest)
}
