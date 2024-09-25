package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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
	Records = append(Records, record)
	fileContent, err := json.MarshalIndent(Records, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("./data/records.json", fileContent, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
