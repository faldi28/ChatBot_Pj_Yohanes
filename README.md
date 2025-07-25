Proyek Aplikasi Chat dengan Go
Ini adalah proyek aplikasi chat sederhana yang dibangun menggunakan bahasa pemrograman Go. Proyek ini mengimplementasikan arsitektur client-server yang memungkinkan banyak pengguna untuk terhubung dan berkomunikasi secara real-time.

Dibuat oleh:
Nama: Yohanes J Palis

NIM: 213400009

Deskripsi Proyek

Aplikasi ini terdiri dari dua bagian utama:

Server: Berfungsi sebagai pusat yang menerima koneksi dari para client, menerima pesan, dan menyiarkannya ke semua client lain yang terhubung.

Client: Merupakan aplikasi yang dijalankan oleh pengguna. Client akan terhubung ke server, meminta nama pengguna, lalu dapat mengirim dan menerima pesan.

Fitur

Arsitektur Client-Server: Komunikasi terpusat melalui satu server.

Multi-Client: Server dapat menangani beberapa koneksi client secara bersamaan (konkuren) menggunakan goroutine.

Pesan Real-Time: Pesan yang dikirim oleh satu client akan langsung diterima oleh semua client lain.

Notifikasi Pengguna: Server akan memberi tahu semua client ketika ada pengguna baru yang bergabung atau pengguna yang keluar.

Input Nama Pengguna: Setiap client dapat mengatur namanya sendiri saat pertama kali terhubung.

Struktur Folder
/
├── client/
│   ├── client.go       # Kode sumber untuk aplikasi client
│   └── go.mod

└── server/
    ├── server.go       # Kode sumber untuk aplikasi server
    └── go.mod
    
Cara Menjalankan
Anda perlu menjalankan server terlebih dahulu, kemudian menjalankan client untuk terhubung ke server.

1. Menjalankan Server
Buka terminal Anda.
Masuk ke direktori server.

Bash

cd server

Jalankan file server.go.

Bash

go run server.go

Server sekarang akan berjalan dan mendengarkan koneksi pada port 8080.

Server berjalan di port 8080...

3. Menjalankan Client
   
Buka terminal baru (biarkan terminal server tetap berjalan).

Masuk ke direktori client.

Bash

cd client

Jalankan file client.go.

Bash

go run client.go

Aplikasi client akan meminta Anda untuk memasukkan nama.

Masukkan nama Anda:

Setelah memasukkan nama, Anda akan terhubung ke server dan dapat mulai mengirim dan menerima pesan.

Untuk keluar dari chat, ketik exit lalu tekan Enter.

Anda dapat menjalankan beberapa instance dari client.go di terminal yang berbeda untuk menyimulasikan obrolan dengan banyak pengguna.
