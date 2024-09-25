package main

type User struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

type Record struct {
	Id            int     `json:"Id"`
	OwnerId       int     `json:"OwnerId"`
	Owner         User    `json:"-"`
	AccountNumber int     `json:"AccountNumber"`
	CreationDate  string  `json:"CreationDate"`
	Country       string  `json:"Country"`
	PhoneNumber   string  `json:"PhoneNumber"`
	Amount        float64 `json:"Amount"`
	AccountType   string  `json:"AccountType"`
}
