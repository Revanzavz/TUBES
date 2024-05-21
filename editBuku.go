package main

import "fmt"

func editBuku(buku *tabBuku, i int, index int) {
	var keyword string
	var found bool
	fmt.Print("Masukkan judul buku yang ingin diubah: ")
	fmt.Scan(&keyword)
	for j := 0; j < i; j++ {
		if (*buku)[j].judul == keyword {
			index = j
			found = true
		}
	}
	if found {
		fmt.Println("Masukkan data buku baru:")
		fmt.Print("Masukkan judul buku (ketik - jika tidak ingin diubah): ")
		fmt.Scan(&(*buku)[index].judul)
		if (*buku)[index].judul == "-" {
			(*buku)[index].judul = buku[index].judul
		}
		fmt.Print("Masukkan nama pengarang (ketik - jika tidak ingin diubah): ")
		fmt.Scan(&(*buku)[index].pengarang)
		if (*buku)[index].pengarang == "-" {
			(*buku)[index].pengarang = buku[index].pengarang
		}
		fmt.Print("Masukkan nomor ISBN (ketik -1 jika tidak ingin diubah): ")
		fmt.Scan(&(*buku)[index].nomorISBN)
		if (*buku)[index].nomorISBN == "-" {
			(*buku)[index].nomorISBN = buku[index].nomorISBN
		fmt.Print("Masukkan jumlah eksemplar (ketik -1 jika tidak ingin diubah): ")
		fmt.Scan(&(*buku)[index].jumlahEksemplar)
		if (*buku)[index].jumlahEksemplar == -1 {
			(*buku)[index].jumlahEksemplar = buku[index].jumlahEksemplar
		}
		fmt.Print("Masukkan tahun terbit (ketik -1 jika tidak ingin diubah): ")
		fmt.Scan(&(*buku)[index].tahunTerbit)
		if (*buku)[index].tahunTerbit == -1 {
			(*buku)[index].tahunTerbit = buku[index].tahunTerbit
		}
		fmt.Println("Data buku berhasil diubah")
		fmt.Println(" ")
	} else {
		fmt.Println("Buku tidak ditemukan")
		fmt.Println(" ")
	}
}
}