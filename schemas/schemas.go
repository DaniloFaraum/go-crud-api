package schemas

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name      string
	Age       int
	Gender    string
	Ocupation string
	Salary    int
}
