package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type (

	Base struct {
		ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid.New()" json:"id"`
		CreatedAt time.Time	 `gorm:"type:timestamp" json:"CreatedAt"`
		UpdatedAt time.Time	 `gorm:"type:timestamp" json:"UpdatedAt"`
		DeletedAt *time.Time `sql:"index"`
	   }

	User struct {
		Base
				// UUID uuid.UUID  `gorm:"primarykey;autoIncrement:false" json:"uuid"`
		Name string 	`json:"name"`
	}
)

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.New()
	return nil
   }
