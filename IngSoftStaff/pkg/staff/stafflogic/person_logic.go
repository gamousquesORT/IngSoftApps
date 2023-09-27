package staff

import "IngSoftStaff/pkg/staff"

type PersonService struct {
	repository staff.PersonRepository
}

func NewPersonService(repo staff.PersonRepository) *PersonService {
	return &PersonService{
		repository: repo,
	}
}

func (s *PersonService) ReadPersons() (*[]staff.PersonData, error) {

	staffList, err := s.repository.ImportData()

	if err != nil {
		return nil, err
	}

	return staffList, nil
	
}