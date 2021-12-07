package users

import (
	"go-apk-users/utils/errors"
	"strings"
	"time"
)

type User struct {
	Id        int64     `json:"id" form:"id"`
	Name      string    `json:"name" gorm:"name" form:"name"`
	Family    string    `json:"family" gorm:"family" form:"family"`
	Username  string    `json:"username" gorm:"username" form:"username"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	Password  string    `json:"password" gorm:"password" form:"password"`
	Role      string    `json:"role" form:"role"`
}

type Users []User

func (u *User) Validate(isEdit bool) *errors.RestErr {
	u.Name = strings.TrimSpace(u.Name)
	u.Family = strings.TrimSpace(u.Family)
	u.Username = strings.TrimSpace(strings.ToLower(u.Username))
	if u.Username == "" {
		return errors.NewBadRequestError("username can not be empty")
	}
	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" && !isEdit {
		return errors.NewBadRequestError("password can not be empty")
	}
	return nil
}
