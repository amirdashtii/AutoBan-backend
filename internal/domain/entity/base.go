package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
}

type BaseModel struct {
	gorm.Model
	ID uint64 `gorm:"primary_key" json:"id"`
}
