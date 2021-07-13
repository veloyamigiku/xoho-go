package model

type Type struct {
	Id   int    `gorm:"primaryKey"`
	Name string `json:"name"`
}
