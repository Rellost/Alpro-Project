package main
import "fmt"
const NMAX int = 1000
type proyek struct {
	judul, klien string
	hari, bulan, tahun int
	bayaran int
	status string
}
var arrProyek [NMAX]proyek
var jumlahProyek int = 10

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

func cetakProyek() {
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Printf("   | %-30s | %-10s | %-10s | %-13s | %-10s\n", "Proyek", "Klien", "Deadline", "Bayaran", "Status")
	fmt.Println("---------------------------------------------------------------------------------------")
	for i := 0; i < jumlahProyek; i++ {
		bayaran := arrProyek[i].bayaran
		digitCount := 0
		temp := bayaran
		for temp != 0 {
			digitCount++
			temp /= 10
		}
		if digitCount > 10 {
			temp := bayaran
			factor := 1
			for j := 0; j < digitCount-9; j++ {
				factor *= 10
			}
			potong := temp / factor
			fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | Rp.%09d~ | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, potong, arrProyek[i].status)
		} else {
			fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | Rp.%-10d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, bayaran, arrProyek[i].status)
		}
	}
	fmt.Println("---------------------------------------------------------------------------------------")
}

func daftarProyek() {
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("|                                      DAFTAR PROYEK                                  |")
	cetakProyek()
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
	fmt.Println("4. Edit bayaran")
	fmt.Println("5. Edit status")
	fmt.Print(">>")
	var input int
	fmt.Scan(&input)
	for input != 1 && input != 2 && input != 3 && input != 4 && input != 5 {
		fmt.Print(">>")
		fmt.Scan(&input)
	}
	switch input {
		case 1: 
			fmt.Printf(">> (%s) -> ", arrProyek[n-1].judul)
			var judul string
			fmt.Scan(&judul)
			arrProyek[n-1].judul = ubahUnderscoreKeSpasi(judul)
		case 2:  
			fmt.Printf(">> (%s) -> ", arrProyek[n-1].klien)
			var klien string
			fmt.Scan(&klien)
			arrProyek[n-1].klien = ubahUnderscoreKeSpasi(klien)
		case 3: 
			fmt.Printf(">> (%02d %02d %04d) -> ", arrProyek[n-1].hari, arrProyek[n-1].bulan, arrProyek[n-1].tahun)
			fmt.Scan(&arrProyek[n-1].hari, &arrProyek[n-1].bulan, &arrProyek[n-1].tahun)
		case 4: 
			fmt.Printf(">> (Rp.%d) -> Rp.", arrProyek[n-1].bayaran)
			fmt.Scan(&arrProyek[n-1].bayaran)
		case 5: 
			fmt.Printf(">> (%s) -> ", arrProyek[n-1].status)
			fmt.Scan(&arrProyek[n-1].status)
	}
	daftarProyek()
}

