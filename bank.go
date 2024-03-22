/*
1. Check balance
2. Deposit money
3. Withdraw mone
*/

package main

import "fmt"

var accountBalance float64 = 1000

func main() {

	//Prints welcome  message
	welcomeMessage()
	var choice int
	fmt.Scan(&choice)

	//wantsCheckBalance := choice == 1

	if choice == 1 {
		fmt.Println("Your account balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Println("Enter amount to deposit")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount")
			return
		}
		accountBalance = accountBalance + depositAmount

		fmt.Println("Your total balance is: ", accountBalance)
	} else if choice == 3 {
		fmt.Println("Enter amount to withdraw")
		var amountToWithdraw float64
		fmt.Scan(&amountToWithdraw)
		if amountToWithdraw > accountBalance {
			fmt.Println("You don't have enough balance to withdraw")
			return
		}
		accountBalance = accountBalance - amountToWithdraw
		fmt.Println("Your withdraw has done")
		fmt.Println("Your total balance is: ", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}

func welcomeMessage() {
	fmt.Println("Welcome to bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Println("Please enter your choice")
}
