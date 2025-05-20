package main
import (
	"fmt"
	"strings"
)
const NMAX int = 1000
type proyek struct {
	judul, klien string
	hari, bulan, tahun int
	status string
}
var arrProyek [NMAX]proyek
var jumlahProyek int = 4
type akun struct {
	nama, username, password string
}
var arrAkun [NMAX]akun
var jumlahAkun int = 0

func menuLogin() {
	fmt.Println("---------------------------")
	fmt.Println("|         LOGIN           |")
	fmt.Println("---------------------------")
	fmt.Println("|1. Login                 |")
	fmt.Println("|2. Daftar                |")
	fmt.Println("|3. Keluar                |")
	fmt.Println("---------------------------")
	fmt.Print(">>")
	var input int
	fmt.Scan(&input)
	switch input {
		case 1: login()
		case 2: daftar()
		case 3: return
	}
}

func login() {
	var indexAkun int
	fmt.Println("Silahkan masukkan username dan password")
	fmt.Print("Username: ")
	var username string
	fmt.Scan(&username)
	fmt.Print("Password: ")
	var password string
	fmt.Scan(&password)
	indexAkun = findIndexAkun(username)
	if indexAkun != -1 && username == arrAkun[indexAkun].username && password == arrAkun[indexAkun].password {
		menu()
	} else {
		fmt.Println("*WARN* Username atau password salah!")
		menuLogin()
	}
}

func daftar() {
	fmt.Println("Silahkan masukkan data diri anda")
	fmt.Print("Nama: ")
	fmt.Scan(&arrAkun[jumlahAkun].nama)
	fmt.Print("Username: ")
	fmt.Scan(&arrAkun[jumlahAkun].username)
	fmt.Print("Password: ")
	fmt.Scan(&arrAkun[jumlahAkun].password)
	jumlahAkun++
	menuLogin()
}

func findIndexAkun(username string) int {
	for i := 0; i < jumlahAkun; i++ {
		if arrAkun[i].username == username {
			return i
		}
	}
	return -1
}

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
	for input != 1 && input != 2 && input != 3 && input != 4 {
		fmt.Print(">>")
		fmt.Scan(&input)
	}
	switch input {
		case 1: daftarProyek()
		case 2: tambahProyek()
		case 3: hapusProyek()
		case 4: return
	}
}

func daftarProyek() {
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("|                             DAFTAR PROYEK                           |")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Printf("   | %-30s | %-10s | %-10s | %-10s\n", "Proyek", "Klien", "Deadline", "Status")
	fmt.Println("-----------------------------------------------------------------------")
	for i := 0; i < jumlahProyek; i++ {
		fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
	}
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("e:edit  s:sort  f:find  x:exit") 
	fmt.Print(">>")
	var input string
	fmt.Scan(&input)
	for input != "e" && input != "s" && input != "f" && input != "x" {
		fmt.Print(">>")
		fmt.Scan(&input)
	}
	switch input {
		case "e": editProyek()
		case "s": sortProyek()
		case "f": findProyek()
		case "x": menu()
	}
}

func editProyek() {
	fmt.Println("Silahkan pilih nomor projek yang anda ingin edit")
	fmt.Print(">>")
	var n int
	fmt.Scan(&n)
	fmt.Println("1. Edit judul")
	fmt.Println("2. Edit klien")
	fmt.Println("3. Edit deadline")
	fmt.Println("4. Edit status")
	fmt.Print(">>")
	var input int
	fmt.Scan(&input)
	for input != 1 && input != 2 && input != 3 && input != 4 {
		fmt.Print(">>")
		fmt.Scan(&input)
	}
	switch input {
		case 1: 
			fmt.Print(">>Judul: ")
			fmt.Scan(&arrProyek[n-1].judul)
		case 2: 
			fmt.Print(">>Klien: ")
			fmt.Scan(&arrProyek[n-1].klien)
		case 3: 
			fmt.Print(">>Deadline: (DD MM YYYY) ")
			fmt.Scan(&arrProyek[n-1].hari, &arrProyek[n-1].bulan, &arrProyek[n-1].tahun)
		case 4: 
			fmt.Print(">>Status: (working/pending/done) ")
			fmt.Scan(&arrProyek[n-1].status)
	}
	daftarProyek()
}

