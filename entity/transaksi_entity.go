package entity

import (
	"github.com/google/uuid"
)

type Transaksi struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
	Type   string    `json:"type"`
	Amount int       `json:"amount"`
	Notes  string    `json:"notes"`

	User *User `gorm:"foreignKey:UserID"`

	Timestamp
}
