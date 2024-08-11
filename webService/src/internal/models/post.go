package models

type Post struct {
	Id              string `gorm:"primarykey"`
	ShortDecription string `gorm:"type:varchar(1000)"`
	Decription      string `gorm:"type:varchar(10000)"`
}
