package main

var Users []User
var Records []Record

func main() {
	//Get Users and Records
	FetchUsers()
	FetchRecords()

	LoginMenu()
}
