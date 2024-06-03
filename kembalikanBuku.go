package main

import (
	"fmt"
	"time"
)

func kembalikanBuku(buku *tabBuku, i int, riwayat *tabPeminjaman, jmlP int) {
	var keyword, namaPeminjam string
	var hari, bulan, tahun int
	var found bool
	fmt.Print("Masukkan judul buku yang ingin dikembalikan: ")
	fmt.Scan(&keyword)
	fmt.Print("Masukkan nama peminjam: ")
	fmt.Scan(&namaPeminjam)
	for j := 0; j < i; j++ {
		if (*buku)[j].judul == keyword  {
			for k := 0; k < jmlP; k++ {
				if (*riwayat)[k].judul == keyword && (*riwayat)[k].namaPeminjam == namaPeminjam {
					found = true
					(*buku)[j].jumlahEksemplar++
					fmt.Println("Buku ditemukan!")
					fmt.Println("Judul buku: ", (*buku)[j].judul)
					fmt.Println("Pengarang: ", (*buku)[j].pengarang)
					fmt.Println("Nomor ISBN: ", (*buku)[j].nomorISBN)
					fmt.Println("Jumlah eksemplar: ", (*buku)[j].jumlahEksemplar)
					fmt.Println("Tahun terbit: ", (*buku)[j].tahunTerbit)
					fmt.Println(" ")

					// Input tanggal pengembalian
					fmt.Print("Masukkan tanggal pengembalian (dd mm yyyy): ")
					fmt.Scan(&hari, &bulan, &tahun)
					for hari > 31 || bulan > 12 {
						fmt.Println("Tanggal atau Bulan tidak valid!")
						fmt.Println("Periksa Kembali!")
						fmt.Print("Masukkan tanggal pengembalian (dd mm yyyy): ")
						fmt.Scan(&hari, &bulan, &tahun)
					}

					tanggalPengembalian := time.Date(tahun, time.Month(bulan), hari, 0, 0, 0, 0, time.UTC)

					tanggalPeminjaman := (*riwayat)[k].tanggalPeminjaman
					durasi := tanggalPengembalian.Sub(tanggalPeminjaman).Hours() / 24
					if durasi > 0 && durasi < 7 {
						fmt.Println("Durasi Peminjaman: ", durasi, "hari")
					}
					if durasi >= 7 {
						fmt.Println("Anda terlambat mengembalikan buku selama", durasi, "hari")
						fmt.Println("Denda yang harus dibayar sebesar Rp", (durasi-7)*1000)
					} else {
						(*buku)[j].jumlahEksemplar--
						fmt.Print("Pengembalian seharusnya lebih dari masa tanggal peminjaman\nSilahkan input ulang")
					}
				}
			}
		}
	}
	if !found {
		fmt.Println("Buku tidak ditemukan atau nama peminjam tidak sesuai")
		fmt.Println(" ")
	}
}

func hapusBuku(buku *tabBuku, i *int, index int) {
	cetakBuku(*buku, *i)
	var keyword string
	var found bool
	fmt.Print("Masukkan judul buku yang ingin dihapus: ")
	fmt.Scan(&keyword)
	for j := 0; j < *i; j++ {
		if (*buku)[j].judul == keyword {
			index = j
			found = true
		}
	}
	if found {
		for k := index; k < *i; k++ {
			(*buku)[k] = (*buku)[k+1]
		}
		*i--
		fmt.Println("Buku berhasil dihapus")
		fmt.Println(" ")
	} else {
		fmt.Println("Buku tidak ditemukan")
		fmt.Println(" ")
	}
}