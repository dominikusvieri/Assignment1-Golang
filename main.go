package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dominikusvieri/Assignment1-Golang/execute"
	"github.com/dominikusvieri/Assignment1-Golang/models"
	"github.com/dominikusvieri/Assignment1-Golang/parameter"
)

func main() {
	var siswa []models.Siswa

	if len(os.Args) > 1 {
		RequestNumber, err := strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println("Please make sure you set the number not with string!")
			return
		}

		siswa1 := parameter.SetNewSiswa(1, "Domi", "Jakarta", "Frontend", "Ingin banyak belajar")
		siswa2 := parameter.SetNewSiswa(2, "Domi", "Jakarta", "Frontend", "Ingin banyak belajar")
		siswa3 := parameter.SetNewSiswa(3, "Domi", "Jakarta", "Frontend", "Ingin banyak belajar")

		siswa = append(siswa, execute.SetSiswa(siswa1))
		siswa = append(siswa, execute.SetSiswa(siswa2))
		siswa = append(siswa, execute.SetSiswa(siswa3))

		for _, s := range siswa {
			s.Print(RequestNumber)
		}
	} else {
		fmt.Println("Please set the number!")
	}

}
