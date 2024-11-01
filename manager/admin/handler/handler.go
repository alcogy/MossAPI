package handler

import (
	"manager/admin/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func GetIndexHtml(c echo.Context) error {
	return c.File("admin/public/index.html")
}

func GetAllContainer(c echo.Context) error {
	data := models.GetAllContainers()
	return c.JSON(http.StatusOK, data)
}

func GetAllTables(c echo.Context, mysql *sqlx.DB) error {
	data := models.GetAllTables(mysql)
	return c.JSON(http.StatusOK, data)
}