package model

// User role model.
type UserRole struct {
	Id       	int64
	Role 		*Role 	`orm:"rel(fk)"`
	User		*User 	`orm:"rel(fk)"`
}
type UserRoles []UserRole

// Custom table name
func (u *UserRole) TableName() string {
	return "user_role"
}