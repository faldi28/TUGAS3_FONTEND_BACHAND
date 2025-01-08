package controllers

import (
	"encoding/json"
	"net/http"
	"penjualan-sepatu/config"
	"penjualan-sepatu/models"
	"strconv"
	"github.com/gorilla/mux"
)

func TampilkanSepatu(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM sepatu")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sepatuList []models.Sepatu
	for rows.Next() {
		var sepatu models.Sepatu
		if err := rows.Scan(&sepatu.ID, &sepatu.Merk, &sepatu.Jenis, &sepatu.Ukuran, &sepatu.Harga); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sepatuList = append(sepatuList, sepatu)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sepatuList)
}

func TambahSepatu(w http.ResponseWriter, r *http.Request) {
	var sepatu models.Sepatu
	if err := json.NewDecoder(r.Body).Decode(&sepatu); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := config.DB.Exec("INSERT INTO sepatu (merk, jenis, ukuran, harga) VALUES (?, ?, ?, ?)",
		sepatu.Merk, sepatu.Jenis, sepatu.Ukuran, sepatu.Harga)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Gagal mendapatkan ID terakhir", http.StatusInternalServerError)
		return
	}
	sepatu.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sepatu)
}

func HapusSepatu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("DELETE FROM sepatu WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Data berhasil dihapus"}`))
}

func UpdateSepatu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	var sepatu models.Sepatu
	if err := json.NewDecoder(r.Body).Decode(&sepatu); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("UPDATE sepatu SET merk = ?, jenis = ?, ukuran = ?, harga = ? WHERE id = ?",
		sepatu.Merk, sepatu.Jenis, sepatu.Ukuran, sepatu.Harga, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sepatu.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sepatu)
}
