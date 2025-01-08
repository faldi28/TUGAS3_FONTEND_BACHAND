package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Fungsi untuk menghubungkan ke database
func KoneksiDatabase() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/penjualan_sepatu") // Ganti dengan kredensial yang sesuai
	if err != nil {
		log.Fatal(err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database terkoneksi dengan baik")
}
