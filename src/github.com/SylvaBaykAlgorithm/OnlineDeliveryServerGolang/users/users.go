package users

import (
	"errors"
)

func CreateUser(user User) (string, error) {
	var newUser UserRequest
	// var uuid, err = exec.Command("uuidgen").Output()
	// if err != nil {
	// } else {
	// 	fakeUser.uid = string(uuid)
	// }
	newUser.FName = user.FName
	newUser.LName = user.LName
	newUser.Address = user.Address
	newUser.City = user.City
	newUser.State = user.State
	newUser.Zipcode = user.Zipcode
	newUser.PhoneNumber = user.PhoneNumber
	if newUser.FName == "" {
		return "", errors.New("No First Name")
	}

	return newUser.FName, nil
}

func Checkout() string {
	return "OK"
}

func GetUser(user User) User {

	return user
}
