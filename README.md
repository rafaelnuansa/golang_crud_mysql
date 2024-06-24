# Backend API dengan Go dan Gin Framework

Ini adalah proyek backend API sederhana yang dibangun dengan bahasa pemrograman Go dan framework web Gin. Proyek ini menunjukkan operasi dasar CRUD (Create, Read, Update, Delete) menggunakan database MySQL. Tujuan utama dari proyek ini adalah memberikan platform pembelajaran bagi para penggemar Go-Lang.

## Fitur

- Koneksi ke database MySQL
- Membuat, Membaca, Memperbarui, dan Menghapus postingan
- API RESTful sederhana dengan respons JSON

## Prasyarat

Sebelum menjalankan proyek ini, pastikan Anda telah menginstal:

- [Go](https://golang.org/dl/)
- [MySQL](https://www.mysql.com/downloads/)
- [Git](https://git-scm.com/)

## Memulai

### Instalasi

1. **Kloning repository:**

    ```bash
    git clone https://github.com/rafaelnuansa/backend-api.git
    cd backend-api
    ```

2. **Siapkan database MySQL:**

    Buat database bernama `db_go_api` dan pengguna dengan akses ke dalamnya. Perbarui paket `models/setup.go` untuk menggunakan kredensial MySQL Anda.

3. **Jalankan aplikasi:**

    ```bash
    go run main.go
    ```

    Server akan berjalan di port 3000.

### Endpoint API

- **GET /** - Pesan selamat datang
- **GET /api/posts** - Mendapatkan semua postingan
- **POST /api/posts** - Membuat postingan baru
- **GET /api/posts/:id** - Mendapatkan postingan berdasarkan ID
- **PUT /api/posts/:id** - Memperbarui postingan berdasarkan ID
- **DELETE /api/posts/:id** - Menghapus postingan berdasarkan ID

## Struktur Proyek

- **controllers/** - Berisi logika untuk menangani permintaan HTTP.
- **models/** - Berisi logika untuk berinteraksi dengan database.
- **main.go** - Titik masuk utama dari aplikasi.

## Lisensi

Proyek ini dilisensikan di bawah MIT License.
