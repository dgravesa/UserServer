package model

// UserData is an interface to a persistent user store.
type UserData interface {
	Insert(u User)
}

var userData UserData

// SetUserDataLayer sets the backend user data store.
func SetUserDataLayer(ud UserData) {
	userData = ud
}
