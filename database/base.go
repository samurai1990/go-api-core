package database

import "time"

type BaseModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"default:current_timestamp" binding:"-" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" binding:"-" json:"updated_at"`
}
