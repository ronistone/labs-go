package repository

import (
	"github.com/gocraft/dbr/v2"
	"github.com/ronistone/labs-go/model"
)

type PersonRepository interface {
	GetPersonById(id int) (model.Person, error)
}

type Person struct {
	DbConnection *dbr.Connection
}

func CreatePersonRepository(connection *dbr.Connection) PersonRepository {
	return &Person{
		connection,
	}
}

func (p Person) GetPersonById(id int) (model.Person, error) {
	session := p.DbConnection.NewSession(nil)

	rows, err := session.SelectBySql("SELECT id, name, created_at, updated_at FROM person WHERE id = ?", id).Rows()
	if err != nil {
		return model.Person{}, err
	}

	rows.Next()
	result := model.Person{}
	err = rows.Scan(&result.ID, &result.Name, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return model.Person{}, err
	}

	return result, nil
}
