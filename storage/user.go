package storage

// User class for user table
type User struct {
	ID    int
	Email string
	Name  string
}

// GetUsers get all users
func GetUsers() []User {
	var u []User
	GopiDB.Find(&u)

	return u
}
