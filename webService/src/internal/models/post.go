package models

type Post struct {
	Id               string `gorm:"type:uuid;primaryKey"`
	ShortDescription string `gorm:"type:varchar(1000)"`
	Description      string `gorm:"type:varchar(10000)"`
}
