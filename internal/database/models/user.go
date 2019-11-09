package models

import (
	"encoding/json"
	"github.com/netsaj/transporte-backend/internal/utils"
)

type User struct {
	Base
	Name     string `gorm:"size:255" json:"name"`
	Email    string `gorm:"unique_index;size:100;not null" json:"email"`
	Password string `gorm:"not null;size:255" json:"-"`
	Username string `gorm:"not null;unique_index;size:100" json:"username"`
	Role     string `gorm:"not null;DEFAULT:'standard';size:100" json:"role"`
	Active   bool   `gorm:"not null;default:'true'" json:"active"`
}

//CheckPassword : Verify user password hash with another string.
func (u *User) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.Password)
}

func (u User) Json() string {
	b, err := json.Marshal(u)
	if err != nil {
		print(err)
	}
	return string(b)
}

func (u User) IsAdmin() bool {
	return u.Role == "Administrator";
}

func (u *User) SetPassword(s string) (err error) {
	u.Password, err = utils.HashPassword(s)
	return
}
