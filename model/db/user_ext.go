package db

type UserExt struct {
	Id            int `gorm:"primaryKey"`
	AuthMissCount int `gorm:"column:auth_miss_count"`
}
