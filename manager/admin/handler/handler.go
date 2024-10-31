package handler

import (
	"manager/admin/container"
	"manager/admin/table"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetIndexHtml(c echo.Context) error {
	return c.File("admin/public/index.html")
}

func GetAllContainer(c echo.Context) error {
	data := container.GetAllContainers()
	return c.JSON(http.StatusOK, data)
}

func GetAllTables(c echo.Context) error {
	data := table.GetAllTables()
	return c.JSON(http.StatusOK, data)
}