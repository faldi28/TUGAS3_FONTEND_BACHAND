package routes

import (
	"net/http"
	"penjualan-sepatu/controllers"
	"penjualan-sepatu/middleware"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
    r := mux.NewRouter()

    // Gunakan middleware EnableCORS
    r.Use(middleware.EnableCORS)

    // Rute API
    r.HandleFunc("/sepatu", controllers.TampilkanSepatu).Methods(http.MethodGet)
    r.HandleFunc("/sepatu", controllers.TambahSepatu).Methods(http.MethodPost)
    r.HandleFunc("/sepatu/{id}", controllers.HapusSepatu).Methods(http.MethodDelete)
    r.HandleFunc("/sepatu/{id}", controllers.UpdateSepatu).Methods(http.MethodPut)

    return r
}