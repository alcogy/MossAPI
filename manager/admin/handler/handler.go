package handler

import (
	"manager/admin/models"
	"manager/admin/types"
	"manager/table"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func GetIndexHtml(c echo.Context) error {
	return c.File("admin/public/index.html")
}

func GetInfrastructureInfo(c echo.Context, db *sqlx.DB) error {
	var infra types.InfrastructureInfo
	infra.Gateway = models.IsActiveGateway()
	err := db.Ping()
	if err == nil {
		infra.Database = true
	}	
	return c.JSON(http.StatusOK, infra)
}

func GetAllServices(c echo.Context) error {
	data := models.GetAllServices()
	return c.JSON(http.StatusOK, data)
}

func PostService(c echo.Context) error {
	var arg types.CreateServiceBody
	if err := c.Bind(&arg); err != nil {
		panic(err)
	}
	models.CreateService(arg)
	return c.JSON(http.StatusOK, types.Message{Message: "ok"})
}

func StartService(c echo.Context) error {
	id := c.Param("id")
	models.RunService(id)
	return c.JSON(http.StatusOK, types.Message{ Message: "ok" })
}

func StopService(c echo.Context) error {
	id := c.Param("id")
	models.StopService(id)
	return c.JSON(http.StatusOK, types.Message{ Message: "ok" })
}

func RemoveService(c echo.Context) error {
	service := c.Param("service")
	models.RemoveService(service)
	return c.JSON(http.StatusOK, types.Message{ Message: "ok" })
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

func CrateTable(c echo.Context, db *sqlx.DB) error {
	var arg table.Table
	if err := c.Bind(&arg); err != nil {
		panic(err)
	}
	err := models.CreateTable(db, arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Message{ Message: err.Error() })
	}
	return c.JSON(http.StatusOK, types.Message{ Message: "ok" })
}

func DeleteTableDetail(c echo.Context, mysql *sqlx.DB) error {
	table := c.Param("table")	
	err := models.DeleteTableDetail(mysql, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Message{ Message: err.Error() })
	}

	return c.JSON(http.StatusOK, types.Message{ Message: "ok" })
}