package main

import (
	"fmt"
)

func lihatRiwayat(riwayat tabPeminjaman, jmlP int) {
	if jmlP == 0 {
		fmt.Println("Belum ada riwayat peminjaman.")
		fmt.Println(" ")
		
	}
	fmt.Println("Riwayat Peminjaman:")
	for j := 0; j < jmlP; j++ {
		fmt.Println("Peminjaman ke-", j+1)
		fmt.Println("Judul Buku: ", riwayat[j].judul)
		fmt.Println("Nama Peminjam: ", riwayat[j].namaPeminjam)
		fmt.Println("Tanggal Peminjaman: ", riwayat[j].tanggalPeminjaman.Format("02 January 2006"))
		fmt.Println()
	}
}