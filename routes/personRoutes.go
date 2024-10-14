package routes

import (
	"Mereb3/controllers"

	"github.com/labstack/echo/v4"
)

func PersonRoutes(personRoutes *echo.Echo) {
	personRoutes.POST("/persons", controllers.CreatePersonController)
	personRoutes.GET("/persons", controllers.GetAllPersonController)
	personRoutes.GET("/persons/:id", controllers.GetPersonController)
	personRoutes.PUT("/persons/:id", controllers.UpdatePerson)
	personRoutes.DELETE("/persons/:id", controllers.DeletePersonController)

}
