package model

import "gorm.io/gorm"

type Note struct {
	Model   gorm.Model `gorm:"embedded"` // <- When this field is active, the date and ID objects are activated.
	Id      int        `gorm:"type:int;primary_key"`
	Content string     `gorm:"not null" json:"content,omitempty"`
}
