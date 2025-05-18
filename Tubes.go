package main

import "fmt"

type akun struct{
	username string
	password string

}
type proyek struct {
    namaProyek   string
    namaKlien    string
    deadline     string
    status       string
}

const NMAX int = 100
type arrAkun [NMAX] akun
type arrProyek [NMAX] proyek

func login(){
	fmt.Println("=========================================================")
	fmt.Println("Selamat Datang di Aplikasi Freelace Managemen & Tracking")
	fmt.Println("=========================================================")
	fmt.Println("1.Login")
	fmt.Println("2.Daftar akun")
	fmt.Println("3.Exit")
}
func daftarAkun(A *akun, n int){


}

func menu(){
	fmt.Println("=========================================================")
	fmt.Println("|                                                       |")
	fmt.Println("|      Aplikasi Manajemen dan Tracking Freelance        |")
	fmt.Println("|                                                       |")
	fmt.Println("=========================================================")
	fmt.Println("|1.Daftar Proyek                                        |")
	fmt.Println("|2.Tambah Proyek                                        |")
	fmt.Println("|3.Hapus Proyek                                         |")
	fmt.Println("|4.Ubah Proyek                                          |")
	fmt.Println("|5.Status Proyek                                        |")
	fmt.Println("|6.Cari Proyek                                          |")
	fmt.Println("|7.Laporan Proyek                                       |")
	fmt.Println("|8.Exit                                                 |")
	fmt.Println("=========================================================")
}

func daftarProyek(A arrProyek, n int){
	var i int
	fmt.Println("=========================================================")
	fmt.Println("|                   Daftar Proyek                       |")
	fmt.Println("=========================================================")
	for i = 0; i < n; i++ {
	fmt.Printf("|Proyek ke-%-42d   |\n", i+1)
	fmt.Printf("|Nama Proyek :%-42s|\n", A[i].namaProyek)
	fmt.Printf("|Nama Klien  :%-42s|\n",  A[i].namaKlien)
	fmt.Printf("|Deadline    :%-42s|\n", A[i].deadline)
	fmt.Printf("|Status      :%-42s|\n",  A[i].status)
	}
	fmt.Println("=========================================================")
}

func tambahProyek(A *arrProyek, n *int) {
	var jumlah, i int
	fmt.Println("=========================================================")
	fmt.Println("|                                                       |")
	fmt.Println("|      Aplikasi Manajemen dan Tracking Freelance        |")
	fmt.Println("|                                                       |")
	fmt.Println("=========================================================")
	fmt.Println("|                                                       |")
	fmt.Println("|       Selamat datang di menu tambah proyek            |")
	fmt.Println("|                                                       |")
	fmt.Println("|  Silakan masukkan informasi proyek yang ingin Anda    |")
	fmt.Println("|  tambahkan ke dalam daftar                            |")	
	fmt.Println("|                                                       |")
	fmt.Println("|                                                       |")
	fmt.Println("|                                                       |")
	fmt.Println("=========================================================")
	fmt.Print(">Jumlah proyek yang ingin ditambahkan: ")
	fmt.Scan(&jumlah)

	for i = 0; i < jumlah; i++ {
		fmt.Printf("Proyek ke-%d:\n", *n+1)
		fmt.Print("Nama Proyek: ")
		fmt.Scan(&A[*n].namaProyek)
		fmt.Print("Nama Klien: ")
		fmt.Scan(&A[*n].namaKlien)
		fmt.Print("Deadline: ")
		fmt.Scan(&A[*n].deadline)
		fmt.Print("Status: ")
		fmt.Scan(&A[*n].status)

		*n++
	}
	fmt.Println("Proyek anda berhasil ditambahkan")
}

func hapusProyek(A *arrProyek, n *int){
	var nama string
	var i, b int
	fmt.Println("=========================================================")
	fmt.Println("|                                                       |")
	fmt.Println("|      Aplikasi Manajemen dan Tracking Freelance        |")
	fmt.Println("|                                                       |")
	fmt.Println("=========================================================")
	fmt.Println("|                                                       |")
	fmt.Println("|       Selamat Datang di Menu Hapus Proyek             |")
	fmt.Println("|                                                       |")
	fmt.Println("|  Silakan masukkan nama proyek yang ingin Anda         |")
	fmt.Println("|  anda hapus dalam daftar                              |")	
	fmt.Println("|                                                       |")
	fmt.Println("|                                                       |")
	fmt.Println("|                                                       |")
	fmt.Println("=========================================================")
	fmt.Print(">Nama proyek yang ingin dihapus: ")
	fmt.Scan(&nama)
	for i = 0; i < *n; i++{
		if nama == A[i].namaProyek{
			for b = i; b < *n-1; b++{
				A[b] = A[b+1]
			}
			*n--
		}
	}

}

func ubahProyek(){

}
// masi salah
func statusProyek(A *arrProyek, n *int){
	var pilih int
	fmt.Println("===================================")
	fmt.Println("Apakah anda ingin memperbarui status proyek :")
	fmt.Println("1.Ya")
	fmt.Println("2.Tidak")
	fmt.Print("Silahkan pilih opsi yang tersedia: ")
	fmt.Scan(&pilih)
	switch pilih{
	case 1 :
		fmt.Print("Status proyek :")
		fmt.Scan(&A[*n].status)
		
	case 2 :
		return
	}
}

func cariProyek(){


}

func laporanProyek(){

}

func main(){
	var a arrProyek
	var n int
	var pilih int
	for{
	menu()
	fmt.Print(">Silahkan pilih opsi yang tersedia: ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1 :
		daftarProyek(a, n)
	case 2: 
		tambahProyek(&a, &n)
	case 3 :
		hapusProyek(&a, &n)
	case 4 :
		ubahProyek()
	case 5 :
		statusProyek(&a, &n)
	case 6 :
		cariProyek()
	case 7 :
		laporanProyek()
	case 8 :
		fmt.Println("Anda keluar dari menu")

	}
}
}
