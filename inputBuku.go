package main

import (
	"fmt"
)

func inputBuku(buku *tabBuku, i *int) {
	var again string
	fmt.Println("Masukkan data buku:")
	for again != "n" && *i < NMAX {
		fmt.Print("Jika buku lebih dari 1 kata gunakan _ sebagai pengganti spasi\n")
		fmt.Print("Masukkan judul buku: ")
		fmt.Scan(&buku[*i].judul)
		fmt.Print("Jika nama pengarang lebih dari 1 kata gunakan _ sebagai pengganti spasi\n")
		fmt.Print("Masukkan nama pengarang: ")
		fmt.Scan(&buku[*i].pengarang)
		fmt.Print("Masukkan nomor ISBN: ")
		fmt.Scan(&buku[*i].nomorISBN)
		fmt.Print("Masukkan jumlah eksemplar: ")
		fmt.Scan(&buku[*i].jumlahEksemplar)
		fmt.Print("Masukkan tahun terbit: ")
		fmt.Scan(&buku[*i].tahunTerbit)

		for j := 0; j < *i; j++ {
			if buku[j].judul == buku[*i].judul {
				buku[*i].jumlahEksemplar += buku[j].jumlahEksemplar
			}
		}

		*i++
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