func sortProyek() {
	fmt.Println("1. Judul")
	fmt.Println("2. Klien")
	fmt.Println("3. Deadline")
	fmt.Println("4. Status")
	var input int
	fmt.Print(">>")
	fmt.Scan(&input)
	for input != 1 && input != 2 && input != 3 && input != 4 {
		fmt.Print(">>")
		fmt.Scan(&input)
	}
	switch input {
		case 1: 
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Printf("   | %-30s | %-10s | %-10s | %-10s\n", "Proyek", "Klien", "Deadline", "Status")
			fmt.Println("-----------------------------------------------------------------------")
			for pass := 1; pass < jumlahProyek; pass++ {
				idx := pass-1
				for i := pass; i < jumlahProyek; i++ {
					if strings.Compare(arrProyek[idx].judul, arrProyek[i].judul) > 0 {
						idx = i
					}
				}
				temp := arrProyek[pass-1]
				arrProyek[pass-1] = arrProyek[idx]
				arrProyek[idx] = temp
			}
			for i := 0; i < jumlahProyek; i++ {
				fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
			}
		case 2:
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Printf("   | %-30s | %-10s | %-10s | %-10s\n", "Proyek", "Klien", "Deadline", "Status")
			fmt.Println("-----------------------------------------------------------------------")
			for pass := 1; pass < jumlahProyek; pass++ {
				idx := pass-1
				for i := pass; i < jumlahProyek; i++ {
					if strings.Compare(arrProyek[idx].klien, arrProyek[i].klien) > 0 {
						idx = i
					}
				}
				temp := arrProyek[pass-1]
				arrProyek[pass-1] = arrProyek[idx]
				arrProyek[idx] = temp
			}
			for i := 0; i < jumlahProyek; i++ {
				fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
			}
		case 3:
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Printf("   | %-30s | %-10s | %-10s | %-10s\n", "Proyek", "Klien", "Deadline", "Status")
			fmt.Println("-----------------------------------------------------------------------")
			for pass := 1; pass < jumlahProyek; pass++ {
				idx := pass-1
				for i := pass; i < jumlahProyek; i++ {
					if arrProyek[i].tahun < arrProyek[idx].tahun || (arrProyek[i].tahun == arrProyek[idx].tahun && arrProyek[i].bulan < arrProyek[idx].bulan) || (arrProyek[i].tahun == arrProyek[idx].tahun && arrProyek[i].bulan == arrProyek[idx].bulan && arrProyek[i].hari < arrProyek[idx].hari) {
						idx = i
					}
				}
				temp := arrProyek[pass-1]
				arrProyek[pass-1] = arrProyek[idx]
				arrProyek[idx] = temp
			}
			for i := 0; i < jumlahProyek; i++ {
				fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
			}
		case 4:
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Printf("   | %-30s | %-10s | %-10s | %-10s\n", "Proyek", "Klien", "Deadline", "Status")
			fmt.Println("-----------------------------------------------------------------------")
			for pass := 1; pass < jumlahProyek; pass++ {
				idx := pass-1
				for i := pass; i < jumlahProyek; i++ {
					if arrProyek[i].tahun < arrProyek[idx].tahun || (arrProyek[i].tahun == arrProyek[idx].tahun && arrProyek[i].bulan < arrProyek[idx].bulan) || (arrProyek[i].tahun == arrProyek[idx].tahun && arrProyek[i].bulan == arrProyek[idx].bulan && arrProyek[i].hari < arrProyek[idx].hari) {
						idx = i
					}
				}
				temp := arrProyek[pass-1]
				arrProyek[pass-1] = arrProyek[idx]
				arrProyek[idx] = temp
			}
			for i := 0; i < jumlahProyek; i++ {
				fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
			}
	}
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("e:edit  s:sort  f:find  x:exit") 
	fmt.Print(">>")
	var s string
	fmt.Scan(&s)
	for s != "e" && s != "s" && s != "f" && s != "x" {
		fmt.Print(">>")
		fmt.Scan(&s)
	}
	switch s {
		case "e": editProyek()
		case "s": sortProyek()
		case "f": findProyek()
		case "x": menu()
	}
}

