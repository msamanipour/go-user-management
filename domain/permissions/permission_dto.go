package permissions

import (
	"time"
)

type Permission struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name" gorm:"name" form:"name"`
	GuardName string    `json:"guard_name" gorm:"guard_name" form:"guard_name"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}
