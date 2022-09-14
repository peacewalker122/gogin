package datareq

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

type Mockrepo struct {
	mock mock.Mock
}

func (m *Mockrepo) Getbyid(id int) (DBdata, error) {
	args := m.mock.Called(id)
	if args.Get(0) == nil {
		return DBdata{}, errors.New("error")
	}
	indec := args.Get(0).(DBdata)
	return indec, nil
}

//	func (m Mockrepo) Getall() ([]DBdata, error) {
//		args := m.mock.Called()
//		if args.Get(0) == nil {
//			return []DBdata{}, errors.New("error")
//		}
//		indec := []string{}
//		for _,r := range args.Get(0).(DBdata){
//			indec = append(indec, r)
//		}
//		return indec, nil
//	}

func (m *Mockrepo) Update(db DBdata) (DBdata, error) {
	args := m.mock.Called(db)
	if args.Get(0) == nil {
		return DBdata{}, errors.New("error")
	}
	indec := args.Get(0).(DBdata)
	return indec, nil
}

func (m *Mockrepo) Delete(db DBdata) (DBdata, error) {
	args := m.mock.Called(db)
	if args.Get(0) == nil {
		return DBdata{}, errors.New("error")
	}
	indec := args.Get(0).(DBdata)
	return indec, nil
}

func (m *Mockrepo) Create(db DBdata) (DBdata, error) {
	args := m.mock.Called(db)
	if args.Get(0) == nil {
		return DBdata{}, errors.New("error")
	}
	indec := args.Get(0).(DBdata)
	return indec, nil
}
