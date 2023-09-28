package applogic

import "IngSoftStaff/pkg/staff/domain"

type PersonService struct {
	repository domain.PersonRepository
}

func NewPersonService(repo domain.PersonRepository) *PersonService {
	return &PersonService{
		repository: repo,
	}
}

func (s *PersonService) ReadPersons() (*[]domain.PersonData, error) {

	staffList, err := s.repository.ImportData()

	if err != nil {
		return nil, err
	}

	return staffList, nil
	
}