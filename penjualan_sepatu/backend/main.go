package main

import (
	"log"
	"net/http"

	"penjualan-sepatu/config"
	"penjualan-sepatu/routes"
)

func main() {
	// Koneksi ke database
	config.KoneksiDatabase()

	// Dapatkan router dengan middleware CORS
	r := routes.Routes()

	// Jalankan server pada port 8080
	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
