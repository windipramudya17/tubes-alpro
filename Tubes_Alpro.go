package main

import "fmt"

const NMAX = 100

type Tim struct {
	Nama  string
	Peran string
}

type Startup struct {
	Nama         string
	BidangUsaha  string
	TahunBerdiri int
	TotalDana    float64
	JumlahTim    int
	Tim          [NMAX]Tim
}

var daftarStartup [NMAX]Startup
var jumlahStartup int

func main() {
	var pilihan int

	for {
		menu()
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahStartup()
		} else if pilihan == 2 {
			ubahStartup()
		} else if pilihan == 3 {
			hapusStartup()
		} else if pilihan == 4 {
			cariStartup()
		} else if pilihan == 5 {
			urutkanStartup()
		} else if pilihan == 6 {
			laporanBidangUsaha()
		} else if pilihan == 7 {
			cariPendanaan()
		} else if pilihan == 8 {
			fmt.Println("Terima Kasih Telah Menggunakan Aplikasi")
			return
		} else {
			fmt.Println("Pilihan Tidak Valid")
			return
		}
	}
}

func menu() {
	fmt.Println("\n=== APLIKASI MANAJEMEN STARTUP SEDERHANA ===")
	fmt.Println("1. Tambah Data Startup")
	fmt.Println("2. Ubah Data Startup")
	fmt.Println("3. Hapus Data Startup")
	fmt.Println("4. Cari Startup")
	fmt.Println("5. Urutkan Data Startup")
	fmt.Println("6. Laporan per Bidang Usaha")
	fmt.Println("7. Cari Pendanaan Tertinggi / Terendah")
	fmt.Println("8. Keluar")
}

func validasiNama(prompt string, maxLen int) string {
	var input string
	var valid bool = false
	for !valid {
		fmt.Print(prompt)
		fmt.Scan(&input)
		valid = len(input) <= maxLen
		if !valid {
			fmt.Printf("Input terlalu panjang (maks %d karakter)\n", maxLen)
		}
	}
	return input
}

func validasiJumlah(prompt string, min, max int) int {
	var input int
	var valid bool = false
	for !valid {
		fmt.Print(prompt)
		fmt.Scan(&input)
		valid = input >= min && input <= max
		if !valid {
			fmt.Printf("Input harus antara %d dan %d\n", min, max)
		}
	}
	return input
}

func validasiDana(prompt string, min float64) float64 {
	var input float64
	var valid bool = false
	for !valid {
		fmt.Print(prompt)
		fmt.Scan(&input)
		valid = input >= min
		if !valid {
			fmt.Printf("Input tidak boleh kurang dari %.2f\n", min)
		}
	}
	return input
}


func tampilkanPilihanUrutan() {
	fmt.Println("Urutan: 1. Ascending 2. Descending")
}

func tambahStartup() {
	if jumlahStartup >= NMAX {
		fmt.Println("Data penuh")
		return
	}
	var s Startup

	s.Nama = validasiNama("Nama Startup: ", 20)

	fmt.Print("Bidang Usaha: ")
	fmt.Scan(&s.BidangUsaha)

	s.TahunBerdiri = validasiJumlah("Tahun Berdiri: ", 2000, 2025)
	s.TotalDana = validasiDana("Total Pendanaan: ", 1000000)
	s.JumlahTim = validasiJumlah("Jumlah Anggota Tim: ", 0, NMAX)

	if s.JumlahTim > NMAX {
		fmt.Println("Jumlah anggota tim melebihi batas maksimum")
		return
	}

	for i := 0; i < s.JumlahTim; i++ {
		fmt.Printf("- Nama Anggota #%d: ", i+1)
		fmt.Scan(&s.Tim[i].Nama)
		fmt.Printf("  Peran Anggota #%d: ", i+1)
		fmt.Scan(&s.Tim[i].Peran)
	}

	daftarStartup[jumlahStartup] = s
	jumlahStartup++
	fmt.Println("Data startup berhasil ditambahkan")
}

func ubahStartup() {
	var nama string
	fmt.Print("Nama Startup yang ingin diubah: ")
	fmt.Scan(&nama)
	idx := sequentialSearch(nama)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}

	fmt.Println("Data ditemukan. Masukkan data baru")
	var s Startup
	fmt.Print("Nama Startup: ")
	fmt.Scan(&s.Nama)
	fmt.Print("Bidang Usaha: ")
	fmt.Scan(&s.BidangUsaha)
	fmt.Print("Tahun Berdiri: ")
	fmt.Scan(&s.TahunBerdiri)
	fmt.Print("Total Pendanaan: ")
	fmt.Scan(&s.TotalDana)
	fmt.Print("Jumlah Anggota Tim: ")
	fmt.Scan(&s.JumlahTim)

	if s.JumlahTim > NMAX {
		fmt.Println("Jumlah anggota tim melebihi batas maksimum.")
		return
	}

	for i := 0; i < s.JumlahTim; i++ {
		fmt.Printf("- Nama Anggota #%d: ", i+1)
		fmt.Scan(&s.Tim[i].Nama)
		fmt.Printf("  Peran Anggota #%d: ", i+1)
		fmt.Scan(&s.Tim[i].Peran)
	}

	daftarStartup[idx] = s
	fmt.Println("Data startup berhasil diubah")
}

