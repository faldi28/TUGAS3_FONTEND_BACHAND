package models

type Sepatu struct {
	ID     int     `json:"id"`
	Merk   string  `json:"merk"`
	Jenis  string  `json:"jenis"`
	Ukuran int     `json:"ukuran"`
	Harga  float64 `json:"harga"`
}

