package datareq

import "gorm.io/gorm"

type Repository interface {
	//Getall() ([]DBdata, error)
	Update(data DBdata) (DBdata, error)
	Create(data DBdata) (DBdata, error)
	Delete(data DBdata) (DBdata, error)
	Getbyid(id int) (DBdata, error)
}
type repository struct {
	db *gorm.DB
}

func Newrepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Getall() ([]DBdata, error) {
	var data []DBdata
	err := r.db.Find(data).Error
	return data, err
}
func (r *repository) Getbyid(id int) (DBdata, error) {
	var data DBdata
	err := r.db.First(data, id).Error
	return data, err
}
func (r *repository) Update(data DBdata) (DBdata, error) {
	err := r.db.Save(data).Error
	return data, err
}
func (r *repository) Create(data DBdata) (DBdata, error) {
	err := r.db.Create(data).Error
	return data, err
}
func (r *repository) Delete(data DBdata) (DBdata, error) {
	err := r.db.Delete(data).Error
	return data, err
}
