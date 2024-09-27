package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func FetchUsers() {
	fileContent, err := os.ReadFile("./data/users.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(fileContent, &Users)
	if err != nil {
		fmt.Println(err)
	}
}
func FetchRecords() {
	fileContent, err := os.ReadFile("./data/records.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(fileContent, &Records)
	if err != nil {
		fmt.Println(err)
	}
	for i, record := range Records {
		user, err := FetchUserById(record.OwnerId)
		if err != nil {
			fmt.Println(err)
		}
		Records[i].Owner = user
	}
}

func FetchUserById(id int) (User, error) {
	for _, user := range Users {
		if id == user.Id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func SaveUser(user User) {
	Users = append(Users, user)
	fileContent, err := json.MarshalIndent(Users, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("./data/users.json", fileContent, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
func SaveRecord(record Record) {
	alreadyExists := false
	for _, rec := range Records {
		if rec.Id == record.Id {
			alreadyExists = true
			break
		}
	}
	if !alreadyExists {
		Records = append(Records, record)
	} else {
		for i, rec := range Records {
			if rec.Id == record.Id {
				Records[i] = record
				break
			}
		}
	}
	fileContent, err := json.MarshalIndent(Records, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("./data/records.json", fileContent, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
func UpdateRecords() {
	fileContent, err := json.MarshalIndent(Records, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("./data/records.json", fileContent, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func FetchRecordByAccountNumber(accountNumber int, userId int) (Record, error) {
	for _, record := range Records {
		if record.OwnerId == userId && accountNumber == record.AccountNumber {
			return record, nil
		}
	}
	return Record{}, errors.New("record not found")
}

func GetAccountByAccountNumber(u *User, title string) Record {
	fmt.Printf("Welcome %s, please fill out the following form\n\n", u.Name)

	var record Record
	availableAccounts := []Record{}
	for _, rec := range Records {
		if rec.Owner.Id == u.Id {
			availableAccounts = append(availableAccounts, rec)
		}
	}
	// Get the account number
	for record.Id == 0 {
		availableAccountsString := "Available Accounts: "
		for _, rec := range availableAccounts {
			availableAccountsString += strconv.Itoa(rec.AccountNumber) + ", "
		}
		fmt.Println(availableAccountsString[:len(availableAccountsString)-2])
		fmt.Println("Account Number (Press 0 to cancel): ")
		var input string
		fmt.Scanln(&input)
		accountNumber, err := strconv.Atoi(input)
		if err != nil {
			ClearScreenAndTitle(title)
			fmt.Println("Invalid input. Account Number must be a non-zero integer.")
			continue
		} else {
			if accountNumber == 0 {
				MainMenu(u)
			}
			record, err = FetchRecordByAccountNumber(accountNumber, u.Id)
			if err != nil {
				ClearScreenAndTitle(title)
				fmt.Println("Account not found. Please enter a valid account number.")
				record.Id = 0
			}
		}
	}
	return record
}
func DeleteRecord(record Record) {
	for i, rec := range Records {
		if rec.Id == record.Id {
			Records = append(Records[:i], Records[i+1:]...)
			break
		}
	}
	UpdateRecords()
}

func FetchUserByName(str string) User {
	for _, user := range Users {
		if user.Name == str {
			return user
		}
	}
	return User{}
}
