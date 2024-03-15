package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Status struct untuk menyimpan status water dan wind
type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

// Data struct untuk menyimpan data status
type Data struct {
	Status Status `json:"status"`
}

// getStatusWaterStatus mengembalikan status berdasarkan nilai water
func getStatusWaterStatus(water int) string {
	switch {
	case water < 5:
		return "Aman"
	case water >= 6 && water <= 8:
		return "Siaga"
	default:
		return "Bahaya"
	}
}

// getStatusWindStatus mengembalikan status berdasarkan nilai wind
func getStatusWindStatus(wind int) string {
	switch {
	case wind < 6:
		return "Aman"
	case wind >= 7 && wind <= 15:
		return "Siaga"
	default:
		return "Bahaya"
	}
}

// updateStatus secara periodik mengupdate file JSON dengan data status yang baru
func updateStatus() {
	for {
		status := Status{
			Water: rand.Intn(100) + 1,
			Wind:  rand.Intn(100) + 1,
		}

		data := Data{Status: status}

		jsonData, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		file, err := os.Create("status.json")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		_, err = file.Write(jsonData)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("Status updated:", data)

		time.Sleep(15 * time.Second)
	}
}

func main() {
	// Mulai goroutine untuk memperbarui status secara periodik
	go updateStatus()

	// Serve HTTP untuk menampilkan status
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Baca file status.json
		file, err := os.Open("status.json")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Decode JSON
		var data Data
		err = json.NewDecoder(file).Decode(&data)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Menampilkan status
		waterStatus := getStatusWaterStatus(data.Status.Water)
		windStatus := getStatusWindStatus(data.Status.Wind)

		fmt.Fprintf(w, "Status Air: %d meter (%s)\n", data.Status.Water, waterStatus)
		fmt.Fprintf(w, "Status Angin: %d meter/detik (%s)\n", data.Status.Wind, windStatus)
	})

	// Mulai server HTTP
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
