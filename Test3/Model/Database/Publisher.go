package Database

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	PublisherName string `gorm:"type:varchar(100)"`
}
