package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name      string
	Age       int
	Gender    string
	Ocupation string
	Salary    int
	Alive     bool
}

type PersonResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"` //exlude field of json if empty or negative
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	Ocupation string    `json:"ocupation"`
	Salary    int       `json:"salary"`
	Alive     bool      `json:"alive"`
}
