package data

import "github.com/dgravesa/WaterLogger-UserServer/model"

// InMemoryUserStore is an in-memory container for user data.
type InMemoryUserStore struct {
	users  []model.User
	nextID uint64
}

// NewInMemoryUserStore creates a new in-memory container for user data.
func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		nextID: 0,
	}
}

// Insert adds a new user to the data.
func (s *InMemoryUserStore) Insert(u model.User) {
	u.ID = s.nextID
	s.nextID++
	s.users = append(s.users, u)
}

// Remove removes users from the data by ID.
func (s *InMemoryUserStore) Remove(uid uint64) {
	for i, user := range s.users {
		if user.ID == uid {
			s.users = append(s.users[:i], s.users[i+1:]...)
		}
	}
}

// FindName returns the user with a given name and true if found, false if not found.
func (s *InMemoryUserStore) FindName(name string) (model.User, bool) {
	for _, user := range s.users {
		if user.Name == name {
			return user, true
		}
	}
	return model.User{}, false
}

// FindID returns the user with a given ID and true if found, false if not found.
func (s *InMemoryUserStore) FindID(id uint64) (model.User, bool) {
	for _, user := range s.users {
		if user.ID == id {
			return user, true
		}
	}
	return model.User{}, false
}
