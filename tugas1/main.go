package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {

	// Mencegah error ketika tidak menginput angka absen
	if len(os.Args) <= 1 {
		errKosong()
	} else {
		//Menampung angka dari terminal
		argsRaw := os.Args
		num := argsRaw[1]

		//Merubah num string menjadi int
		number, _ := strconv.Atoi(num)

		//Data siswa
		var dataStudent = []Student{
			{
				nama:      "Yusuf",
				alamat:    "Bandung",
				pekerjaan: "Fresh Graduate",
				alasan:    "Mengenal Golang",
			},
			{
				nama:      "Jaka",
				alamat:    "Purwakarta",
				pekerjaan: "Fresh Graduate",
				alasan:    "Mencari ilmu baru",
			},
		}

		//Mencegah error apabila angka absen tidak sesuai dengan data
		if number > len(dataStudent) {
			err()
		} else {
			showData(dataStudent[number-1])
		}
	}
}

// Fungsi Error ketika absen tidak tersedia
func err() {
	fmt.Println("Data tidak tersedia!")
	return
}

// Fungsi untuk menampilkan data pada CLI
func showData(student Student) {
	fmt.Println("Nama : ", student.nama)
	fmt.Println("Alamat : ", student.alamat)
	fmt.Println("Pekerjaan : ", student.pekerjaan)
	fmt.Println("Alasan : ", student.alasan)
}

// Fungsi Error ketika tidak memberikan input pada CLI
func errKosong() {
	fmt.Println("Pastikan memberikan angka absen!")
	return
}
