package controllers

import (
	"D-2/lib/database"
	"D-2/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func V1GetAllUserController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// var users []models.User
	// if err := config.DB.Find(&users).Error; err != nil {
	// }
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"data":    users,
	})
}

func V1GetUserById(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	user, err := database.GetDetailUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"data":    user,
	})
}

func V1CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	newUser, err := database.CreateNewUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"user":    newUser,
	})
}

func V1UpdateUser(c echo.Context) error {
	user := models.User{}
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Bind(&user)
	updatedUser, err := database.UserUpdate(userId, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"user":    updatedUser,
	})
}

func V1DeleteUser(c echo.Context) error {
	user := models.User{}
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	deletedUser, err := database.DeletedUserById(userId, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "user Has been deleted",
		"code":        http.StatusOK,
		"userDeleted": deletedUser,
	})
}

func V1LoginUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "emmail/password invalid"})
	}
	token, err := database.LoginUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "Login success",
		"code":         http.StatusOK,
		"access_token": token,
	})
}
