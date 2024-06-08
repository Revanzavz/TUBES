package main

import "fmt"

func MenuUser() {
	var choiceUser int
	fmt.Println("Welcome User!")
	for choiceUser != 5 {
		fmt.Println("Menu User:")
		fmt.Println("1. Cari data buku")
		fmt.Println("2. Pinjam buku")
		fmt.Println("3. Kembalikan buku")
		fmt.Println("4. Lihat riwayat peminjaman")
		fmt.Println("5. Exit")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&choiceUser)
		fmt.Println()
		switch choiceUser {
		case 1:
			if i == 0 {
				fmt.Println("Data buku masih kosong")
				fmt.Println("Input Data Buku terlebih dahulu dengan memilih menu 1")
				fmt.Println(" ")
			} else {
			clearscreen()
			cariBuku(buku, i)
			}
		case 2:
			if i == 0 {
				fmt.Println("Data buku masih kosong")
				fmt.Println("Input Data Buku terlebih dahulu dengan memilih menu 1")
				fmt.Println(" ")
			} else {
			clearscreen()
			cetakBuku(buku, i)
			pinjamBuku(&buku, &i, &riwayatPeminjaman, &jumlahPeminjaman)
			}
		case 3:
			if i == 0 {
				fmt.Println("Data buku masih kosong")
				fmt.Println("Input Data Buku terlebih dahulu dengan memilih menu 1")
				fmt.Println(" ")
			} else {
			clearscreen()
			kembalikanBuku(&buku, i, &riwayatPeminjaman, jumlahPeminjaman)
			}
		case 4:
			if i == 0 {
				fmt.Println("Data buku masih kosong")
				fmt.Println("Input Data Buku terlebih dahulu dengan memilih menu 1")
				fmt.Println(" ")
			} else {
			clearscreen()
			lihatRiwayat(riwayatPeminjaman, jumlahPeminjaman)
			}
		case 5:
			return
		default:
			fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia")
		}
	}
}