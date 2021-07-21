package db

type Type struct {
	Id   int `gorm:"primaryKey"`
	Name string
}
