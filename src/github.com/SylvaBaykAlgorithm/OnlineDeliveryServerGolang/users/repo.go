package users

type UserStore interface {
	Checkout() string
	CreateUser() (string, error)
	GetAllUsers() ([]User, error)
	GetUser(id string) (User, error)
	UpdateUser(user UserRequest) (string, error)
	DeleteUser(user UserRequest) (string, error)
}
