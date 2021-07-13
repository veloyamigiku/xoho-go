package model

type Theater struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Type []Type `json:"types" gorm:"many2many:theater_types"`
}
