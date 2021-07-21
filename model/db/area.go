package db

type Area struct {
	Id   int `gorm:"primaryKey"`
	Name string
	Sub  string
}
