// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bellatrijuliana/agoratix-app/factory" // Impor factory Anda
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 1. Siapkan "Bahan Baku"
	dsn := "user:215544@tcp(127.0.0.1:3306)/agoratix?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("database is not responding:", err)
	}
	fmt.Println("Successfully connected to the database")

	// 2. Pesan "Mobil Jadi" dari Pabrik
	e := factory.Initialize(db)

	// 3. Nyalakan Mesin
	log.Fatal(e.Start(":8080"))
}
