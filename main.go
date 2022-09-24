package main

import (
	"fmt"
)

type Peserta struct {
	ID        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	var absen int
	fmt.Printf("Masukkan nomor Absensi: ")
	fmt.Scanln(&absen)

	var peserta = []Peserta{
		{
			ID:        1,
			Nama:      "Restu",
			Alamat:    "Pasteur",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Seru",
		},
		{
			ID:        2,
			Nama:      "Budi",
			Alamat:    "Samarinda",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Seru",
		},
		{
			ID:        3,
			Nama:      "Fathur",
			Alamat:    "Sangatta",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Seru",
		},
	}

	AbsenPeserta(peserta, absen)
}

func AbsenPeserta(peserta []Peserta, absen int) {

	for i, s := range peserta {
		if s.ID == absen {
			fmt.Println("Absen-Ke						: ", peserta[i].ID)
			fmt.Println("Nama							: ", peserta[i].Nama)
			fmt.Println("Alamat							: ", peserta[i].Alamat)
			fmt.Println("Pekerjaan						: ", peserta[i].Pekerjaan)
			fmt.Println("Alasan							: ", peserta[i].Alasan)
			return
		}
	}
}
