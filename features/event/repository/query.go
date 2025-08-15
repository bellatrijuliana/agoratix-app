package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"

	agoratix "github.com/bellatrijuliana/agoratix-app/features/event"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) agoratix.RepositoryInterface {
	return &repository{
		db: db,
	}
}

func (r *repository) GetEventList() ([]agoratix.Event, error) {
	var events []agoratix.Event
	// CONTOH QUERY: Ganti 'events' dengan nama tabel Anda
	rows, err := r.db.Query("SELECT id, title, description, date_time, location, category, organizer_id, organizer_name, image_url, ticket_categories FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event agoratix.Event
		// Pindai data dari baris ke struct Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.DateTime, &event.Location, &event.Category, &event.OrganizerId, &event.OrganizerName, &event.ImageUrl, &event.TicketCategories)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *repository) GetEventByID(id int) (agoratix.Event, error) {
	var event agoratix.Event
	// Gunakan '?' sebagai placeholder untuk mencegah SQL Injection
	err := r.db.QueryRow("SELECT id, title, description, date_time, location, category, organizer_id, organizer_name, image_url, ticket_categories FROM events WHERE id = ?", id).Scan(&event.ID, &event.Title, &event.Description, &event.DateTime, &event.Location, &event.Category, &event.OrganizerId, &event.OrganizerName, &event.ImageUrl, &event.TicketCategories)
	if err != nil {
		return agoratix.Event{}, err
	}
	return event, nil
}

func (r *repository) InsertEvent(input agoratix.Event) (agoratix.Event, error) {
	// Query untuk memasukkan data baru
	query := "INSERT INTO events (title, description, date_time, location, category, organizer_id, organizer_name, image_url, ticket_categories) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, input.Title, input.Description, input.DateTime, input.Location, input.Category, input.OrganizerId, input.OrganizerName, input.ImageUrl, input.TicketCategories)
	if err != nil {
		return agoratix.Event{}, err
	}

	lastInsertId, _ := result.LastInsertId()
	input.ID = int(lastInsertId)
	return input, nil
}

func (r *repository) UpdateEvent(id int, input agoratix.Event) (agoratix.Event, error) {
	query := "UPDATE events SET title=?, description=?, date_time=?, location=?, category=?, organizer_id=?, organizer_name=?, image_url=?, ticket_categories=? WHERE id=?"
	result, err := r.db.Exec(query, input.Title, input.Description, input.DateTime, input.Location, input.Category, input.OrganizerId, input.OrganizerName, input.ImageUrl, input.TicketCategories, id)
	if err != nil {
		return agoratix.Event{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return agoratix.Event{}, errors.New("no rows were updated")
	}

	// Ambil kembali data yang baru saja diperbarui dan kembalikan
	return r.GetEventByID(id)
}

func (r *repository) DeleteEvent(id int) error {
	result, err := r.db.Exec("DELETE FROM events WHERE id=?", id)
	if err != nil {
		// Ini untuk error koneksi atau sintaks
		return err
	}

	// Cek berapa banyak baris yang terpengaruh
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Jarang terjadi, tapi baik untuk diperiksa
		return err
	}

	// INI KUNCINYA: Jika tidak ada baris yang terhapus, kita anggap itu error
	if rowsAffected == 0 {
		return errors.New("event not found or no rows were deleted")
	}

	return nil // Sukses, karena setidaknya 1 baris terhapus
}
