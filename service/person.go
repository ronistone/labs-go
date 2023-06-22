package service

import (
	"context"
	"github.com/ronistone/labs-go/model"
	"github.com/ronistone/labs-go/repository"
)

type PersonService interface {
	GetPersonById(id int64) (model.Person, error)
	Create(ctx context.Context, person model.Person) (model.Person, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, person model.Person) (model.Person, error)
}

type Person struct {
	PersonRepository repository.PersonRepository
}

func (p Person) GetPersonById(id int64) (model.Person, error) {
	return p.PersonRepository.GetPersonById(id)
}

func (p Person) Create(ctx context.Context, person model.Person) (model.Person, error) {
	return p.PersonRepository.CreatePerson(ctx, person)
}

func (p Person) Delete(ctx context.Context, id int64) error {
	return p.PersonRepository.DeletePerson(ctx, id)
}

func (p Person) Update(ctx context.Context, person model.Person) (model.Person, error) {
	oldPerson, err := p.GetPersonById(*person.ID)

	if err != nil {
		return model.Person{}, err
	}

	person.CreatedAt = oldPerson.CreatedAt
	person.ID = oldPerson.ID

	return p.PersonRepository.UpdatePerson(ctx, person)
}

func CreatePersonService(personRepository repository.PersonRepository) PersonService {
	return &Person{
		personRepository,
	}
}
