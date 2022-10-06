package users

import (
	"os/exec"
)

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

func CreateUser(fName, lName, address, city, state, zipcode, phoneNumber string) User {

	var fakeUser User
	var uuid, err = exec.Command("uuidgen").Output()
	if err != nil {
	} else {
		fakeUser.uid = string(uuid)
	}
	fakeUser.fName = fName
	fakeUser.lName = lName
	fakeUser.address = address
	fakeUser.city = city
	fakeUser.state = state
	fakeUser.zipcode = zipcode
	fakeUser.phoneNumber = phoneNumber
	return fakeUser
}

func Checkout() string {
	return "OK"
}

func DoGetUser(user User) User {

	return user
}
