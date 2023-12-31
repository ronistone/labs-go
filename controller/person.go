package controller

import (
	"context"
	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
	"github.com/ronistone/labs-go/model"
	"github.com/ronistone/labs-go/service"
	"github.com/ronistone/labs-go/trie"
	"net/http"
	"strconv"
)

type PersonController struct {
	personService service.PersonService
	trie          *trie.Trie
}

func (p PersonController) Register(echo *echo.Echo, ctx context.Context) error {
	v1 := echo.Group("/v1/person")
	v1.GET("/:id", p.GetUserById)
	v1.POST("", p.CreatePerson)
	v1.GET("", p.List)
	v1.PUT("/:id", p.UpdateUser)
	v1.DELETE("/:id", p.DeleteUser)
	v1.GET("/search", p.SearchNames)

	persons, err := p.personService.ListPersonsName(ctx)
	if err != nil {
		return err
	}

	for _, person := range persons {
		p.trie.AddWord(person.Name)
	}

	return nil
}

func (p PersonController) SearchNames(echo echo.Context) error {
	name := echo.QueryParam("name")

	names := p.trie.ListWordsByPrefix(name)

	return echo.JSON(http.StatusOK, names)
}

func (p PersonController) List(ctx echo.Context) error {
	persons, err := p.personService.ListPerson(ctx.Request().Context())
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, err)
	}
	ctx.Logger().Infof("Doing List! requestId=%s", ctx.Response().Header().Get(echo.HeaderXRequestID))
	return ctx.JSON(http.StatusOK, persons)
}

func (p PersonController) DeleteUser(echo echo.Context) error {
	id := echo.Param("id")

	idValue, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return handleError(echo, http.StatusBadRequest, err)
	}

	err = p.personService.Delete(echo.Request().Context(), idValue)
	if err != nil {
		return handleError(echo, http.StatusUnprocessableEntity, err)
	}

	return echo.NoContent(http.StatusOK)
}

func handleError(echo echo.Context, statusCode int, err error) error {
	echo.Logger().Error(err)
	return echo.JSON(statusCode, err)
}

func (p PersonController) UpdateUser(echo echo.Context) error {
	id := echo.Param("id")
	idValue, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return handleError(echo, http.StatusBadRequest, err)
	}

	var person model.Person
	if err := echo.Bind(&person); err != nil {
		return handleError(echo, http.StatusBadRequest, err)
	}

	person.ID = &idValue

	person, err = p.personService.Update(echo.Request().Context(), person)

	switch err {
	case nil:
		return echo.JSON(http.StatusOK, person)
	case dbr.ErrNotFound:
		return handleError(echo, http.StatusNotFound, err)
	default:
		return handleError(echo, http.StatusUnprocessableEntity, err)

	}
}

func (p PersonController) CreatePerson(echo echo.Context) error {
	var person model.Person

	if err := echo.Bind(&person); err != nil {
		return handleError(echo, http.StatusBadRequest, err)
	}

	person, err := p.personService.Create(echo.Request().Context(), person)

	if err != nil {
		return handleError(echo, http.StatusUnprocessableEntity, err)
	}
	p.trie.AddWord(person.Name)
	return echo.JSON(http.StatusCreated, person)
}

func (p PersonController) GetUserById(echo echo.Context) error {
	id := echo.Param("id")

	idValue, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return handleError(echo, http.StatusBadRequest, err)
	}

	person, err := p.personService.GetPersonById(idValue)

	switch err {
	case nil:
		return echo.JSON(http.StatusOK, person)
	case dbr.ErrNotFound:
		return handleError(echo, http.StatusNotFound, err)
	default:
		return handleError(echo, http.StatusInternalServerError, err)

	}
}

func CreatePersonController(personService service.PersonService) *PersonController {
	return &PersonController{
		personService: personService,
		trie:          trie.CreateTrie(),
	}
}
