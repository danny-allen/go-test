package model

// User model.
type User struct {
	Id 			int
	Name 		string
	Email 		string
	Password	string
	UserRoles 	[]*UserRole		`orm:"reverse(many)"`
}
type Users []User

