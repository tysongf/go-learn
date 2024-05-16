package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var version float64 = 0.1
var password = "g"
var account = Account{
	Name: "Moon Bag",
	Balance: 500.00,
}

func main() {

	var userInput string
	
	fmt.Println("🏦", version)
	fmt.Print("🔒 Password: ")
	fmt.Scan(&userInput)

	if userInput == password {
		for {	showAccountMenu(&account) }
	} else {
		fmt.Println("⛔ Access Denied")
		return
	}
}

func showAccountMenu(account *Account) {
	fmt.Printf("\n%s: $ %.2f\n\n", account.Name, account.Balance)
	fmt.Println("[D] Deposit")
	fmt.Println("[W] Withdraw")
	fmt.Println("[Q] Quit")
	fmt.Println()

	var userInput string = "";

	fmt.Print("Select Option: ")

	fmt.Scan(&userInput)
	userInput = strings.ToUpper(userInput);

	switch userInput {
		case "D":
			showDepositMenu(account)
		case "W":
			showWithdrawMenu(account)
		case "Q":
			fmt.Print("\nGoodbye ❤️\n\n")
			os.Exit(0)
		default:
			fmt.Print("\n⛔ Invalid Input\n")
	}
}

func showDepositMenu(account *Account) {
	var userInput string;
	fmt.Print("Deposit Amount: ")
	fmt.Scan(&userInput)

	inputAmt, err := strconv.ParseFloat(userInput, 64);

	if (err == nil && inputAmt > 0) {
		account.Balance += inputAmt
		fmt.Printf("\n💸 $%.2f Deposited\n", inputAmt)
	} else {
		fmt.Print("\n⛔ Invalid Input\n")
	}
}

func showWithdrawMenu(account *Account) {
	var userInput string;
	fmt.Print("Withdraw Amount: ")
	fmt.Scan(&userInput)

	inputAmt, err := strconv.ParseFloat(userInput, 64);

	if (err == nil && inputAmt > 0) {
		if(inputAmt <= account.Balance) {
			account.Balance -= inputAmt
			fmt.Printf("\n💵 $%.2f Dispensed\n", inputAmt)
		} else {
			fmt.Printf("\n⛔ Insufficient Funds\n")
		}
	}	else {
		fmt.Printf("\n⛔ Invalid Input\n")
	}
}