package main

import "fmt"

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
