package repository

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"xoho-go/database"
	"xoho-go/model/db"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestTheater(t *testing.T) {

	curDir, _ := os.Getwd()
	envDir := fmt.Sprintf("%s/../..", curDir)
	database.Connect(envDir)
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	t.Run("GetAllTheater_Exists_With_Type", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {

			newArea := db.Area{
				Name: "Area1",
				Sub:  "area1",
			}
			newAreaCreateErr := tx.
				Debug().
				Create(&newArea).Error
			assert.Nil(
				t,
				newAreaCreateErr)

			firstArea := db.Area{}
			firstAreaErr := tx.
				Debug().
				First(&firstArea).Error
			assert.Nil(
				t,
				firstAreaErr)

			newPref := db.Prefecture{
				Name: "Pref1",
				Sub:  "pref1",
			}
			newPrefCreateErr := tx.
				Debug().
				Create(&newPref).Error
			assert.Nil(
				t,
				newPrefCreateErr)

			firstPref := db.Prefecture{}
			firstPrefErr := tx.
				Debug().
				First(&firstPref).Error
			assert.Nil(
				t,
				firstPrefErr)

			newType := db.Type{
				Name:       "Type1",
				Title:      "TypeTitle1",
				Sub:        "type1",
				Opt:        "typeOpt1",
				IconPrefix: "typeIconPrefix1",
				IconClass:  "typeIconClass1"}
			newTypeCreateErr := tx.
				Debug().
				Create(&newType).
				Error
			assert.Nil(
				t,
				newTypeCreateErr)

			firstType := db.Type{}
			firstTypeErr := tx.
				Debug().
				First(&firstType).
				Error
			assert.Nil(
				t,
				firstTypeErr)

			newTheater := db.Theater{
				Name:         "Theater1",
				Sub:          "theater1",
				Url:          "http://localhost/",
				AreaId:       firstArea.Id,
				PrefectureId: firstPref.Id}
			newTheaterCreateErr := tx.
				Debug().
				Create(&newTheater).Error
			assert.Nil(
				t,
				newTheaterCreateErr)

			firstTheater := db.Theater{}
			firstTheaterErr := tx.
				Debug().
				First(&firstTheater).
				Error
			assert.Nil(
				t,
				firstTheaterErr)

			newTheaterType := db.TheaterType{
				TheaterId: firstTheater.Id,
				TypeId:    firstType.Id}
			newTheaterTypeCreateErr := tx.
				Debug().
				Create(&newTheaterType).
				Error
			assert.Nil(
				t,
				newTheaterTypeCreateErr)

			firstTheaterType := db.TheaterType{}
			firstTheaterTypeErr := tx.
				Debug().
				First(&firstTheaterType).
				Error
			assert.Nil(
				t,
				firstTheaterTypeErr)

			theaters, getAllTheaterErr := GetAllTheater(tx)
			assert.Nil(
				t,
				getAllTheaterErr)
			assert.Equal(
				t,
				len(theaters),
				1)
			theater := theaters[0]
			assert.Equal(
				t,
				firstArea.Name,
				theater.Area.Name)
			assert.Equal(
				t,
				firstArea.Sub,
				theater.Area.Sub)
			assert.Equal(
				t,
				firstPref.Name,
				theater.Prefecture.Name)
			assert.Equal(
				t,
				firstPref.Sub,
				theater.Prefecture.Sub)
			assert.Equal(
				t,
				firstType.Name,
				theater.Type[0].Name)
			assert.Equal(
				t,
				firstType.IconClass,
				theater.Type[0].IconClass)
			assert.Equal(
				t,
				firstType.IconPrefix,
				theater.Type[0].IconPrefix)
			assert.Equal(
				t,
				firstType.Opt,
				theater.Type[0].Opt)
			assert.Equal(
				t,
				firstType.Sub,
				theater.Type[0].Sub)
			assert.Equal(
				t,
				firstType.Title,
				theater.Type[0].Title)
			assert.Equal(
				t,
				firstTheater.Name,
				theater.Name)
			assert.Equal(
				t,
				firstTheater.Sub,
				theater.Sub)
			assert.Equal(
				t,
				firstTheater.Url,
				theater.Url)

			return errors.New("rollback GetAllTheater_Exists_With_Type")
		})

	})

	t.Run("GetAllTheater_Exists_No_Type", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {
			return errors.New("rollback GetAllTheater_Exists_No_Type")
		})

	})

	t.Run("GetAllTheater_Not_Exists_No_Area", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {
			return errors.New("rollback GetAllTheater_Not_Exists_No_Area")
		})

	})

	t.Run("GetAllTheater_Not_Exists_No_Prefecture", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {
			return errors.New("rollback GetAllTheater_Not_Exists_No_Prefecture")
		})

	})

	t.Run("GetTheaterWithTypeId_Exists_With_Type_And_TypeID", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {
			return errors.New("rollback GetTheaterWithTypeId_Exists_With_Type_And_TypeID")
		})

	})

	t.Run("GetTheaterWithTypeId_Exists_With_Type_And_NoTypeID", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {
			return errors.New("rollback GetTheaterWithTypeId_Exists_With_Type_And_NoTypeID")
		})

	})

	t.Run("GetTheaterWithTypeId_Exists_No_Type", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {
			return errors.New("rollback GetTheaterWithTypeId_Exists_No_Type")
		})

	})

	t.Run("GetTheaterWithTypeId_Exists_No_Area", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {
			return errors.New("rollback GetTheaterWithTypeId_Exists_No_Area")
		})

	})

	t.Run("GetTheaterWithTypeId_Exists_No_Prefecture", func(t *testing.T) {

		database.DB.Transaction(func(tx *gorm.DB) error {
			return errors.New("rollback GetTheaterWithTypeId_Exists_No_Prefecture")
		})

	})

}
