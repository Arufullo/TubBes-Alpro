package main

import "fmt"

const NMAX int = 100

type data struct {
	idTransaksi string
	namaBarang  string
	jumlah      int
	totalHarga  int
}
type dataTransaksi [NMAX]data

//Func Utama
func main() {
	var pilihan, jumlahData, idx int
	var A dataTransaksi

	for pilihan < 7 {
		menu()
		fmt.Scan(&pilihan)
		fmt.Println()
		switch pilihan {
		case 1:
			baca(&A, &jumlahData)
		case 2:
			cetakData(A, jumlahData)
		case 3:
			search(&A, &jumlahData, &idx)
		case 4:
			editData(&A, &jumlahData, &idx)
		case 5:
			hapus(&A, &jumlahData, &idx)
		case 6:
			reset(&A, &jumlahData)
		}
	}
}

func menu() { // Tampilan interaktif
	fmt.Println("+-------------------------+")
	fmt.Println("|     K O P E R A S I     |")
	fmt.Println("|-------------------------|")
	fmt.Println("| 1. Input Data Transaksi |")
	fmt.Println("| 2. Cetak Data Transaksi |")
	fmt.Println("| 3. Search               |")
	fmt.Println("| 4. Edit Data            |")
	fmt.Println("| 5. Hapus Data           |")
	fmt.Println("| 6. Reset Data           |")
	fmt.Println("| 7. Exit                 |")
	fmt.Println("+-------------------------+")
	fmt.Print("Pilih (1/2/3/4/5/6/7)? ")
}

//Func baca
func baca(A *dataTransaksi, jumlahData *int) {
	var namaBarang string
	var jumlah, totalHarga int
	var stop bool = false
	fmt.Println("Silakan masukan 0 0 0 jika ingin berhenti")
	fmt.Println("Nama barang | Jumlah | Total Harga")
	for !stop {
		fmt.Scan(&namaBarang, &jumlah, &totalHarga)
		if namaBarang == "0" && jumlah == 0 && totalHarga == 0 {
			stop = true
		} else {
			A[*jumlahData].idTransaksi = fmt.Sprintf("KOP%04d", *jumlahData+1) // penginputan id transaksi dilakukan secara otomatis
			A[*jumlahData].namaBarang = namaBarang
			A[*jumlahData].jumlah = jumlah
			A[*jumlahData].totalHarga = totalHarga
			*jumlahData++
		}
	}
	fmt.Println()
}

//Func Cetak
func headingTable() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Printf("| %-12s | %-12s | %-6s | %-11s |\n", "ID Transaksi", "Nama Barang", "Jumlah", "Total Harga")
	fmt.Println("|--------------|--------------|--------|-------------|")
}

func footerTable() {
	fmt.Println("+----------------------------------------------------+")
}

func cetak(A dataTransaksi, i int) {
	fmt.Printf("| %-12s | %-12s | %-6d | %-11d |\n", A[i].idTransaksi, A[i].namaBarang, A[i].jumlah, A[i].totalHarga)
}

func cetakData(A dataTransaksi, jumlahData int) {
	headingTable()
	for i := 0; i < jumlahData; i++ {
		cetak(A, i)
	}
	fmt.Println("|--------------------------------------|-------------|")
	fmt.Printf("| %-36s | %-11d |\n", "Total Nilai Transaksi", totalNilaiTransaksi(A, jumlahData))
	footerTable()
	fmt.Println()
}

func totalNilaiTransaksi(A dataTransaksi, jumlahData int) int {
	var total int
	for i := 0; i < jumlahData; i++ {
		total = total + A[i].totalHarga
	}
	return total
}

//Func Cari
func search(A *dataTransaksi, jumlahData, i *int) {
	var pilihan int

	for pilihan < 3 {
		menuSearch()
		fmt.Scan(&pilihan)
		fmt.Println()
		switch pilihan {
		case 1:
			cariIdTransaksi(A, jumlahData, i)
		case 2:
			cariNamaBarang(A, jumlahData)
		}
	}
}

