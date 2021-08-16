package convert

import (
	"testing"
	"xoho-go/model/db"

	"github.com/stretchr/testify/assert"
)

func TestConvertTheaterFromDbToJson(t *testing.T) {
	dbType1 := db.Type{
		Id:         1,
		Name:       "name1",
		Title:      "title1",
		Sub:        "sub1",
		Opt:        "opt1",
		IconPrefix: "icon_prefix1",
		IconClass:  "icon_class1",
	}
	dbType2 := db.Type{
		Id:         2,
		Name:       "name2",
		Title:      "title2",
		Sub:        "sub2",
		Opt:        "opt2",
		IconPrefix: "icon_prefix2",
		IconClass:  "icon_class2",
	}
	dbTheater := db.Theater{
		Name: "name1",
		Sub:  "sub1",
		Url:  "url1",
		Type: []db.Type{
			dbType1,
			dbType2,
		},
	}

	jsonTheater := ConvertTheaterFromDbToJson(dbTheater)

	assert.Equal(
		t,
		jsonTheater.Name,
		dbTheater.Name)
	assert.Equal(
		t,
		jsonTheater.Sub,
		dbTheater.Sub)
	assert.Equal(
		t,
		jsonTheater.Url,
		dbTheater.Url)

	assert.Equal(
		t,
		len(jsonTheater.Type),
		len(dbTheater.Type),
	)

	for idx, typ := range jsonTheater.Type {
		assert.Equal(
			t,
			typ.Name,
			dbTheater.Type[idx].Name,
		)
		assert.Equal(
			t,
			typ.Sub,
			dbTheater.Type[idx].Sub,
		)
		assert.Equal(
			t,
			typ.Title,
			dbTheater.Type[idx].Title,
		)
		assert.Equal(
			t,
			typ.Option,
			dbTheater.Type[idx].Opt,
		)
		assert.Equal(
			t,
			typ.Icon[0],
			dbTheater.Type[idx].IconPrefix,
		)
		assert.Equal(
			t,
			typ.Icon[1],
			dbTheater.Type[idx].IconClass,
		)
	}

}
