package transport

import (
	"backend/database"
	"encoding/json"
	"net/http"
)

type Kost struct {
	ID         int     `json:"id"`
	Nama       string  `json:"nama"`
	Alamat     string  `json:"alamat"`
	Fasilitas  string  `json:"fasilitas"`
	Harga      string  `json:"harga"`
	GambarURL  string  `json:"gambar_url"`
	MapsURL    string  `json:"maps_url"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

func GetKostsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`
		SELECT id, nama, alamat, fasilitas, harga, gambar_url, maps_url, latitude, longitude
		FROM kosts`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var kosts []Kost
	for rows.Next() {
		var k Kost
		err := rows.Scan(&k.ID, &k.Nama, &k.Alamat, &k.Fasilitas, &k.Harga, &k.GambarURL, &k.MapsURL, &k.Latitude, &k.Longitude)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		kosts = append(kosts, k)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kosts)
}