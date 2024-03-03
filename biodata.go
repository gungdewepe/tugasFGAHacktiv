package main

import (
	"fmt"
	"os"
)

// Teman adalah struktur untuk menyimpan data teman
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// DatabaseTeman adalah slice untuk menyimpan data teman-teman
var DatabaseTeman = []Teman{
	{Nama: "Muhammad", Alamat: "Depok", Pekerjaan: "Staff IT", Alasan: "Belajar"},
	{Nama: "Agung", Alamat: "Bogor", Pekerjaan: "CS", Alasan: "Lolos Saja"},
	{Nama: "Dwi", Alamat: "Jakarta", Pekerjaan: "Manager", Alasan: "Gabut"},
	{Nama: "Prasetiyo", Alamat: "Tangerang", Pekerjaan: "Teller", Alasan: "Nambah Skill"},
	// Tambahkan data teman selanjutnya di sini
}

// TampilkanDataTeman adalah fungsi untuk menampilkan data teman berdasarkan nomor absen
func TampilkanDataTeman(nomor int) {
	if nomor <= 0 || nomor > len(DatabaseTeman) {
		fmt.Println("Nomor absen tidak valid")
		return
	}

	teman := DatabaseTeman[nomor-1]
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Gunakan: go run biodata.go [nomor_absen]")
		return
	}

	nomorAbsen := os.Args[1]
	var nomor int
	_, err := fmt.Sscanf(nomorAbsen, "%d", &nomor)
	if err != nil {
		fmt.Println("Nomor absen harus berupa bilangan bulat")
		return
	}

	TampilkanDataTeman(nomor)
}
