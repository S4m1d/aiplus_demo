package da

type Employee struct {
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	City        string `json:"city"`
}
