package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/ronistone/labs-go/service"
	"strconv"
)

type PersonController struct {
	personService service.PersonService
}

func (p PersonController) Register(echo *echo.Echo) {
	echo.GET("/:id", p.GetUserById)
}

func (p PersonController) GetUserById(echo echo.Context) error {
	id := echo.Param("id")

	idValue, err := strconv.Atoi(id)

	if err != nil {
		echo.Logger().Error(err)
		return err
	}

	person, err := p.personService.GetPersonById(idValue)
	if err != nil {
		echo.Logger().Error(err)
		return err
	}

	return echo.JSON(200, person)
}

func CreatePersonController(personService service.PersonService) *PersonController {
	return &PersonController{
		personService: personService,
	}
}
