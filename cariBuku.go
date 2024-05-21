package main

import (
	"fmt"
)

func cariBuku(buku tabBuku, i int) {
	var caridatabuku int
	var keyword string
	var keywordint int
	var found bool

	fmt.Println("1. Judul buku")
	fmt.Println("2. Pengarang")
	fmt.Println("3. Nomor ISBN")
	fmt.Println("4. Tahun terbit")
	fmt.Println("5. Exit")
	fmt.Print("Masukkan data buku yang ingin dicari: ")

	fmt.Scan(&caridatabuku)
	switch caridatabuku {
	case 1:
		fmt.Print("Jika judul buku lebih dari 1 kata gunakan _ sebagai pengganti spasi\n")
		fmt.Print("Masukkan judul buku yang ingin dicari: ")
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
				found = true
			}
			if !found {
				fmt.Println("Buku tidak ditemukan")
				fmt.Println(" ")
			}
		}
	case 2:
		fmt.Print("Jika nama pengarang lebih dari 1 kata gunakan _ sebagai pengganti spasi\n")
		fmt.Print("Masukkan pengarang buku yang ingin dicari: ")
		fmt.Scan(&keyword)
		for j := 0; j < i; j++ {
			if buku[j].pengarang == keyword {
				fmt.Println("Buku ditemukan!")
				fmt.Println("Judul buku: ", buku[j].judul)
				fmt.Println("Pengarang: ", buku[j].pengarang)
				fmt.Println("Nomor ISBN: ", buku[j].nomorISBN)
				fmt.Println("Jumlah eksemplar: ", buku[j].jumlahEksemplar)
				fmt.Println("Tahun terbit: ", buku[j].tahunTerbit)
				fmt.Println(" ")
				found = true
			}
			if !found {
				fmt.Println("Buku tidak ditemukan")
				fmt.Println(" ")
			}
		}
	case 3:
		fmt.Print("Masukkan nomor ISBN buku yang ingin dicari: ")
		fmt.Scan(&keywordint)
		for j := 0; j < i; j++ {
			if buku[j].nomorISBN == keywordint {
				fmt.Println("Buku ditemukan!")
				fmt.Println("Judul buku: ", buku[j].judul)
				fmt.Println("Pengarang: ", buku[j].pengarang)
				fmt.Println("Nomor ISBN: ", buku[j].nomorISBN)
				fmt.Println("Jumlah eksemplar: ", buku[j].jumlahEksemplar)
				fmt.Println("Tahun terbit: ", buku[j].tahunTerbit)
				fmt.Println(" ")
				found = true
			}
			if !found {
				fmt.Println("Buku tidak ditemukan")
				fmt.Println(" ")
			}
		}
	case 4:
		fmt.Print("Masukkan tahun terbit buku yang ingin dicari: ")
		fmt.Scan(&keywordint)
		for j := 0; j < i; j++ {
			if buku[j].tahunTerbit == keywordint {
				fmt.Println("Buku ditemukan!")
				fmt.Println("Judul buku: ", buku[j].judul)
				fmt.Println("Pengarang: ", buku[j].pengarang)
				fmt.Println("Nomor ISBN: ", buku[j].nomorISBN)
				fmt.Println("Jumlah eksemplar: ", buku[j].jumlahEksemplar)
				fmt.Println("Tahun terbit: ", buku[j].tahunTerbit)
				fmt.Println(" ")
				found = true
			}
		}
		if !found {
			fmt.Println("Buku tidak ditemukan")
			fmt.Println(" ")
		}
	case 5:
		return
	default:
		fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia")

	}

}