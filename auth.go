package main

import (
	"errors"
	"fmt"
)

func LoginMenu() {
	var option int
	isInvalidOption := false
	for {
		ClearScreen()
		fmt.Print("\t\t\tWelcome to the bank, please choose an option\n\n")
		fmt.Println("\t\t\t\t[1] Login")
		fmt.Println("\t\t\t\t[2] Register")
		fmt.Print("\t\t\t\t[3] Exit\n\n")
		if isInvalidOption {
			fmt.Println("Invalid option, please try again :")
			isInvalidOption = false
		} else {
			fmt.Println("Choose an option: ")
		}
		fmt.Scanln(&option)
		switch option {
		case 1:
			Login()
		case 2:
			Register()
		case 3:
			return
		default:
			isInvalidOption = true
		}
	}
}

func Login() {
	isInvalidCredentials := false
	var username string
	var password string
	title := "Welcome to the bank, please choose an option"
	for {
		ClearScreenAndTitle(title)
		if isInvalidCredentials {
			fmt.Println("Invalid credentials, please try again")
			isInvalidCredentials = false
		}
		fmt.Print("Username: ")
		fmt.Scanln(&username)
		fmt.Print("Password: ")
		fmt.Scanln(&password)
		user, err := CheckLogin(username, password)
		if err != nil {
			isInvalidCredentials = true
		}
		if err == nil {
			MainMenu(&user)
		}
	}
}
func CheckLogin(username, pwd string) (User, error) {
	for _, user := range Users {
		if user.Name == username && user.Password == pwd {
			return user, nil
		}
	}
	return User{}, errors.New("invalid credentials")
}
func Register() {
	var username string
	var password string
	isDupe := false
	title := "Welcome to the bank"
	for {
		ClearScreenAndTitle(title)
		if isDupe {
			fmt.Println("Username already exists")
		}
		isDupe = false
		fmt.Print("Username: ")
		fmt.Scanln(&username)
		fmt.Print("Password: ")
		fmt.Scanln(&password)
		user := User{Id: len(Users) + 1, Name: username, Password: password}
		for _, u := range Users {
			if u.Name == username {
				isDupe = true
			}
		}
		if !isDupe {
			SaveUser(user)
			MainMenu(&user)
		}

	}
}
