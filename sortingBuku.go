package main

import (
	"fmt"
)

func sortingBuku(buku *tabBuku, i int) {
	var choiceSort int
	for choiceSort != 5 {
		fmt.Println("Menu Sorting:")
		fmt.Println("1. Sorting berdasarkan judul secara ascending")
		fmt.Println("2. sorting berdasarkan judul secara descending")
		fmt.Println("3. Sorting berdasarkan tahun terbit secara ascending")
		fmt.Println("4. Sorting berdasarkan tahun terbit secara descending")
		fmt.Println("5. Exit")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&choiceSort)
		fmt.Print("\033[2J")
		clearscreen()
		switch choiceSort {
		case 1:
			//selection sort ascending
			pass := 1
			for pass < i {
				idx := pass - 1
				j := pass
				for j < i {
					if (*buku)[idx].judul > (*buku)[j].judul {
						idx = j
					}
					j++
				}
				temp := (*buku)[pass-1]
				(*buku)[pass-1] = (*buku)[idx]
				(*buku)[idx] = temp
				pass++
			}
			cetakBuku(*buku, i)
		case 2:
			//selection sort descending
			pass := 1
			for pass < i {
				idx := pass - 1
				j := pass
				for j < i {
					if (*buku)[idx].judul < (*buku)[j].judul {
						idx = j
					}
					j++
				}
				temp := (*buku)[pass-1]
				(*buku)[pass-1] = (*buku)[idx]
				(*buku)[idx] = temp
				pass++
			}
			cetakBuku(*buku, i)
		case 3:
			//insertion sort ascending
			pass := 1
			for pass < i {
				j := pass
				temp := (*buku)[pass]
				for j > 0 && temp.tahunTerbit < (*buku)[j-1].tahunTerbit {
					(*buku)[j] = (*buku)[j-1]
					j--
				}
				(*buku)[j] = temp
				pass++
			}
			cetakBuku(*buku, i)
		case 4:
			//insertion sort descending
			pass := 1
			for pass < i {
				j := pass
				temp := (*buku)[pass]
				for j > 0 && temp.tahunTerbit > (*buku)[j-1].tahunTerbit {
					(*buku)[j] = (*buku)[j-1]
					j--
				}
				(*buku)[j] = temp
				pass++
			}
			cetakBuku(*buku, i)
		case 5:
			return
		default:
			fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia")
		}

	}
}
