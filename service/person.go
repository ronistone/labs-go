package service

import (
	"github.com/ronistone/labs-go/model"
	"github.com/ronistone/labs-go/repository"
)

type PersonService interface {
	GetPersonById(id int) (model.Person, error)
}

type Person struct {
	PersonRepository repository.PersonRepository
}

func (p Person) GetPersonById(id int) (model.Person, error) {
	return p.PersonRepository.GetPersonById(id)
}

func CreatePersonService(personRepository repository.PersonRepository) PersonService {
	return &Person{
		personRepository,
	}
}
