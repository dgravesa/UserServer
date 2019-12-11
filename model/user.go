package model

import (
	"encoding/json"
	"fmt"
)

// User contains profile info for a user account.
type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// UnmarshalJSON gets a User from a JSON encoding.
func (u *User) UnmarshalJSON(data []byte) error {
	nillable := struct {
		ID   *uint64 `json:"id"`
		Name *string `json:"name"`
	}{}

	err := json.Unmarshal(data, &nillable)

	if err != nil {
		return err
	} else if nillable.ID == nil {
		return fmt.Errorf("missing \"id\" field")
	} else if nillable.Name == nil {
		return fmt.Errorf("missing \"name\" field")
	}

	u.ID = *nillable.ID
	u.Name = *nillable.Name
	return nil
}
