package datareq

type Service interface {
	//Getall() ([]DBdata, error)
	Update(data Userreq, id int) (DBdata, error)
	Create(data Userreq) (DBdata, error)
	Delete(id int) (DBdata, error)
	Getbyid(id int) (DBdata, error)
}
type service struct {
	Repository Repository
}

// a "call" to "hub" into repository
func Newservice(repo Repository) *service {
	return &service{repo}
}

// Implement of Get all repository
//func (s *service) Getall() ([]DBdata, error) {
//	test, err := s.Repository.Getall()
//	return test, err
//}

// Create implements Service
func (s *service) Create(data Userreq) (DBdata, error) {
	test := DBdata{
		Name:    data.Name,
		Deposit: data.Deposit,
	}
	assert, err := s.Repository.Create(test)
	return assert, err
}

// Delete implements Service
func (s *service) Delete(id int) (DBdata, error) {
	test, _ := s.Repository.Getbyid(id)
	return s.Repository.Delete(test)
	// return DBdata{},nil
}

// Getbyid implements Service
func (s *service) Getbyid(id int) (DBdata, error) {
	test, err := s.Repository.Getbyid(id)
	return test, err
}

// Update implements Service
func (s *service) Update(data Userreq, id int) (DBdata, error) {
	getid, _ := s.Repository.Getbyid(id)
	getid.Name = data.Name
	getid.Deposit = data.Deposit
	return s.Repository.Update(getid)
}
