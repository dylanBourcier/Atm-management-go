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
}

func UpdateAccount(u *User) {
	title := "ATM Management - Update account information"
	ClearScreenAndTitle(title)
	record := GetAccountByAccountNumber(u, title)
	// Get the field to update
	ClearScreenAndTitle(title)
	for {
		fmt.Println("What would you like to update?")
		fmt.Println("[1] - Phone Number")
		fmt.Println("[2] - Country")
		fmt.Println("\n[0] - Cancel")
		fmt.Print("\nChoose an option: ")
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			ClearScreenAndTitle(title)
			fmt.Print("New Phone Number: ")
			fmt.Scanln(&record.PhoneNumber)
		case 2:
			ClearScreenAndTitle(title)
			fmt.Print("New Country: ")
			fmt.Scanln(&record.Country)
		case 0:
			MainMenu(u)
		default:
			ClearScreenAndTitle(title)
			fmt.Println("Invalid option")
		}
		if option == 1 || option == 2 {
			break
		}
	}

	// Save the updated record
	SaveRecord(record)

	fmt.Println("Account updated successfully")

}
func CheckAccountDetails(u *User) {
	title := "ATM Management - Check details of an account"
	ClearScreenAndTitle(title)
	fmt.Printf("\t\t\t\n\n")
	record := GetAccountByAccountNumber(u, title)
	ClearScreenAndTitle(title)
	record.PrintDetails()
	Pause()
}
func ViewAccountsList(u *User) {
	title := "ATM Management - View the list of your accounts"
	ClearScreenAndTitle(title)
	fmt.Printf("\t\t\t\n\n")
	fmt.Printf("Welcome %s, here is a list of your accounts\n\n", u.Name)
	for _, rec := range Records {
		if rec.Owner.Id == u.Id {
			rec.PrintDetails()
			fmt.Println()
		}
	}
	Pause()
}

func MakeTransaction(u *User) {
	title := "ATM Management - Make transactions"
	ClearScreenAndTitle(title)
	fmt.Printf("\t\t\t\n\n")
	record := GetAccountByAccountNumber(u, title)
	ClearScreenAndTitle(title)
	for {
		fmt.Println("What would you like to do?")
		fmt.Println("[1] - Deposit")
		fmt.Println("[2] - Withdraw")
		fmt.Println("\n[0] - Cancel")
		fmt.Print("\nChoose an option: ")
		var option string
		fmt.Scanln(&option)
		switch option {
		case "1":
			ClearScreenAndTitle(title)
			record.PrintDetails()
			fmt.Println()
			fmt.Print("Amount to deposit: ")
			var amount float64
			fmt.Scanln(&amount)
			record.Deposit(amount)
			SaveRecord(record)
			fmt.Println("Deposit successful")
		case "2":
			ClearScreenAndTitle(title)
			for {
				record.PrintDetails()
				fmt.Println()
				fmt.Print("Amount to withdraw: ")
				var amount float64
				fmt.Scanln(&amount)
				if amount > record.Amount {
					ClearScreenAndTitle(title)
					fmt.Println("Insufficient funds. Please enter a smaller amount.")
					continue
				}
				record.Withdraw(amount)
				SaveRecord(record)
				fmt.Println("Withdrawal successful")
				break
			}
		case "0":
			MainMenu(u)
			return
		default:
			ClearScreenAndTitle(title)
			fmt.Println("Invalid option")
		}
		if option == "1" || option == "2" {
			break
		}
	}
}
func DeleteAccount(u *User) {
	title := "ATM Management - Delete an account"
	isInvalidOpt := false
	ClearScreenAndTitle(title)
	record := GetAccountByAccountNumber(u, title)
	ClearScreenAndTitle(title)
	for {
		fmt.Println("Are you sure you want to delete this account?")
		record.PrintDetails()
		fmt.Println("\n[1] - Yes")
		fmt.Println("[2] - No")
		if isInvalidOpt {
			fmt.Println("\nInvalid option")
			isInvalidOpt = false
		}
		fmt.Print("Choose an option: ")
		var option string
		fmt.Scanln(&option)
		if option == "1" {
			DeleteRecord(record)
			fmt.Println("Account deleted successfully")
			break
		} else if option == "2" {
			break
		} else {
			ClearScreenAndTitle(title)
			isInvalidOpt = true
		}
	}
}

func TransferOwnership(u *User) {
	title := "ATM Management - Delete an account"
	ClearScreenAndTitle(title)
	record := GetAccountByAccountNumber(u, title)
	for {
		fmt.Println("Please enter the username of the new owner (Press 0 to cancel): ")
		var newOwnerName string
		fmt.Scanln(&newOwnerName)
		if newOwnerName == "0" {
			MainMenu(u)
		}
		newOwner := FetchUserByName(newOwnerName)
		Pause()
		if newOwner.Id == 0 {
			ClearScreenAndTitle(title)
			fmt.Println("User not found. Please enter a valid username.")
		} else {
			record.Owner = newOwner
			record.OwnerId = newOwner.Id
			SaveRecord(record)
			fmt.Println("Account ownership transferred successfully")
			break
		}
	}
}
