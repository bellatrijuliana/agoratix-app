package repository

import (
	"testing"
	"time"

	agoratix "github.com/bellatrijuliana/agoratix-app/features/event"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sqlx.DB {
	dsn := "root:215544@tcp(127.0.0.1:3306)/agoratix_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("Failed to connect to test database")
	}

	// Hapus tabel lama jika ada
	db.MustExec("DROP TABLE IF EXISTS events")

	// Buat tabel baru dengan skema yang LENGKAP
	schema := `
    CREATE TABLE events (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255),
        description TEXT,
        date_time DATETIME,
        location VARCHAR(255),
        category VARCHAR(100),
        organizer_id INT,
        organizer_name VARCHAR(255),
        image_url TEXT,
        ticket_categories JSON
    );`
	db.MustExec(schema)

	return db
}

// Tes untuk fungsi Insert
// Test Case 1: positive case
func TestInsert_positive_case(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	repo := NewRepository(db)

	// data input
	newEvent := agoratix.Event{
		Title:            "Konser Spektakuler",
		Description:      "Deskripsi konser yang sangat menarik.",
		DateTime:         time.Now(),
		Location:         "Stadion GBK",
		Category:         "Musik",
		OrganizerId:      "123",
		OrganizerName:    "Event Organizer Pro",
		ImageUrl:         "http://example.com/image.jpg",
		TicketCategories: `[{"name": "VIP", "price": 500000}]`, // Contoh data JSON sebagai string
	}

	// Act
	result, err := repo.InsertEvent(newEvent)

	// Assert
	assert.NoError(t, err)                              // Pastikan tidak ada error
	assert.NotEqual(t, 0, result.ID)                    // Pastikan ID berhasil dibuat
	assert.Equal(t, "Konser Spektakuler", result.Title) // Cek salah satu field
	assert.Equal(t, "Stadion GBK", result.Location)     // Cek field lainnya untuk validasi lebih
}

// Test Case 2: negative case
func TestInsert_negative_case_invalid_organizer_id(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	repo := NewRepository(db)

	// Siapkan data input dengan OrganizerId yang TIDAK VALID (misal: 9999)
	newEvent := agoratix.Event{
		Title:            "Event Hantu",
		Description:      "Event ini tidak akan pernah dibuat.",
		DateTime:         time.Now(),
		Location:         "Tidak Diketahui",
		Category:         "Misteri",
		OrganizerId:      "99", // ID ini tidak ada di tabel organizers
		OrganizerName:    "Organizer Hantu",
		ImageUrl:         "http://example.com/ghost.jpg",
		TicketCategories: `[]`,
	}

	// Act (Aksi)
	result, err := repo.InsertEvent(newEvent)

	// Assert (Penegasan)
	assert.Error(t, err)                                            // HARUS mengembalikan error
	assert.Contains(t, err.Error(), "FOREIGN KEY constraint fails") // Opsional: Cek isi pesan error
	assert.Equal(t, 0, result.ID)                                   // ID HARUS 0 karena data gagal masuk
}
