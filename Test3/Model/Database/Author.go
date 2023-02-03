package Database

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Book       []Book
	BookId     uint
	AuthorName string `gorm:"type:varchar(100)"`
}
