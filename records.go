package main

import (
	"fmt"
	"strconv"
	"time"
)

func CreateAccount(u *User) {
	title := "ATM Management - Create a new account"
	ClearScreenAndTitle(title)
	fmt.Printf("\t\t\t\n\n")
	fmt.Printf("Welcome %s, please fill out the following form\n\n", u.Name)

	var record Record
	record.Owner = *u

	// Get the account number
	for {
		fmt.Print("Account Number: ")
		var input string
		fmt.Scanln(&input)
		accountNumber, err := strconv.Atoi(input)
		if err != nil || accountNumber == 0 {
			ClearScreenAndTitle(title)
			fmt.Println("Invalid input. Account Number must be a non-zero integer.")
			continue
		}

		// Check if the account number already exists for the user
		exists := false
		for _, rec := range Records {
			if rec.Owner.Name == u.Name && rec.AccountNumber == accountNumber {
				exists = true
				break
			}
		}

		if exists {
			ClearScreenAndTitle(title)
			fmt.Println("Account Number already exists for this user. Please choose a different number.")
		} else {
			record.AccountNumber = accountNumber
			break
		}
	}

	// Get the country
	ClearScreenAndTitle(title)
	fmt.Print("Country: ")
	fmt.Scanln(&record.Country)

	// Get the phone number
	ClearScreenAndTitle(title)
	fmt.Print("Phone Number: ")
	fmt.Scanln(&record.PhoneNumber)

	// Get the amount
	ClearScreenAndTitle(title)
	fmt.Print("Amount: ")
	fmt.Scanln(&record.Amount)

	// Get the account type
	ClearScreenAndTitle(title)
	for {
		fmt.Println("Account Type:")
		fmt.Println("[1] - Savings")
		fmt.Println("[2] - Current")
		fmt.Println("[3] - Fixed01")
		fmt.Println("[4] - Fixed02")
		fmt.Println("[5] - Fixed03")
		fmt.Print("\nChoose an option: ")
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			record.AccountType = "Savings"
		case 2:
			record.AccountType = "Current"
		case 3:
			record.AccountType = "Fixed01"
		case 4:
			record.AccountType = "Fixed02"
		case 5:
			record.AccountType = "Fixed03"
		default:
			ClearScreenAndTitle(title)
			fmt.Println("Invalid option")
		}
		if option >= 1 && option <= 5 {
			break
		}
	}

	//Generate the record ID
	record.Id = len(Records) + 1

	//Set the creation date
	record.CreationDate = time.Now().Format("2006-01-02")

	//Add the record to the list of records
	SaveRecord(record)

	fmt.Println("Account created successfully")
	Pause()
}
