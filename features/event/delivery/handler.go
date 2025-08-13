package delivery

import (
	"net/http"
	"strconv"

	agoratix "github.com/bellatrijuliana/agoratix-app/features/event"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service agoratix.ServiceInterface
}

// NewHandler membuat instance handler baru
func NewHandler(srv agoratix.ServiceInterface) *handler {
	return &handler{
		service: srv,
	}
}

// GetEventList menangani permintaan GET /events
func (h *handler) GetEventList(c echo.Context) error {
	events, err := h.service.GetEventList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "cannot get events"})
	}
	return c.JSON(http.StatusOK, events)
}

// GetEventByID menangani permintaan GET /events/:id
func (h *handler) GetEventByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("ID"))
	event, err := h.service.GetEventByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "event not found"})
	}
	return c.JSON(http.StatusOK, event)
}

// InsertEvent menangani permintaan POST /events
func (h *handler) InsertEvent(c echo.Context) error {
	var input agoratix.Event
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid input"})
	}

	newEvent, err := h.service.InsertEvent(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "cannot insert event"})
	}
	return c.JSON(http.StatusCreated, newEvent)
}

// UpdateEvent menangani permintaan PUT /events/:id
func (h *handler) UpdateEvent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var input agoratix.Event
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid input"})
	}

	err := h.service.UpdateEvent(id, input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "cannot update event"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "event updated successfully"})
}

// DeleteEvent menangani permintaan DELETE /events/:id
func (h *handler) DeleteEvent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteEvent(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "cannot delete event"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "event deleted successfully"})
}
