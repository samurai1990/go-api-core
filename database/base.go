package database

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"default:current_timestamp" binding:"-" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" binding:"-" json:"updated_at"`
}