func menuSearch() {
	fmt.Println("+-----------------------------+")
	fmt.Println("|         S E A R C H         |")
	fmt.Println("|-----------------------------|")
	fmt.Println("| 1. Berdasarkan ID Transaksi |")
	fmt.Println("| 2. Berdasarkan Nama Barang  |")
	fmt.Println("| 3. Exit                     |")
	fmt.Println("+-----------------------------+")
	fmt.Print("Pilih (1/2/3)? ")
}

func cariIdTransaksi(A *dataTransaksi, jumlahData, i *int) {
	var n string

	fmt.Print("Masukan ID Transaksi (contoh: KOP0001): ")
	fmt.Scan(&n)

	ketemu := false
	*i = searchIdx(*A, *jumlahData, n)

	if *i > -1 {
		headingTable()
		cetak(*A, *i)
		footerTable()
		ketemu = true
	}
	if !ketemu {
		fmt.Println("Maaf, data tidak ditemukan")
	}
	fmt.Println()
}

func searchIdx(A dataTransaksi, jD int, n string) int {
	var low, mid, high, idx int

	low = 0
	high = jD - 1
	idx = -1

	for low <= high && idx == -1 {
		mid = (low + high) / 2
		if A[mid].idTransaksi == n {
			idx = mid
		} else if A[mid].idTransaksi < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}

func cariNamaBarang(A *dataTransaksi, jumlahData *int) {
	var n string
	var ketemu bool = false
	fmt.Print("Masukan Nama Barang (contoh: Tempat_Pensil): ")
	fmt.Scan(&n)

	headingTable()
	for i := 0; i < *jumlahData; i++ {
		if n == A[i].namaBarang {
			cetak(*A, *jumlahData)
			ketemu = true
		}
	}
	footerTable()

	if !ketemu {
		fmt.Println("Maaf, data tidak ditemukan")
	}
	fmt.Println()
}

func editData(A *dataTransaksi, jumlahData, idx *int) {
	var namaBarang string
	var jumlah, totalHarga int

	cariIdTransaksi(A, jumlahData, idx)
	if *idx > -1 {
		fmt.Println("Silakan edit data")
		fmt.Scan(&namaBarang, &jumlah, &totalHarga)
		if konfirmasi() == true {
			A[*idx].namaBarang = namaBarang
			A[*idx].jumlah = jumlah
			A[*idx].totalHarga = totalHarga
			fmt.Printf("Data ke %d sudah berhasil di edit\n", *idx+1)
			fmt.Println()
		}
	}
}

func hapus(A *dataTransaksi, jumlahData, idx *int) {
	cariIdTransaksi(A, jumlahData, idx)
	if *idx > -1 && konfirmasi() == true {
		for i := *idx; i < *jumlahData-1; i++ {
			A[i] = A[i+1]
		}
		A[*jumlahData-1] = data{}

		*jumlahData--
		fmt.Printf("Data ke %d sudah terhapus\n", *idx+1)
		fmt.Println()
	}
}

//Func Reset
func reset(A *dataTransaksi, jumlahData *int) {
	fmt.Println("Data yang akan dihapus: ")
	cetakData(*A, *jumlahData)

	if konfirmasi() == true {
		for i := 0; i < *jumlahData; i++ {
			A[i].idTransaksi = "0"
			A[i].namaBarang = "0"
			A[i].jumlah = 0
			A[i].totalHarga = 0
		}
		*jumlahData = 0
		fmt.Println("Semua data sudah terhapus")
		fmt.Println()
	}
}

func konfirmasi() bool {
	var pilihan string
	var k bool = false
	fmt.Print("Apakah anda yakin akan mengedit atau menghapus data ini? (y/n)")
	fmt.Scan(&pilihan)
	fmt.Println()

	if pilihan == "y" || pilihan == "Y" {
		k = true
	}
	return k
}


// calon fitur baru