package repository

import (
	"context"
	"github.com/gocraft/dbr/v2"
	"github.com/ronistone/labs-go/model"
)

type PersonRepository interface {
	GetPersonById(id int64) (model.Person, error)
	CreatePerson(ctx context.Context, person model.Person) (model.Person, error)
	DeletePerson(ctx context.Context, id int64) error
	UpdatePerson(ctx context.Context, person model.Person) (model.Person, error)
	List(ctx context.Context) ([]model.Person, error)
	ListNames(ctx context.Context) ([]model.Person, error)
}

type Person struct {
	DbConnection *dbr.Connection
}

func CreatePersonRepository(connection *dbr.Connection) PersonRepository {
	return &Person{
		connection,
	}
}

func (p Person) List(ctx context.Context) ([]model.Person, error) {
	statement := p.DbConnection.NewSession(nil).SelectBySql(`SELECT * FROM person;`)

	results := []model.Person{}

	_, err := statement.LoadContext(ctx, &results)
	if err != nil {
		return []model.Person{}, err
	}
	return results, nil
}

func (p Person) ListNames(ctx context.Context) ([]model.Person, error) {
	statement := p.DbConnection.NewSession(nil).SelectBySql(`SELECT name FROM person;`)

	results := []model.Person{}

	_, err := statement.LoadContext(ctx, &results)
	if err != nil {
		return []model.Person{}, err
	}
	return results, nil
}

func (p Person) GetPersonById(id int64) (model.Person, error) {
	session := p.DbConnection.NewSession(nil)

	statement := session.SelectBySql("SELECT id, name, created_at, updated_at FROM person WHERE id = ?", id)

	var result model.Person
	err := statement.LoadOne(&result)
	if err != nil {
		return model.Person{}, err
	}

	return result, nil
}

func (p Person) DeletePerson(ctx context.Context, id int64) error {
	statement := p.DbConnection.NewSession(nil).SelectBySql(`
		DELETE FROM person WHERE id = ?
	`, &id)

	_, err := statement.RowsContext(ctx)

	return err
}

func (p Person) UpdatePerson(ctx context.Context, person model.Person) (model.Person, error) {
	statement := p.DbConnection.NewSession(nil).SelectBySql(`
		UPDATE person set name = ?, updated_at = NOW() WHERE ID = ?
		returning *
	`, &person.Name, person.ID)

	_, err := statement.LoadContext(ctx, &person)
	if err != nil {
		return model.Person{}, nil
	}

	return person, nil

}

func (p Person) CreatePerson(ctx context.Context, person model.Person) (model.Person, error) {

	statement := p.DbConnection.NewSession(nil).SelectBySql(`
		INSERT INTO person (id, name, created_at, updated_at) VALUES (default, ?, default, default)
		returning *;
	`, &person.Name)

	var result model.Person
	if _, err := statement.LoadContext(ctx, &result); err != nil {
		return model.Person{}, err
	}

	return result, nil

}