func findProyek() {
	fmt.Println("1. Judul")
	fmt.Println("2. Klien")
	fmt.Println("3. Deadline")
	fmt.Println("4. Status")
	fmt.Print(">>")
	var input int
	fmt.Scan(&input)
	for input != 1 && input != 2 && input != 3 && input != 4 {
		fmt.Print(">>")
		fmt.Scan(&input)
	}
	switch input {
		case 1: 
			fmt.Print(">> (judul) ")
			var judul string
			fmt.Scan(&judul)
			fmt.Printf("   | %-30s | %-10s | %-10s | %-10s\n", "Proyek", "Klien", "Deadline", "Status")
			fmt.Println("-----------------------------------------------------------------------")
			for i := 0; i < jumlahProyek; i++ {
				if strings.Contains(strings.ToLower(arrProyek[i].judul), strings.ToLower(judul)) {
					fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
				}
			}
		case 2:
			fmt.Print(">> (klien) ")
			var klien string
			fmt.Scan(&klien)
			fmt.Printf("   | %-30s | %-10s | %-10s | %-10s\n", "Proyek", "Klien", "Deadline", "Status")
			fmt.Println("-----------------------------------------------------------------------")
			for i := 0; i < jumlahProyek; i++ {
				if arrProyek[i].klien == klien {
					fmt.Printf("%d. %s %s %d-%d-%d %s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
				}
			}
		case 3:
			fmt.Print(">> (deadline) ")
			var tanggal int
			fmt.Scan(&tanggal)
			fmt.Printf("   | %-30s | %-10s | %-10s | %-10s\n", "Proyek", "Klien", "Deadline", "Status")
			fmt.Println("-----------------------------------------------------------------------")
			for i := 0; i < jumlahProyek; i++ {
				if arrProyek[i].tahun == tanggal {
					fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
				} else if arrProyek[i].bulan == tanggal {
					fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
				} else if arrProyek[i].hari == tanggal {
					fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
				}
			}
		case 4:
			fmt.Print(">> (status) ")
			var status string
			fmt.Scan(&status)
			fmt.Printf("   | %-30s | %-10s | %-10s | %-10s\n", "Proyek", "Klien", "Deadline", "Status")
			fmt.Println("-----------------------------------------------------------------------")
			for i := 0; i < jumlahProyek; i++ {
				if arrProyek[i].status == status {
					fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
				}
			}
	}
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("e:edit  s:sort  f:find  x:exit") 
	fmt.Print(">>")
	var s string
	fmt.Scan(&s)
	for s != "e" && s != "s" && s != "f" && s != "x" {
		fmt.Print(">>")
		fmt.Scan(&s)
	}
	switch s {
		case "e": editProyek()
		case "s": sortProyek()
		case "f": findProyek()
		case "x": menu()
	}
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
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Printf("   | %-30s | %-10s | %-10s | %-10s\n", "Proyek", "Klien", "Deadline", "Status")
	fmt.Println("-----------------------------------------------------------------------")
	for i := 0; i < jumlahProyek; i++ {
		fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].status)
	}
	fmt.Println("-----------------------------------------------------------------------")
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
	arrProyek[0] = proyek{judul: "Aplikasi Managemen Freelance", klien: "Andi", hari: 20, bulan: 12, tahun: 2025, status: "working"}
	arrProyek[1] = proyek{judul: "Membuat Vidio Perkenalan", klien: "Vestia", hari: 25, bulan: 7, tahun: 2023, status: "done"}
	arrProyek[2] = proyek{judul: "Membuat Game Simple", klien: "Windah", hari: 16, bulan: 7, tahun: 2025, status: "working"}
	arrProyek[3] = proyek{judul: "Membuat Website Bitcoin", klien: "Hirata", hari: 20, bulan: 12, tahun: 2027, status: "pending"}
	menuLogin()
}
