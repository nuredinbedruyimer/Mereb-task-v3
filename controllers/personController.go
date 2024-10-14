package controllers

import (
	"Mereb3/constants"
	"Mereb3/helpers"
	"Mereb3/models"
	"Mereb3/services"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func CreatePersonController(c echo.Context) error {
	var person models.Person
	if err := c.Bind(&person); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  constants.ERROR_STATUS,
			Message: constants.ERROR_IN_DATA_BINDING,
			Error:   err.Error(),
		})
	}
	personValidator := helpers.NewValidatorService()

	if validationError := personValidator.ValidateData(person); validationError != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  constants.ERROR_STATUS,
			Message: constants.VALIDATION_FAILED,
			Error:   validationError.Error(),
		})
	}
	ctx, cancell := context.WithTimeout(context.Background(), constants.TIME_OUT)

	defer cancell()

	filter := bson.M{"name": person.Name}

	personCount, err := services.PersonCollecion.CountDocuments(ctx, filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  constants.ERROR_STATUS,
			Message: constants.INTERNAL_SERVER_ERROR,
			Error:   err.Error(),
		})

	}

	if personCount > 0 {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  constants.ERROR_STATUS,
			Message: constants.PERSON_ALEADY_EXISTS,
		})

	}

	if err := services.CreatePersonService(&person); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  constants.ERROR_STATUS,
			Message: constants.INTERNAL_SERVER_ERROR,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, Response{
		Status:  constants.SUCCESS_STATUS,
		Message: constants.PERSON_CREATED,
		Data:    person,
	})

}

func GetAllPersonController(c echo.Context) error {

	//  Do simple Pagination

	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")

	perPage := 3
	page := 0

	if currValue, err := strconv.Atoi(limitStr); err == nil && currValue >= 1 {
		perPage = currValue
	}
	if currValue, err := strconv.Atoi(offsetStr); err == nil && currValue >= 1 {
		page = currValue
	}

	persons, err := services.GetAllPersonsService(perPage, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  constants.ERROR_STATUS,
			Message: constants.INTERNAL_SERVER_ERROR,
			Error:   err.Error(),
		})
	}

	//  We can add Page and perPage as Response if we want
	return c.JSON(http.StatusOK, Response{
		Status:  constants.SUCCESS_STATUS,
		Message: constants.PERSONS_FETCHED,
		Data:    persons,
	})

}

func GetPersonController(c echo.Context) error {
	id := c.Param("id")
	person, err := services.GetPersonService(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, Response{
			Status:  constants.ERROR_STATUS,
			Message: constants.PERSON_DOES_NOT_EXIST,
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status:  constants.SUCCESS_STATUS,
		Message: constants.PERSON_FETCHED,
		Data:    person,
	})
}

func UpdatePerson(c echo.Context) error {
	id := c.Param("id")
	var updatedPerson models.Person
	if err := c.Bind(&updatedPerson); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  constants.ERROR_STATUS,
			Message: constants.ERROR_IN_DATA_BINDING,
			Error:   err.Error(),
		})
	}

	newPerson, err := services.UpdatePersonService(id, updatedPerson)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  constants.ERROR_STATUS,
			Message: constants.INTERNAL_SERVER_ERROR,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  constants.SUCCESS_STATUS,
		Message: constants.PERSON_UPDATE,
		Data:    newPerson,
	})
}

// DeletePerson handles DELETE requests to delete a person by ID
func DeletePersonController(c echo.Context) error {
	id := c.Param("id")
	if err := services.DeletePersonService(id); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  constants.ERROR_STATUS,
			Message: constants.INTERNAL_SERVER_ERROR,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  constants.SUCCESS_STATUS,
		Message: constants.PERSON_DELETED,
	})
}
