package delivery

import (
	"net/http"
	"strconv"

	agoratix "github.com/bellatrijuliana/agoratix-app/features/event"
	"github.com/bellatrijuliana/agoratix-app/utils/responses"
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
		// PERBAIKAN: Gunakan FailedResponse
		return c.JSON(http.StatusInternalServerError, responses.FailedResponse("cannot get events", err.Error()))
	}
	// PERBAIKAN: Gunakan SuccessWithDataResponse
	return c.JSON(http.StatusOK, responses.SuccessWithDataResponse(events, "successfully retrieved all events"))
}

// DeleteEvent menangani permintaan DELETE /events/:id
func (h *handler) DeleteEvent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteEvent(id)
	if err != nil {
		// PERBAIKAN: Gunakan FailedResponse
		return c.JSON(http.StatusInternalServerError, responses.FailedResponse("cannot delete event", err.Error()))
	}
	// PERBAIKAN: Gunakan SuccessResponse
	return c.JSON(http.StatusOK, responses.SuccessResponse("event deleted successfully"))
}

// ... (Ubah fungsi handler lainnya dengan pola yang sama) ...

// GetEventByID
func (h *handler) GetEventByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	event, err := h.service.GetEventByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, responses.FailedResponse("event not found", err.Error()))
	}
	return c.JSON(http.StatusOK, responses.SuccessWithDataResponse(event, "successfully retrieved event"))
}

// InsertEvent
func (h *handler) InsertEvent(c echo.Context) error {
	var input agoratix.Event
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("invalid input", err.Error()))
	}

	newEvent, err := h.service.InsertEvent(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.FailedResponse("cannot insert event", err.Error()))
	}
	return c.JSON(http.StatusCreated, responses.SuccessWithDataResponse(newEvent, "event created successfully"))
}

// UpdateEvent
func (h *handler) UpdateEvent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var input agoratix.Event
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("invalid input", err.Error()))
	}

	updatedEvent, err := h.service.UpdateEvent(id, input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.FailedResponse("cannot update event", err.Error()))
	}
	return c.JSON(http.StatusOK, responses.SuccessWithDataResponse(updatedEvent, "event updated successfully"))
}
