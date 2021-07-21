package db

type Prefecture struct {
	Id   int `gorm:"primaryKey"`
	Name string
	Sub  string
}
