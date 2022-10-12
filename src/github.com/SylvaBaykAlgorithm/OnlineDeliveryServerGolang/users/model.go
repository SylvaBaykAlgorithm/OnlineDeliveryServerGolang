package users

type User struct {
	UID         string
	FName       string
	LName       string
	Address     string
	City        string
	State       string
	Zipcode     string
	PhoneNumber string
	CreatedAt   string
	UpdatedAt   string
}

type UserRequest struct {
	UID         string
	FName       string
	LName       string
	Address     string
	City        string
	State       string
	Zipcode     string
	PhoneNumber string
}
