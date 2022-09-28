package users

type User struct {
	fName       string
	lName       string
	address     string
	city        string
	state       string
	zipcode     string
	uid         string
	phoneNumber string
}

func Checkout() string {
	return "OK"
}

func DoGetUser(user User) User {

	return user
}