func hapusStartup() {
	var nama string
	fmt.Print("Nama Startup yang ingin dihapus: ")
	fmt.Scan(&nama)
	idx := sequentialSearch(nama)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}
	for i := idx; i < jumlahStartup-1; i++ {
		daftarStartup[i] = daftarStartup[i+1]
	}
	jumlahStartup--
	fmt.Println("Data berhasil dihapus")
}

func cariStartup() {
	var nama string
	var metode int
	fmt.Print("Nama Startup yang dicari: ")
	fmt.Scan(&nama)
	fmt.Println("Metode Pencarian: 1. Sequential 2. Binary")
	fmt.Scan(&metode)

	var idx int = -1
	if metode == 1 {
		idx = sequentialSearch(nama)
	} else {
		insertionSort(1, true) 
		idx = binarySearch(nama)
	}

	if idx != -1 {
		fmt.Println("Startup ditemukan:")
		tampilkanStartup(daftarStartup[idx])
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func sequentialSearch(nama string) int {
	for i := 0; i < jumlahStartup; i++ {
		if daftarStartup[i].Nama == nama {
			return i
		}
	}
	return -1
}

func binarySearch(nama string) int {
	low := 0
	high := jumlahStartup - 1
	for low <= high {
		mid := (low + high) / 2
		if daftarStartup[mid].Nama == nama {
			return mid
		} else if daftarStartup[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func urutkanStartup() {
	var metode, kriteria, urutan int
	fmt.Println("Metode Urut: 1. Selection 2. Insertion")
	fmt.Scan(&metode)
	fmt.Println("Kriteria: 1. Nama 2. Tahun Berdiri 3. Total Dana")
	fmt.Scan(&kriteria)
	fmt.Println("Urutan: 1. Ascending 2. Descending")
	fmt.Scan(&urutan)

	asc := urutan == 1
	if metode == 1 {
		selectionSort(kriteria, asc)
	} else {
		insertionSort(kriteria, asc)
	}

	fmt.Println("Data berhasil diurutkan")
}

func selectionSort(kriteria int, asc bool) {
	for i := 0; i < jumlahStartup-1; i++ {
		idx := i
		for j := i + 1; j < jumlahStartup; j++ {
			if banding(daftarStartup[j], daftarStartup[idx], kriteria, asc) {
				idx = j
			}
		}
		daftarStartup[i], daftarStartup[idx] = daftarStartup[idx], daftarStartup[i]
	}
}

func insertionSort(kriteria int, asc bool) {
	for i := 1; i < jumlahStartup; i++ {
		temp := daftarStartup[i]
		j := i - 1
		for j >= 0 && banding(temp, daftarStartup[j], kriteria, asc) {
			daftarStartup[j+1] = daftarStartup[j]
			j--
		}
		daftarStartup[j+1] = temp
	}
}

func banding(a, b Startup, kriteria int, asc bool) bool {
	switch kriteria {
	case 1:
		if asc {
			return a.Nama < b.Nama
		}
		return a.Nama > b.Nama
	case 2:
		if asc {
			return a.TahunBerdiri < b.TahunBerdiri
		}
		return a.TahunBerdiri > b.TahunBerdiri
	case 3:
		if asc {
			return a.TotalDana < b.TotalDana
		}
		return a.TotalDana > b.TotalDana
	default:
		return false
	}
}

func laporanBidangUsaha() {
	if jumlahStartup == 0 {
		fmt.Println("Tidak ada data")
		return
	}
	fmt.Println("Laporan Startup per Bidang Usaha:")
	for i := 0; i < jumlahStartup; i++ {
		bidang := daftarStartup[i].BidangUsaha
		jumlah := 1
		for j := 0; j < i; j++ {
			if daftarStartup[j].BidangUsaha == bidang {
				jumlah = 0
			}
		}
		if jumlah > 0 {
			hitung := 0
			for k := 0; k < jumlahStartup; k++ {
				if daftarStartup[k].BidangUsaha == bidang {
					hitung++
				}
			}
			fmt.Printf("%s: %d startup\n", bidang, hitung)
		}
	}
}

func tampilkanStartup(s Startup) {
	fmt.Println("====================================")
	fmt.Println("Nama Startup     :", s.Nama)
	fmt.Println("Bidang Usaha     :", s.BidangUsaha)
	fmt.Println("Tahun Berdiri    :", s.TahunBerdiri)
	fmt.Printf("Total Pendanaan  : %.2f\n", s.TotalDana)
	fmt.Println("Jumlah Tim       :", s.JumlahTim)
	fmt.Println("Anggota Tim:")
	for i := 0; i < s.JumlahTim; i++ {
		fmt.Printf("  - %s (%s)\n", s.Tim[i].Nama, s.Tim[i].Peran)
	}
	fmt.Println("====================================")
}

func cariPendanaan() {
	if jumlahStartup == 0 {
		fmt.Println("Tidak ada data")
		return
	}

	fmt.Println("Cari: 1. Pendanaan Tertinggi 2. Pendanaan Terendah")
	var pilihan int
	fmt.Scan(&pilihan)

	idx := 0
	for i := 1; i < jumlahStartup; i++ {
		if (pilihan == 1 && daftarStartup[i].TotalDana > daftarStartup[idx].TotalDana) {
			idx = i
		} else if (pilihan == 2 && daftarStartup[i].TotalDana < daftarStartup[idx].TotalDana) {
			idx = i
		}
	}

	if pilihan == 1 {
		fmt.Println("Startup dengan pendanaan tertinggi:")
	} else {
		fmt.Println("Startup dengan pendanaan terendah:")
	}
	tampilkanStartup(daftarStartup[idx])
}
