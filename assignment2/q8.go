package main

import (
	"fmt"
	"os"
)

type bankAccount struct {
	accNo      string
	ifscCode   string
	name       string
	typeOfAcc  string
	balance    int
	minBalance int
}

func q8() {
	user := createAccount()
	for {
		fmt.Println("Do you wish to deposit money or withdraw money or check your balance or exit from your bank account(enter your choice)")
		var ch string
		fmt.Scan(&ch)
		switch ch {
		case "deposit":
			{
				var amt int
				fmt.Println("Enter amount to deposit money")
				fmt.Scan(&amt)
				deposit(amt, &user)
			}
		case "withdraw":
			{
				var amt int
				fmt.Println("Enter amount to withdraw money")
				fmt.Scan(&amt)
				withdraw(amt, &user)
			}
		case "check":
			{
				fmt.Println(check(user))
			}
		case "exit":
			{
				os.Exit(0)
			}
		default:
			{
				fmt.Println("Invalid choice")
			}
		}

	}
}
func createAccount() (user bankAccount) {
	user = bankAccount{}
	fmt.Print("Enter account no: ")
	fmt.Scan(&user.accNo)
	fmt.Print("Enter account name: ")
	fmt.Scan(&user.name)
	fmt.Print("Enter account type: ")
	fmt.Scan(&user.typeOfAcc)
	user.minBalance = 1000
	user.balance = 0
	fmt.Println("Minimum balance is 1000, deposited")
	deposit(1000, &user)
	fmt.Println("Thank you for opening an account!")
	return
}
func deposit(amount int, user *bankAccount) {
	user.balance += amount
}
func withdraw(amount int, user *bankAccount) {
	if user.balance-amount < user.minBalance {
		fmt.Println("You cannot withdraw money,since it exceeds minimum balance")
	} else {
		user.balance -= amount
	}

}
func check(account bankAccount) int {
	return account.balance
}
