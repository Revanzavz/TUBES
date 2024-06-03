package main

import "fmt"

type EditBuku struct {
	judul, pengarang, isbn string
	eksemplar, terbit int
}

func editBuku(buku *tabBuku, i int, index int) {
	cetakBuku(*buku, i)
	var keyword string
	var found bool
	var idx EditBuku
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
		
		fmt.Scan(&idx.judul)
		if  idx.judul != "-" {
			(*buku)[index].judul = idx.judul
		}
		fmt.Print("Masukkan nama pengarang (ketik - jika tidak ingin diubah): ")
		
		fmt.Scan(&idx.pengarang)
		if idx.pengarang != "-" {
			(*buku)[index].pengarang = idx.pengarang
		}
		fmt.Print("Masukkan nomor ISBN (ketik -1 jika tidak ingin diubah): ")
		
		fmt.Scan(&idx.isbn)
		if idx.isbn != "-1" {
			(*buku)[index].nomorISBN = idx.isbn
		}
		fmt.Print("Masukkan jumlah eksemplar (ketik -1 jika tidak ingin diubah): ")
		
		fmt.Scan(&idx.eksemplar)
		if idx.eksemplar != -1 {
			(*buku)[index].jumlahEksemplar = idx.eksemplar
		}
		fmt.Print("Masukkan tahun terbit (ketik -1 jika tidak ingin diubah): ")
		
		fmt.Scan(&idx.terbit)
		if idx.terbit != -1 {
			(*buku)[index].tahunTerbit = idx.terbit
		}
		fmt.Println("Data buku berhasil diubah")
		fmt.Println(" ")
		clearscreen()
	} else {
		fmt.Println("Buku tidak ditemukan")
		fmt.Println(" ")
	}
}
