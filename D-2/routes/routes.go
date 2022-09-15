package routes

import (
	"D-2/controllers"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CostumValidator struct {
	validator *validator.Validate
}

func (cv *CostumValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New() *echo.Echo {
	route := echo.New()
	route.Validator = &CostumValidator{validator: validator.New()}

	v1 := route.Group("/v1")
	v1Books := v1.Group("/books")
	v1Users := v1.Group("/users")

	jwtSecret := os.Getenv("JWT_SECRET")
	v1Users.Use(middleware.JWT([]byte(jwtSecret)))

	v1Books.POST("", controllers.CreateBook)
	v1Books.GET("", controllers.GetBooks)
	v1Books.GET("/:id", controllers.GetBooksById)
	v1Books.PUT("/:id", controllers.EditPriceBook)
	v1Books.DELETE("/:id", controllers.DeleteBook)
	v1Users.GET("", controllers.V1GetAllUserController)
	v1Users.GET("/:id", controllers.V1GetUserById)
	v1Users.PUT("/:id", controllers.V1UpdateUser)
	v1Users.DELETE("/:id", controllers.V1DeleteUser)
	route.POST("v1/login", controllers.V1LoginUser)
	route.POST("v1/users", controllers.V1CreateUser)

	return route
}
