package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
func ClearScreenAndTitle(title string) {
	ClearScreen()
	fmt.Print("\t\t\t" + title + "\n\n")
}
func Pause() {
	fmt.Println("Press 'Enter' to continue...")
	fmt.Scanln()
}
func PrintUser(user User) {
	fmt.Println("User ID: ", user.Id)
	fmt.Println("Username: ", user.Name)
	fmt.Println("Password: ", user.Password)
}
func PrintRecord(record Record) {
	fmt.Println("Record ID: ", record.Id)
	fmt.Println("Owner ID: ", record.OwnerId)
	fmt.Println("Owner: ", record.Owner.Name)
	fmt.Println("Account Number: ", record.AccountNumber)
	fmt.Println("Creation Date: ", record.CreationDate)
	fmt.Println("Country: ", record.Country)
	fmt.Println("Phone Number: ", record.PhoneNumber)
	fmt.Println("Amount: ", record.Amount)
	fmt.Println("Account Type: ", record.AccountType)
}

func (r Record) PrintDetails() {
	CreationDate, err := time.Parse("2006-01-02", r.CreationDate)
	if err != nil {
		fmt.Println("Error parsing creation date:", err)
		return
	}
	formattedCreationDate := CreationDate.Format("2006/01/02")
	fmt.Printf("Account Number: %d	Amount: $%.2f\n", r.AccountNumber, r.Amount)
	fmt.Printf("Country: %s		Account Type: %s\n", r.Country, r.AccountType)
	fmt.Printf("Phone Number: %s	Creation Date : %s\n", r.PhoneNumber, formattedCreationDate)
	if r.AccountType == "Savings" {
		fmt.Printf("You will get $%.2f as interests every month\n", r.Amount*0.07)
	} else if r.AccountType == "Fixed01" {
		creationDate, err := time.Parse("2006-01-02", r.CreationDate)
		if err != nil {
			fmt.Println("Error parsing creation date:", err)
			return
		}
		oneYearLater := creationDate.AddDate(1, 0, 0)
		fmt.Printf("You will get $%.2f on %s\n", r.Amount*0.04, oneYearLater.Format("2006/01/02"))
	} else if r.AccountType == "Fixed02" {
		creationDate, err := time.Parse("2006-01-02", r.CreationDate)
		if err != nil {
			fmt.Println("Error parsing creation date:", err)
			return
		}
		twoYearsLater := creationDate.AddDate(2, 0, 0)
		fmt.Printf("You will get $%.2f on %s\n", 2*r.Amount*0.05, twoYearsLater.Format("2006/01/02"))
	} else if r.AccountType == "Fixed03" {
		creationDate, err := time.Parse("2006-01-02", r.CreationDate)
		if err != nil {
			fmt.Println("Error parsing creation date:", err)
			return
		}
		threeYearsLater := creationDate.AddDate(3, 0, 0)
		fmt.Printf("You will get $%.2f on %s\n", 3*r.Amount*0.08, threeYearsLater.Format("2006/01/02"))
	}
}
func (r *Record) Deposit(amount float64) {
	r.Amount += amount
}
func (r *Record) Withdraw(amount float64) {
	r.Amount -= amount
}

func verifyPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
