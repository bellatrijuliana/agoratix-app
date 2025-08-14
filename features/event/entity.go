package agoratix

import "time"

type Event struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	DateTime         time.Time `json:"date_time"`
	Location         string    `json:"location"`
	Category         string    `json:"category"`
	OrganizerId      string    `json:"organizer_id"`
	OrganizerName    string    `json:"organizer_name"`
	ImageUrl         string    `json:"image_url"`
	TicketCategories string    `json:"ticket_categories"`
}

type RepositoryInterface interface {
	GetEventList() ([]Event, error)
	GetEventByID(id int) (Event, error)
	InsertEvent(input Event) (Event, error)
	UpdateEvent(id int, input Event) (Event, error)
	DeleteEvent(id int) error
}

type ServiceInterface interface {
	GetEventList() ([]Event, error)
	GetEventByID(id int) (Event, error)
	InsertEvent(input Event) (Event, error)
	UpdateEvent(id int, input Event) (Event, error)
	DeleteEvent(id int) error
}
