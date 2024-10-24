package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// Fungsi untuk menghubungkan ke server
func hubungkanKeServer(alamat string) (net.Conn, error) {
	koneksi, err := net.Dial("tcp", alamat)
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke server: %v", err)
	}
	return koneksi, nil
}

// Fungsi untuk meminta nama pengguna
func mintaNamaPengguna() string {
	fmt.Print("Masukkan nama Anda: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// Fungsi untuk mendengarkan pesan yang diterima dari server
func dengarkanPesan(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		pesan := scanner.Text()
		fmt.Println("Pesan dari server:", pesan)
	}
}

// Fungsi untuk mengirim pesan ke server
func kirimPesan(conn net.Conn, nama string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pesan := scanner.Text()
		
		// Cek jika pengguna ingin keluar
		if pesan == "exit" {
			fmt.Fprintf(conn, "%s telah keluar.\n", nama)
			fmt.Println("Anda telah keluar dari server.")
			conn.Close()
			os.Exit(0)
		}
		
		fmt.Fprintf(conn, "%s: %s\n", nama, pesan)
	}
}

func main() {
	// Menghubungkan ke server
	conn, err := hubungkanKeServer("localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Mendapatkan nama pengguna
	nama := mintaNamaPengguna()

	fmt.Println("Terhubung ke server. Anda sekarang dapat mengirim pesan:")

	// Goroutine untuk mendengarkan pesan masuk
	go dengarkanPesan(conn)

	// Mengirim pesan ke server
	kirimPesan(conn, nama)
}
