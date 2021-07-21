package db

type Theater struct {
	Id           int `gorm:"primaryKey"`
	Name         string
	Sub          string
	Url          string
	AreaId       int
	Area         Area `gorm:"ForeignKey:area_id;AssociationForeignKey:id"`
	PrefectureId int
	Prefecture   Prefecture `gorm:"ForeignKey:prefecture_id;AssociationForeignKey:id"`
	Type         []Type     `gorm:"many2many:theater_types"`
}
