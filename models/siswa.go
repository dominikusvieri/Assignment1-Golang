package models

import "fmt"

type Siswa struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func (siswa *Siswa) Print(number int) {
	if number == siswa.Absen {
		fmt.Println("Absen \t\t : ", siswa.Absen)
		fmt.Println("Nama \t\t : ", siswa.Nama)
		fmt.Println("Alamat \t\t : ", siswa.Alamat)
		fmt.Println("Pekerjaan \t : ", siswa.Pekerjaan)
		fmt.Println("Alasan \t\t : ", siswa.Alasan)
	}
}
