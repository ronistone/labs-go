package controller

import (
	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
	"github.com/ronistone/labs-go/model"
	"github.com/ronistone/labs-go/service"
	"net/http"
	"strconv"
)

type PersonController struct {
	personService service.PersonService
}

func (p PersonController) Register(echo *echo.Echo) {
	v1 := echo.Group("/v1/person")
	v1.GET("/:id", p.GetUserById)
	v1.POST("", p.CreatePerson)
	v1.PUT("/:id", p.UpdateUser)
	v1.DELETE("/:id", p.DeleteUser)
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
	}
}
