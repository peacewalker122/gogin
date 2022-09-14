package datareq

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var Repositorycategory = &Mockrepo{mock: mock.Mock{}}
var Servicecategory = service{Repository: Repositorycategory}

func TestGetbyid(t *testing.T) {
	Repositorycategory.mock.On("Getbyid", 1).Return(nil)
	category, err := Servicecategory.Getbyid(1)
	assert.Equal(t, 0, category.ID)
	assert.NotNil(t, err)
	//assert.Equal(t, category.Name, category.Name)
}
func TestDel(t *testing.T) {
	data := DBdata{
		ID:      1,
		Name:    "sasda",
		Deposit: Deposit{2000, "idr"},
	}

	Repositorycategory.mock.On("Getbyid", 1).Return(data)
	Repositorycategory.mock.On("Delete", data).Return(data)
	_, err := Servicecategory.Getbyid(1)
	r, err := Servicecategory.Delete(1)
	assert.Nil(t, err)
	assert.Equal(t, data.ID, r.ID)
}
func TestCreate(t *testing.T) {
	update := Userreq{
		Name:    "Noasdd",
		Deposit: Deposit{9345,"idr"},
	}
	Repositorycategory.mock.On("Create", update).Return(DBdata{})
	r,err := Servicecategory.Create(update)
	assert.NotNil(t,err)
	assert.Equal(t, update.Name, r.Name)
}
