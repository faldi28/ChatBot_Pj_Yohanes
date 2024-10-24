package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

// Struktur untuk menyimpan informasi client
type Client struct {
	Conn net.Conn
	Name string
}

// Daftar client yang terhubung
var (
	clients     = make(map[net.Conn]string) // Menyimpan koneksi client dan namanya
	clientMutex sync.Mutex                  // Mutex untuk menghindari kondisi balapan
)

func main() {
	// Memulai server pada port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Kesalahan saat memulai server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server berjalan di port 8080...")

	// Loop untuk menerima koneksi
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Kesalahan saat menerima koneksi:", err)
			continue
		}
		go handleConnection(conn) // Tangani setiap koneksi dalam goroutine baru
	}
}

// Fungsi untuk menangani koneksi client
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Menerima nama client
	clientName := receiveClientName(conn)
	if clientName == "" {
		return // Jika tidak ada nama, tutup koneksi
	}

	// Simpan client baru ke daftar
	addClient(conn, clientName)
	fmt.Printf("Client %s bergabung\n", clientName)

	// Broadcast pesan sambutan ke semua client
	broadcastMessage(fmt.Sprintf("%s telah bergabung", clientName), conn)

	// Baca pesan dari client dan broadcast ke semua client
	readMessages(conn, clientName)

	// Hapus client dari daftar jika terputus
	removeClient(conn, clientName)
}

// Fungsi untuk menerima nama client
func receiveClientName(conn net.Conn) string {
	scanner := bufio.NewScanner(conn)
	fmt.Fprintln(conn, "Masukkan nama Anda:")
	scanner.Scan()
	return scanner.Text()
}

// Fungsi untuk menambahkan client ke daftar
func addClient(conn net.Conn, name string) {
	clientMutex.Lock()
	defer clientMutex.Unlock()
	clients[conn] = name
}

// Fungsi untuk membaca pesan dari client
func readMessages(conn net.Conn, clientName string) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		pesan := scanner.Text()
		fullMessage := fmt.Sprintf("%s: %s", clientName, pesan)
		fmt.Println(fullMessage) // Log pesan di server
		broadcastMessage(fullMessage, conn)
	}
}

// Fungsi untuk menghapus client dari daftar
func removeClient(conn net.Conn, clientName string) {
	clientMutex.Lock()
	defer clientMutex.Unlock()
	delete(clients, conn)
	broadcastMessage(fmt.Sprintf("%s telah keluar", clientName), conn)
}

// Fungsi untuk broadcast pesan ke semua client kecuali pengirim
func broadcastMessage(message string, senderConn net.Conn) {
	clientMutex.Lock()
	defer clientMutex.Unlock()
	for clientConn := range clients {
		if clientConn != senderConn { // Jangan kirim ke pengirim
			fmt.Fprintln(clientConn, message)
		}
	}
}