func sortProyek() {
	fmt.Println("1. Judul")
	fmt.Println("2. Klien")
	fmt.Println("3. Deadline")
	fmt.Println("4. Bayaran")
	fmt.Println("5. Status")
	var input int
	fmt.Print(">>")
	fmt.Scan(&input)
	for input != 1 && input != 2 && input != 3 && input != 4 && input != 5 {
		fmt.Print(">>")
		fmt.Scan(&input)
	}
	switch input {
		case 1: 
			for pass := 1; pass < jumlahProyek; pass++ {
				idx := pass-1
				for i := pass; i < jumlahProyek; i++ {
					if len(arrProyek[idx].judul) > len(arrProyek[i].judul) {
						idx = i
					}
				}
				temp := arrProyek[pass-1]
				arrProyek[pass-1] = arrProyek[idx]
				arrProyek[idx] = temp
			}
		case 2:
			for pass := 1; pass < jumlahProyek; pass++ {
				idx := pass-1
				for i := pass; i < jumlahProyek; i++ {
					if arrProyek[idx].klien[0] > arrProyek[i].klien[0] {
						idx = i
					}
				}
				temp := arrProyek[pass-1]
				arrProyek[pass-1] = arrProyek[idx]
				arrProyek[idx] = temp
			}
		case 3:
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
		case 4:
			for pass := 1; pass < jumlahProyek; pass++ {
				idx := pass-1
				for i := pass; i < jumlahProyek; i++ {
					if arrProyek[idx].bayaran < arrProyek[i].bayaran {
						idx = i
					}
				}
				temp := arrProyek[pass-1] 
				arrProyek[pass-1] = arrProyek[idx]
				arrProyek[idx] = temp
			}
		case 5:
			for pass := 1; pass < jumlahProyek; pass++ {
				idx := pass-1
				for i := pass; i < jumlahProyek; i++ {
					if (arrProyek[idx].status == "done" && arrProyek[i].status == "pending") || (arrProyek[idx].status == "pending" && arrProyek[i].status == "working") || (arrProyek[idx].status == "done" && arrProyek[i].status == "working") {
						idx = i
					}
				}
				temp := arrProyek[pass-1]
				arrProyek[pass-1] = arrProyek[idx]
				arrProyek[idx] = temp
			}
	}
	cetakProyek()
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
	fmt.Println("4. Bayaran")
	fmt.Println("5. Status")
	fmt.Print(">>")
	var input int
	fmt.Scan(&input)
	for input != 1 && input != 2 && input != 3 && input != 4 && input != 5 {
		fmt.Print(">>")
		fmt.Scan(&input)
	}
	switch input {
		case 1: 
			fmt.Print(">> (Judul) ")
			var judul string
			fmt.Scan(&judul)
			fmt.Println("---------------------------------------------------------------------------------------")
			fmt.Printf("   | %-30s | %-10s | %-10s | %-13s | %-10s\n", "Proyek", "Klien", "Deadline", "Bayaran", "Status")
			fmt.Println("---------------------------------------------------------------------------------------")
			for i := 0; i < jumlahProyek; i++ {
				var tot1, tot2 [10]int
				a := 0
				b := 0
				for j := 0; j < len(arrProyek[i].judul); j++ {
					if arrProyek[i].judul[j] >= 'A' && arrProyek[i].judul[j] <= 'Z' {
						tot1[a] += int(arrProyek[i].judul[j]) + 32
					} else if arrProyek[i].judul[j] == ' ' {
						a++
					} else {
						tot1[a] += int(arrProyek[i].judul[j])
					}
				}
				for j := 0; j < len(judul); j++ {
					if judul[j] >= 'A' && judul[j] <= 'Z' {
						tot2[b] += int(judul[j]) + 32
					} else if judul[j] == ' ' {
						b++
					} else {
						tot2[b] += int(judul[j])
					}
				}
				for j := b; j >= 0; j-- {
					for k := a; k >= 0; k-- {
						if tot1[k] == tot2[j] {
							fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | Rp.%-10d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].bayaran, arrProyek[i].status)
						}
					}
				}
			}
		case 2:
			fmt.Print(">> (klien) ")
			var klien string
			fmt.Scan(&klien)
			fmt.Println("---------------------------------------------------------------------------------------")
			fmt.Printf("   | %-30s | %-10s | %-10s | %-13s | %-10s\n", "Proyek", "Klien", "Deadline", "Bayaran", "Status")
			fmt.Println("---------------------------------------------------------------------------------------")
			for i := 0; i < jumlahProyek; i++ {
				var tot1, tot2 [10]int
				a := 0
				b := 0
				for j := 0; j < len(arrProyek[i].klien); j++ {
					if arrProyek[i].klien[j] >= 'A' && arrProyek[i].klien[j] <= 'Z' {
						tot1[a] += int(arrProyek[i].klien[j]) + 32
					} else if arrProyek[i].klien[j] == ' ' {
						a++
					} else {
						tot1[a] += int(arrProyek[i].klien[j])
					}
				}
				for j := 0; j < len(klien); j++ {
					if klien[j] >= 'A' && klien[j] <= 'Z' {
						tot2[b] += int(klien[j]) + 32
					} else if klien[j] == ' ' {
						b++
					} else {
						tot2[b] += int(klien[j])
					}
				}
				for j := b; j >= 0; j-- {
					for k := a; k >= 0; k-- {
						if tot1[k] == tot2[j] {
							fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | Rp.%-10d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].bayaran, arrProyek[i].status)
						}
					}
				}
			}
		case 3:
			fmt.Print(">> (deadline) ")
			var tanggal int
			fmt.Scan(&tanggal)
			fmt.Println("---------------------------------------------------------------------------------------")
			fmt.Printf("   | %-30s | %-10s | %-10s | %-13s | %-10s\n", "Proyek", "Klien", "Deadline", "Bayaran", "Status")
			fmt.Println("---------------------------------------------------------------------------------------")
			for i := 0; i < jumlahProyek; i++ {
				if arrProyek[i].tahun == tanggal {
					fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | Rp.%-10d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].bayaran, arrProyek[i].status)
				} else if arrProyek[i].bulan == tanggal {
					fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | Rp.%-10d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].bayaran, arrProyek[i].status)
				} else if arrProyek[i].hari == tanggal {
					fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | Rp.%-10d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].bayaran, arrProyek[i].status)
				}
			}
		case 4:
			fmt.Print(">> (bayaran) Rp.")
			var bayaran int
			fmt.Scan(&bayaran)
			fmt.Println("---------------------------------------------------------------------------------------")
			fmt.Printf("   | %-30s | %-10s | %-10s | %-13s | %-10s\n", "Proyek", "Klien", "Deadline", "Bayaran", "Status")
			fmt.Println("---------------------------------------------------------------------------------------")
			for i := 0; i < jumlahProyek; i++ {
				if arrProyek[i].bayaran == bayaran {
					fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | Rp.%-10d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].bayaran, arrProyek[i].status)
				}
			}
		case 5:
			fmt.Print(">> (status) ")
			var status string
			fmt.Scan(&status)
			fmt.Println("---------------------------------------------------------------------------------------")
			fmt.Printf("   | %-30s | %-10s | %-10s | %-13s | %-10s\n", "Proyek", "Klien", "Deadline", "Bayaran", "Status")
			fmt.Println("---------------------------------------------------------------------------------------")
			for i := 0; i < jumlahProyek; i++ {
				if arrProyek[i].status == status {
					fmt.Printf(" %-2d| %-30s | %-10s | %02d-%02d-%04d | Rp.%-10d | %-10s\n", i+1, arrProyek[i].judul, arrProyek[i].klien, arrProyek[i].hari, arrProyek[i].bulan, arrProyek[i].tahun, arrProyek[i].bayaran, arrProyek[i].status)
				}
			}
	}
	fmt.Println("---------------------------------------------------------------------------------------")
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

