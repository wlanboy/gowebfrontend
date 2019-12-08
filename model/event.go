package model

import (
	"time"
)

/*Event struct*/
type Event struct {
	//gorm.Model
	ID        uint64
	UUID      string     `json:"uuid,omitempty"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	CreatedAt time.Time  `json:"created,omitempty"`
	UpdatedAt time.Time  `json:"updated,omitempty"`
	DeletedAt *time.Time `json:"deleted,omitempty"`
}
