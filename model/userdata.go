package model

// UserData is an interface to a persistent user store.
type UserData interface {
	Insert(u User)
	FindName(name string) (User, bool)
}

var userData UserData

// SetUserDataLayer sets the backend user data store.
func SetUserDataLayer(ud UserData) {
	userData = ud
}

// FindUserByName returns the user with a given name and true if found, false if not found.
func FindUserByName(name string) (User, bool) {
	return userData.FindName(name)
}
