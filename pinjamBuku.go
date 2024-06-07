package main

import (
	"fmt"
	"time"
)

func pinjamBuku(buku *tabBuku, i *int, riwayat *tabPeminjaman, jmlP *int) {
	var keyword, namaPeminjam string
	var hari, bulan, tahun int
	var found bool

	fmt.Println("Masukkan judul buku yang ingin dipinjam: ")
	fmt.Println("Gunakan tanda '_' untuk menggantikan spasi")
	fmt.Scan(&keyword)
	clearscreen()
	for j := 0; j < *i; j++ {
		if (*buku)[j].judul == keyword {
			found = true

			if (*buku)[j].jumlahEksemplar == 0 {
				fmt.Println("Buku telah habis")
				return
			}

			fmt.Println("Buku ditemukan!")
			fmt.Println("Judul buku: ", (*buku)[j].judul)
			fmt.Println("Pengarang: ", (*buku)[j].pengarang)
			fmt.Println("Nomor ISBN: ", (*buku)[j].nomorISBN)
			fmt.Println("Jumlah eksemplar: ", (*buku)[j].jumlahEksemplar)
			fmt.Println("Tahun terbit: ", (*buku)[j].tahunTerbit)
			fmt.Println(" ")

			// Input data peminjam
			fmt.Print("Masukkan nama peminjam: ")
			fmt.Scan(&namaPeminjam)
			fmt.Print("Masukkan tanggal peminjaman (dd mm yyyy): ")
			fmt.Scan(&hari, &bulan, &tahun)
			for hari > 31 || bulan > 12 || (bulan == 02 && hari > 28) || (bulan%2 == 0 && hari > 30) || (tahun / 400 == 0 && tahun / 100 != 0) || tahun % 4 != 0 && bulan == 02 && hari > 28 {
				fmt.Println("Tanggal tidak valid")
				fmt.Print("Masukkan tanggal peminjaman (dd mm yyyy): ")
				fmt.Scan(&hari, &bulan, &tahun)
			}
			clearscreen()
			tanggalPeminjaman := time.Date(tahun, time.Month(bulan), hari, 0, 0, 0, 0, time.UTC)

			// Simpan riwayat peminjaman
			(*riwayat)[*jmlP] = dataPeminjaman{
				judul:             (*buku)[j].judul,
				namaPeminjam:      namaPeminjam,
				tanggalPeminjaman: tanggalPeminjaman,
			}
			(*jmlP)++

			// Kurangi jumlah eksemplar buku
			(*buku)[j].jumlahEksemplar--
			if (*buku)[j].jumlahEksemplar == 0 {
				fmt.Println("Mohon maaf Buku yang anda ingin pinjam belum tersedia")
			}
		}
	}
	//clearscreen()
	if !found {
		fmt.Println("Buku tidak ditemukan")
		fmt.Println(" ")
	} else {
		fmt.Print("Mohon diingat pengembalian buku tidak lebih dari 7 hari dari tanggal peminjaman\nJika lebih dari 7 hari akan dikenakan denda sebesar Rp 1000/hari\n")
		fmt.Println(" ")
	}
}
