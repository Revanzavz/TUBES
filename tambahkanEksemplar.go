package main

import (
	"fmt"
)

func tambahkanEksemplar(buku *tabBuku, i int) {
	cetakBuku(*buku, i)
	fmt.Println("Masukkan judul buku yang ingin ditambahkan eksemplarnya: ")
	var keyword string
	fmt.Scan(&keyword)
	for j := 0; j < i; j++ {
		if buku[j].judul == keyword {
			fmt.Println("Buku ditemukan!")
			fmt.Println("Judul buku: ", buku[j].judul)
			fmt.Println("Pengarang: ", buku[j].pengarang)
			fmt.Println("Nomor ISBN: ", buku[j].nomorISBN)
			fmt.Println("Jumlah eksemplar: ", buku[j].jumlahEksemplar)
			fmt.Println("Tahun terbit: ", buku[j].tahunTerbit)
			fmt.Println(" ")

			var tambah int
			fmt.Println("Masukkan jumlah eksemplar yang ingin ditambahkan: ")
			fmt.Scan(&tambah)
			buku[j].jumlahEksemplar += tambah
			fmt.Println("Jumlah eksemplar berhasil ditambahkan")
			fmt.Println("Jumlah eksemplar saat ini: ", buku[j].jumlahEksemplar)
			fmt.Println(" ")
		} else {
			fmt.Println("Buku tidak ditemukan")
			fmt.Println(" ")
		}
	}
}