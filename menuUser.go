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
			cariBuku(buku, i)
		case 2:
			clearscreen()
			pinjamBuku(&buku, &i, &riwayatPeminjaman, &jumlahPeminjaman)
		case 3:
			kembalikanBuku(&buku, i, &riwayatPeminjaman, jumlahPeminjaman)
		case 4:
			lihatRiwayat(riwayatPeminjaman, jumlahPeminjaman)
		case 5:
			return
		default:
			fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia")
		}
	}
}