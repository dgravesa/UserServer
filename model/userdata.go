package model

// UserData is an interface to a persistent user store.
type UserData interface {
	Insert(u User)
	Remove(uid uint64)
	FindName(name string) (User, bool)
	FindID(id uint64) (User, bool)
}

var userData UserData

// SetUserDataLayer sets the backend user data store.
func SetUserDataLayer(ud UserData) {
	userData = ud
}

// AddUser adds a new user to the data.
func AddUser(u User) {
	userData.Insert(u)
}

// DeleteUser removes a user from the data by ID.
func DeleteUser(uid uint64) {
	userData.Remove(uid)
}

// FindUserByName returns the user with a given name and true if found, false if not found.
func FindUserByName(name string) (User, bool) {
	return userData.FindName(name)
}

// FindUserByID returns the user with a given ID and true if found, false if not found.
func FindUserByID(id uint64) (User, bool) {
	return userData.FindID(id)
}
