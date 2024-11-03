package admin

import (
	"manager/admin/handler"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/browser"
)

func Serve(mysql *sqlx.DB) {
	e := echo.New()
	
	e.Use(middleware.Logger())
  e.Use(middleware.Recover())

	// Static files.
	e.Static("/", "admin/public")
	e.GET("/", handler.GetIndexHtml)

	// API.
	e.GET("/api/services", handler.GetAllServices)
	e.POST("/api/service/create", handler.PostService)
	e.POST("/api/service/start/:id", handler.StartService)
	e.POST("/api/service/stop/:id", handler.StopService)
	e.POST("/api/service/remove/:service", handler.RemoveService)
	
	e.GET("/api/tables", func (c echo.Context) error {
		return handler.GetAllTables(c, mysql)
	})

	e.GET("/api/table/:table", func (c echo.Context) error {
		return handler.GetTableDetail(c, mysql)
	})
	
	go func() {
		browser.OpenURL("http://localhost:5500")
	}()

	e.Logger.Fatal(e.Start(":5500"))
	
}