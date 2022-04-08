package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Participant struct {
	ID                 uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	FullName           string    `gorm:"size:255;not null"`
	BusinessName       string    `gorm:"size:255;not null"`
	Email              string    `gorm:"size:255;not null"`
	PhoneNumber        string    `gorm:"size:255;not null"`
	IsPrintCertificate bool      `gorm:"not_null"`
	IsPrintNameTag     bool      `gorm:"not_null"`
	CreatedAt          time.Time `gorm:"default:NOW()"`
	UpdatedAt          time.Time `gorm:"default:NOW()"`
}

func (u *Participant) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