func ubahUnderscoreKeSpasi(input string) string {
	var hasil string
	for i := 0; i < len(input); i++ {
		if input[i] == '_' {
			hasil += " "
		} else {
			hasil += string(input[i])
		}
	}
	return hasil
}

func tambahProyek() {
	fmt.Println("------------------------------------")
	fmt.Println("|         INPUT DATA PROYEK        |")
	fmt.Println("------------------------------------")
	fmt.Println("*WARN* Gunakan _ untuk spasi!")
	fmt.Print("Judul: ")
	var judul, klien string
	fmt.Scan(&judul)
	arrProyek[jumlahProyek].judul = ubahUnderscoreKeSpasi(judul)
	fmt.Print("Klien: ")
	fmt.Scan(&klien)
	arrProyek[jumlahProyek].klien = ubahUnderscoreKeSpasi(klien)
	fmt.Print("Deadline: (DD MM YYYY) ")
	fmt.Scan(&arrProyek[jumlahProyek].hari, &arrProyek[jumlahProyek].bulan, &arrProyek[jumlahProyek].tahun)
	fmt.Print("Bayaran: Rp.")
	fmt.Scan(&arrProyek[jumlahProyek].bayaran)
	fmt.Print("Status: (working/pending/done) ")
	fmt.Scan(&arrProyek[jumlahProyek].status)
	jumlahProyek++
	menu()
}

func hapusProyek() {
	fmt.Println("*WARN* Silahkan pilih nomor projek yang anda ingin hapus")
	cetakProyek()
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
	arrProyek[0] = proyek{judul: "Aplikasi Managemen Freelance", klien: "Andi", hari: 20, bulan: 12, tahun: 2025, bayaran: 1000000, status: "working"}
	arrProyek[1] = proyek{judul: "Vidio Presentasi", klien: "Vestia", hari: 25, bulan: 7, tahun: 2023, bayaran: 12000000, status: "done"}
	arrProyek[2] = proyek{judul: "Game Simple", klien: "Windah B", hari: 16, bulan: 7, tahun: 2025, bayaran: 156000000, status: "working"}
	arrProyek[3] = proyek{judul: "Website Penghasil Bitcoin", klien: "Hirata", hari: 20, bulan: 12, tahun: 2027, bayaran: 3000000, status: "pending"}
	arrProyek[4] = proyek{judul: "Aplikasi Absensi Sekolah", klien: "Sakura", hari: 10, bulan: 1, tahun: 2024, bayaran: 5000000, status: "done"}
	arrProyek[5] = proyek{judul: "Sistem Kasir Warung", klien: "Sumeru", hari: 5, bulan: 3, tahun: 2025, bayaran: 2500000, status: "working"}
	arrProyek[6] = proyek{judul: "Desain Logo Perusahaan", klien: "Lina A.S", hari: 22, bulan: 9, tahun: 2023, bayaran: 750000, status: "done"}
	arrProyek[7] = proyek{judul: "Web Portofolio Pribadi", klien: "Yudi Amir", hari: 30, bulan: 11, tahun: 2025, bayaran: 1800000, status: "pending"}
	arrProyek[8] = proyek{judul: "Aplikasi Booking Salon", klien: "Tari", hari: 8, bulan: 8, tahun: 2024, bayaran: 4000000, status: "working"}
	arrProyek[9] = proyek{judul: "Membuat Laporan", klien: "Rika Smth", hari: 18, bulan: 10, tahun: 2022, bayaran: 1500000, status: "done"}
	menu()
}
