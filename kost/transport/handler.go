package transport

import (
	"encoding/json"
	"net/http"
)

type Kost struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Harga  int    `json:"harga"`
}

func GetKostsHandler(w http.ResponseWriter, r *http.Request) {
	kosts := []Kost{
		{ID: 1, Nama: "Kost Mawar", Alamat: "Jl. UG Raya", Harga: 1000000},
		{ID: 2, Nama: "Kost Melati", Alamat: "Jl. UG Timur", Harga: 850000},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kosts)
}