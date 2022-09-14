package routes

import (
	"D-2/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()

	v1 := route.Group("/v1")
	v1Books := v1.Group("/books")
	v1Users := v1.Group("/users")
	v1Books.POST("", controllers.CreateBook)
	v1Books.GET("", controllers.GetBooks)
	v1Books.GET("/:id", controllers.GetBooksById)
	v1Books.PUT("/:id", controllers.EditPriceBook)
	v1Books.DELETE("/:id", controllers.DeleteBook)
	v1Users.GET("", controllers.GetAllUser)
	v1Users.GET("/:id", controllers.GetUserById)
	v1Users.PUT("/:id", controllers.UpdateUser)
	v1Users.POST("", controllers.Createuser)
	v1Users.DELETE("/:id", controllers.DeleteUser)

	return route
}
