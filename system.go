package main

import (
	"fmt"
	"os"
)

func MainMenu(u *User) {
	for {
		ClearScreen()
		fmt.Printf("\t\t\tATM Management System - Main Menu\n\n")
		fmt.Printf("Welcome %s, what do you want to do ?\n\n", u.Name)
		fmt.Println("[1] - Create a new account		[5] - Make transactions")
		fmt.Println("[2] - Update account information 	[6] - Delete an account")
		fmt.Println("[3] - Check details of an account 	[7] - Transfer ownership of an account")
		fmt.Println("[4] - View the list of your accounts 	[8] - Exit")

		fmt.Print("\nChoose an option: ")
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			CreateAccount(u)
		case 2:
			// UpdateAccount(u)
		case 3:
			// CheckAccountDetails(u)
		case 4:
			// ViewAccountsList(u)
		case 5:
			// MakeTransaction(u)
		case 6:
			// DeleteAccount(u)
		case 7:
			// TransferOwnership(u)
		case 8:
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
			Pause()
		}
	}
}
