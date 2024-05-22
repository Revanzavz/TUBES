package main

import (
	"fmt"
)

// Define the structure for tabBuku
// type Buku struct {
// 	judul           string
// 	pengarang       string
// 	nomorISBN       string
// 	jumlahEksemplar int
// 	tahunTerbit     int
// }

func inputBuku(buku *tabBuku, i *int) {
	var again string
	fmt.Println("Masukkan data buku:")

	for again != "n" && *i < NMAX {
		fmt.Print("Jika buku lebih dari 1 kata gunakan _ sebagai pengganti spasi\n")
		fmt.Print("Masukkan judul buku: ")
		var judul string
		fmt.Scan(&judul)

		fmt.Print("Jika nama pengarang lebih dari 1 kata gunakan _ sebagai pengganti spasi\n")
		fmt.Print("Masukkan nama pengarang: ")
		var pengarang string
		fmt.Scan(&pengarang)

		fmt.Print("Masukkan nomor ISBN: ")
		var nomorISBN string
		fmt.Scan(&nomorISBN)

		fmt.Print("Masukkan jumlah eksemplar: ")
		var jumlahEksemplar int
		fmt.Scan(&jumlahEksemplar)

		fmt.Print("Masukkan tahun terbit: ")
		var tahunTerbit int
		fmt.Scan(&tahunTerbit)

		//found := false
		for j := 0; j < *i; j++ {
			if buku[j].judul == judul && buku[j].pengarang == pengarang && buku[j].nomorISBN == nomorISBN {
				buku[j].jumlahEksemplar += jumlahEksemplar
			//	found = true
			}
		}


		fmt.Println("Masukkan data buku lagi? (y/n) ")
		fmt.Scan(&again)

		for again != "n" && again != "y" {
			fmt.Println("Input tidak valid, silahkan input kembali")
			fmt.Scan(&again)
		}

		clearscreen()
		fmt.Println(" ")
	}
}
