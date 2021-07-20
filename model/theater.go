package model

type Theater struct {
	Id           int        `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name"`
	Sub          string     `json:"sub"`
	Url          string     `json:"url"`
	AreaId       int        `json:"-"`
	Area         Area       `json:"area" gorm:"ForeignKey:area_id;AssociationForeignKey:id"`
	PrefectureId int        `json:"-"`
	Prefecture   Prefecture `json:"prefecture" gorm:"ForeignKey:prefecture_id;AssociationForeignKey:id"`
	Type         []Type     `json:"types" gorm:"many2many:theater_types"`
}
