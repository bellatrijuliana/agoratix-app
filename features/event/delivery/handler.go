package delivery

import (
	"fmt"
	"net/http"

	agoratix "github.com/bellatrijuliana/agoratix-app/features/event"
	"github.com/bellatrijuliana/agoratix-app/utils/responses"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service agoratix.ServiceInterface
}

type IdRequest struct {
	ID int `json:"id"`
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

// Fungsi DELETE yang mengambil ID dari body
func (h *handler) DeleteEvent(c echo.Context) error {
	var input IdRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("invalid input data", err.Error()))
	}

	err := h.service.DeleteEvent(input.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.FailedResponse("cannot delete event", err.Error()))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponse("event deleted successfully"))
}

// ... (Ubah fungsi handler lainnya dengan pola yang sama) ...

func (h *handler) GetEventByID(c echo.Context) error {
	var input IdRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("invalid input data", err.Error()))
	}

	// TAMBAHKAN BARIS INI UNTUK DEBUGGING
	fmt.Println("Mencari event dengan ID:", input.ID)

	eventData, err := h.service.GetEventByID(input.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, responses.FailedResponse("event not found", err.Error()))
	}

	return c.JSON(http.StatusOK, responses.SuccessWithDataResponse(eventData, "successfully retrieved event"))
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

// Fungsi UPDATE yang mengambil ID dari body
func (h *handler) UpdateEvent(c echo.Context) error {
	var input agoratix.Event
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("invalid input data", err.Error()))
	}

	id := input.ID
	if id == 0 {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("id is required", ""))
	}

	updatedEvent, err := h.service.UpdateEvent(id, input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.FailedResponse("cannot update event", err.Error()))
	}
	return c.JSON(http.StatusOK, responses.SuccessWithDataResponse(updatedEvent, "event updated successfully"))
}
