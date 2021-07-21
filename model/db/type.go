package db

type Type struct {
	Id         int `gorm:"primaryKey"`
	Name       string
	Title      string
	Sub        string
	Opt        string
	IconPrefix string `gorm:"icon_prefix"`
	IconClass  string `gorm:"icon_class"`
}
