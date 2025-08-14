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

// Initialize akan membuat dan menyambungkan semua komponen, lalu mengembalikan instance Echo
func Initialize(db *sqlx.DB) *echo.Echo {
	// Inisialisasi semua lapisan dari bawah ke atas
	repo := repository.NewRepository(db)
	srv := service.NewService(repo)
	h := delivery.NewHandler(srv)

	// Buat instance Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Definisikan Rute Endpoint
	e.GET("/GetEventList", h.GetEventList)
	e.POST("/GetEventByID", h.GetEventByID)
	e.POST("/InsertEvent", h.InsertEvent)
	e.PUT("/UpdateEvent", h.UpdateEvent)
	e.DELETE("DeleteEvent", h.DeleteEvent)

	return e
}
