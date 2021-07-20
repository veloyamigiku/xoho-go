package model

type Prefecture struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Sub  string `json:"sub"`
}
