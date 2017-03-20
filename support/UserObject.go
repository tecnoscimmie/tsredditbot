package support

import (
	"encoding/json"
	"io"
)

// User telegram object data struct
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// DecodeJSON decodes some JSON into a User
func (u *User) DecodeJSON(r io.ReadCloser) error {
	d := json.NewDecoder(r)
	err := d.Decode(u)
	if err != nil {
		return err
	}

	return nil
}
