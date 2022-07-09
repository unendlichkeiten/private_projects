package conf

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string    `json:"name,omitempty" gorm:"column:name"`
	UUID     uuid.UUID `json:"uuid,omitempty" gorm:"column:uuid"`
	Role     string    `json:"role,omitempty" gorm:"column:role"`
	Age      int       `json:"age,omitempty" gorm:"column:age"`
	Birthday time.Time `json:"birthday,omitempty" gorm:"column:birthday"`
}
