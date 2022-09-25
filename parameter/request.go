package parameter

type SetSiswa struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func SetNewSiswa(absen int, nama string, alamat string, pekerjaan string, alasan string) *SetSiswa {
	return &SetSiswa{
		Absen:     absen,
		Nama:      nama,
		Alamat:    alamat,
		Pekerjaan: pekerjaan,
		Alasan:    alamat,
	}
}
