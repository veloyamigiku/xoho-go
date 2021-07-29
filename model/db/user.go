package db

type User struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Password  string
	UserExtId int
	UserExt   UserExt `gorm:"ForeignKey:user_ext_id;AssociationForeignKey:id"`
}
