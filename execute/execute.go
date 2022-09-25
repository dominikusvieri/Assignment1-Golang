package execute

import (
	"github.com/dominikusvieri/Assignment1-Golang/models"
	"github.com/dominikusvieri/Assignment1-Golang/parameter"
)

func SetSiswa(req *parameter.SetSiswa) models.Siswa {
	var siswa models.Siswa

	siswa.Absen = req.Absen
	siswa.Nama = req.Nama
	siswa.Alamat = req.Alamat
	siswa.Pekerjaan = req.Pekerjaan
	siswa.Alasan = req.Alasan

	return siswa
}
