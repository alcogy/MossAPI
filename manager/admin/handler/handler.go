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

func GetAllContainers(c echo.Context) error {
	data := models.GetAllContainers()
	return c.JSON(http.StatusOK, data)
}

func PostContainer(c echo.Context) error {
	arg := new(CreateArgs)
	if err := c.Bind(arg); err != nil {
		panic(err)
	}
	models.CreateContainer(arg.Service, arg.Port, arg.Artifact)
	return c.JSON(http.StatusOK, Message{Message: "ok"})
}

func GetAllTables(c echo.Context, mysql *sqlx.DB) error {
	data := models.GetAllTables(mysql)
	return c.JSON(http.StatusOK, data)
}