package model

import (
	"time"
)

/*Event struct*/
type Event struct {
	//gorm.Model
	ID        uint64
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	CreatedAt time.Time  `json:"created"`
	UpdatedAt time.Time  `json:"updated"`
	DeletedAt *time.Time `json:"deleted"`
}

/*Validate Event struct*/
func (event *Event) Validate() (string, bool) {

	if event.Name == "" {
		return "name missing", false
	}

	if event.Type == "" {
		return "Event Type missing", false
	}

	return "", true
}
