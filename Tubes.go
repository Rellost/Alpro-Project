package main
import "fmt"
const NMAX int = 1000
type proyek struct {
	judul, klien, status string
	hari, bulan, tahun int
}
var arrProyek [NMAX]proyek
var jumlahProyek int = 4

func menu(){
	fmt.Println("---------------------------")
	fmt.Println("|        MAIN MENU        |")
	fmt.Println("---------------------------")
	fmt.Println("|1. Lihat Daftar Proyek   |")
	fmt.Println("|2. Tambah Proyek         |")
	fmt.Println("|3. Hapus Proyek          |")
	fmt.Println("|4. Keluar                |")
	fmt.Println("---------------------------")
	fmt.Print(">>")
	var input int
	fmt.Scan(&input)
	switch input {
		case 1: daftarProyek()
		case 2: tambahProyek()
		case 3: hapusProyek()
		case 4: return
	}
}

func daftarProyek() {
	fmt.Println("--------------------------------")
	fmt.Println("|         DAFTAR PROYEK        |")
	fmt.Println("--------------------------------")
	if jumlahProyek == 0 {
		fmt.Println("*WARN* Anda belum memiliki proyek!")
	}
	for i := 0; i < jumlahProyek; i++ {
		fmt.Printf("%d. %s %s %d-%d-%d %s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
	}
	fmt.Println("------------------------------")
	fmt.Println("e:edit  s:sort  f:find  x:exit") 
	fmt.Print(">>")
	var input string
	fmt.Scan(&input)
	switch input {
		case "e": editProyek()
		case "s": sortProyek()
		case "f": findProyek()
		case "x": menu()
	}
}

func editProyek() {
	fmt.Println("Silahkan pilih nomor projek yang anda ingin edit")
	for i := 0; i < jumlahProyek; i++ {
		fmt.Printf("%d. %s %s %d-%d-%d %s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
	}
	fmt.Print(">>")
	var n int
	fmt.Scan(&n)
	fmt.Println("1. edit judul")
	fmt.Println("2. edit klien")
	fmt.Println("3. edit deadline")
	fmt.Println("4. edit status")
	fmt.Print(">>")
	var input int
	fmt.Scan(&input)
	switch input {
		case 1: 
			fmt.Print("Judul: ")
			fmt.Scan(&arrProyek[n-1].judul)
		case 2: 
			fmt.Print("Klien: ")
			fmt.Scan(&arrProyek[n-1].klien)
		case 3: 
			fmt.Print("Deadline: (DD MM YYYY) ")
			fmt.Scan(&arrProyek[n-1].hari, &arrProyek[n-1].bulan, &arrProyek[n-1].tahun)
		case 4: 
			fmt.Print("Status: (working/pending/done) ")
			fmt.Scan(&arrProyek[n-1].status)
	}
	daftarProyek()
}

func sortProyek() {
	fmt.Print("1. Judul")
	fmt.Print("2. Klien")
	fmt.Print("3. Deadline")
	fmt.Print("4. Status")
	var input int
	fmt.Scan(&input)
	switch input {
		case 1: 
	}
}

func findProyek() {
	
}

func tambahProyek() {
	fmt.Println("------------------------------------")
	fmt.Println("|         INPUT DATA PROYEK        |")
	fmt.Println("------------------------------------")
	fmt.Print("Judul: ")
	fmt.Scan(&arrProyek[jumlahProyek].judul)
	fmt.Print("Klien: ")
	fmt.Scan(&arrProyek[jumlahProyek].klien)
	fmt.Print("Deadline: (DD MM YYYY) ")
	fmt.Scan(&arrProyek[jumlahProyek].hari, &arrProyek[jumlahProyek].bulan, &arrProyek[jumlahProyek].tahun)
	fmt.Print("Status: (working/pending/done) ")
	fmt.Scan(&arrProyek[jumlahProyek].status)
	jumlahProyek++
	menu()
}

func hapusProyek() {
	fmt.Println("*WARN* Silahkan pilih nomor projek yang anda ingin hapus")
	for i := 0; i < jumlahProyek; i++ {
		fmt.Printf("%d. %s %s %d-%d-%d %s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
	}
	fmt.Print(">>")
	var n int
	fmt.Scan(&n)
	for i := n; i < jumlahProyek; i++ {
		arrProyek[i-1] = arrProyek[i]
	}
	jumlahProyek--
	menu()
}

func main() {
	arrProyek[0].judul = "Pembangunan Taman Baca"
	arrProyek[0].klien = "Andi"
	arrProyek[0].hari = 20
	arrProyek[0].bulan = 12 
	arrProyek[0].tahun = 2022
	arrProyek[0].status = "working"
	arrProyek[1].judul = "Pembangunan Website"
	arrProyek[1].klien = "Galang"
	arrProyek[1].hari = 15
	arrProyek[1].bulan = 6 
	arrProyek[1].tahun = 2025
	arrProyek[1].status = "pending"
	menu()
}
