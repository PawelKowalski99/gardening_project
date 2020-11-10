package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	// Base struct {
	// 	UUID      uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid.New()" json:"id"`
	// 	CreatedAt time.Time  `gorm:"type:timestamp;autoCreateTime" json:"CreatedAt"`
	// 	UpdatedAt time.Time  `gorm:"type:timestamp;autoUpdateTime:milli" json:"UpdatedAt"`
	// 	DeletedAt *time.Time `sql:"index"`
	// }

	User struct {
		gorm.Model
		FirstName    string  `gorm:"type:varchar" json:"first_name" form:"first_name"`
		LastName     string  `gorm:"type:varchar" json:"last_name" form:"last_name"`
		Email        string  `gorm:"type:varchar" json:"email" form:"email"`
		Password     string  `gorm:"type:varchar" json:"password" form:"password"`
		Wallet       float64 `gorm:"float(4)"     json:"wallet,string" form:"wallet"`
		Order        Order
		Subscription Subscription
		Role         string
		Range        float64
	}

	//Subscription Model
	Subscription struct {
		gorm.Model
		UserID      uint      `gorm:"unique"`
		Price       float64   `gorm:"type:float(4)"  json:"wallet"`
		Description string    `gorm:"type:varchar"   json:"description"`
		TimeStarted time.Time `gorm:"type:timestamp;autoCreateTime" json:"time_started"`
		TimeEnd     time.Time `gorm:"type:timestamp" json:"time_end"`
	}

	Order struct {
		gorm.Model
		UserID      uint   `gorm:"unique"`
		Description string `gorm:"type:varchar"   json:"description" form:"description"`
		Difficulty  uint   `gorm:"type:varchar"   json:"difficulty"  form:"difficulty"`
	}
)

// func (base *Base) BeforeCreate(tx *gorm.DB) error {
// 	base.UUID = uuid.New()
// 	return nil
// }
