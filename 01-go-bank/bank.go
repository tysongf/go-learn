package main

import "fmt"

var version float32 = 0.1
var password = "gobank"
var account Account

func main() {

	var userInput string
	
	fmt.Println("Welcome to Go Bank", version)
	fmt.Print("Enter password: ")
	fmt.Scan(&userInput)

	if userInput == password {
		loadAccount()
		showAccountMenu()
	} else {
		fmt.Println("Access Denied")
		return
	}
}

func loadAccount() {
	account = Account{
		"Checking Account",
		1,
		500.00,
	}
}

func showAccountMenu() {
	fmt.Print("---------- [ MAIN MENU ]---------\n\n")
	fmt.Print(account.Name, "\n\n")
	fmt.Print("[D]eposit\n")
	fmt.Print("[W]ithdraw\n\n")
	fmt.Print("---------------------------------\n")
	fmt.Print(" => ")
}