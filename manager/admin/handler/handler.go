package handler

import (
	"manager/admin/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Message struct {
	Message string `json:"message"`
}

type CreateArgs struct {
	Service string `json:"service"`
	Port string `json:"port"`
	Artifact string `json:"artifact"`
}

func GetIndexHtml(c echo.Context) error {
	return c.File("admin/public/index.html")
}

func GetAllServices(c echo.Context) error {
	data := models.GetAllServices()
	return c.JSON(http.StatusOK, data)
}

func PostService(c echo.Context) error {
	arg := new(CreateArgs)
	if err := c.Bind(arg); err != nil {
		panic(err)
	}
	models.CreateService(arg.Service, arg.Port, arg.Artifact)
	return c.JSON(http.StatusOK, Message{Message: "ok"})
}

func StartService(c echo.Context) error {
	id := c.Param("id")
	models.RunService(id)
	return c.JSON(http.StatusOK, Message{ Message: "ok" })
}

func StopService(c echo.Context) error {
	id := c.Param("id")
	models.StopService(id)
	return c.JSON(http.StatusOK, Message{ Message: "ok" })
}

func RemoveService(c echo.Context) error {
	service := c.Param("service")
	models.RemoveService(service)
	return c.JSON(http.StatusOK, Message{ Message: "ok" })
}

func GetAllTables(c echo.Context, mysql *sqlx.DB) error {
	data := models.GetAllTables(mysql)
	return c.JSON(http.StatusOK, data)
}

func GetTableDetail(c echo.Context, mysql *sqlx.DB) error {
	table := c.Param("table")
	data := models.GetTableDetail(mysql, table)
	return c.JSON(http.StatusOK, data)
}

func DeleteTableDetail(c echo.Context, mysql *sqlx.DB) error {
	table := c.Param("table")	
	err := models.DeleteTableDetail(mysql, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Message{ Message: err.Error() })
	}

	return c.JSON(http.StatusOK, Message{ Message: "ok" })
}