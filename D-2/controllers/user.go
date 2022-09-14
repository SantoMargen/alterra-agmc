package controllers

import (
	"D-2/config"
	"D-2/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"data":    users,
	})
}

func GetUserById(c echo.Context) error {
	var user models.User
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Find(&user, userId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"data":    user,
	})
}

func Createuser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func UpdateUser(c echo.Context) error {
	user := models.User{}
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusBadRequest,
			"message": "Cannot parse payload",
		})
	}
	userIput := models.User{}
	json.Unmarshal(body, &userIput)

	config.DB.First(&user, userId)
	if err := config.DB.Model(&user).Updates(userIput).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"user":    user,
	})
}

func DeleteUser(c echo.Context) error {
	user := models.User{}
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	config.DB.First(&user, userId)

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	/*
		Hard Delete
			if err := config.DB.Unscoped().Delete(&user, "id = ? ", userId).Error; err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
	*/

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user Has been deleted",
		"code":    http.StatusOK,
		"data":    nil,
	})
}
