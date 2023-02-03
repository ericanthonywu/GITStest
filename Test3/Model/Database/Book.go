package Database

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Publisher   Publisher
	PublisherId uint
	BookName    string `gorm:"type:varchar(100)"`
}
