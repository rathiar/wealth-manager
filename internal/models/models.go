package models

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
	Role      int
}
