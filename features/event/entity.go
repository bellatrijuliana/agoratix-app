package agoratix

import "time"

type Event struct {
	ID               int
	Title            string
	Description      string
	DateTime         time.Time
	Location         string
	Category         string
	OrganizerId      string
	OrganizerName    string
	ImageUrl         string
	TicketCategories string
}

type RepositoryInterface interface {
	GetEventList() ([]Event, error)
	GetEventByID(id int) (Event, error)
	InsertEvent(input Event) (Event, error)
	UpdateEvent(id int, input Event) error
	DeleteEvent(id int) error
}

type ServiceInterface interface {
	GetEventList() ([]Event, error)
	GetEventByID(id int) (Event, error)
	InsertEvent(input Event) (Event, error)
	UpdateEvent(id int, input Event) error
	DeleteEvent(id int) error
}
