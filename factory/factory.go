// factory/factory.go
package factory

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bellatrijuliana/agoratix-app/features/event/delivery"
	"github.com/bellatrijuliana/agoratix-app/features/event/repository"
	"github.com/bellatrijuliana/agoratix-app/features/event/service"
)

func Initialize(db *sqlx.DB) *echo.Echo {
	repo := repository.NewRepository(db)
	srv := service.NewService(repo)
	h := delivery.NewHandler(srv)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/GetEventList", h.GetEventList)
	e.POST("/GetEventByID", h.GetEventByID)
	e.POST("/InsertEvent", h.InsertEvent)
	e.PUT("/UpdateEvent", h.UpdateEvent)
	e.DELETE("DeleteEvent", h.DeleteEvent)

	return e
}
