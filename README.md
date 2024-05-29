# Proyek Presensee

## Deskripsi Singkat
Proyek ini adalah sistem manajemen jadwal dan absensi untuk mahasiswa dan dosen. Sistem ini dibangun menggunakan Go dan menggunakan GORM sebagai ORM untuk berinteraksi dengan database.

## Cara Menjalankan Proyek

1. **Clone repositori ini:**
    ```bash
    git clone https://github.com/username/repo-name.git
    ```

2. **Masuk ke direktori proyek:**
    ```bash
    cd repo-name
    ```

3. **Install dependencies:**
    ```bash
    npm install
    ```

4. **Jalankan proyek:**
    ```bash
    go run main.go
    ```

## Struktur Folder
- **controller/**: Berisi handler untuk setiap endpoint API.
- **model/**: Berisi definisi model untuk entitas dalam database.
- **repository/**: Berisi implementasi untuk berinteraksi dengan database.
- **usecase/**: Berisi logika bisnis aplikasi.
- **utils/**: Berisi utilitas dan helper yang digunakan dalam proyek.

## Endpoint API
### Mahasiswa
- **GET /mahasiswas**: Mendapatkan semua data mahasiswa.
- **GET /mahasiswa/:id**: Mendapatkan data mahasiswa berdasarkan ID.
- **POST /mahasiswa**: Membuat data mahasiswa baru.
- **PUT /mahasiswa/:id**: Memperbarui data mahasiswa berdasarkan ID.
- **DELETE /mahasiswa/:id**: Menghapus data mahasiswa berdasarkan ID.

### Matakuliah
- **GET /matakuliahs**: Mendapatkan semua data matakuliah.
- **GET /matakuliah/:id**: Mendapatkan data matakuliah berdasarkan ID.
- **POST /matakuliah**: Membuat data matakuliah baru.
- **PUT /matakuliah/:id**: Memperbarui data matakuliah berdasarkan ID.
- **DELETE /matakuliah/:id**: Menghapus data matakuliah berdasarkan ID.

### Jadwal
- **GET /jadwals**: Mendapatkan semua data jadwal.
- **GET /jadwal/:id**: Mendapatkan data jadwal berdasarkan ID.
- **POST /jadwal**: Membuat data jadwal baru.
- **PUT /jadwal/:id**: Memperbarui data jadwal berdasarkan ID.
- **DELETE /jadwal/:id**: Menghapus data jadwal berdasarkan ID.

## Kontribusi
1. Fork repositori ini.
2. Buat branch fitur baru (`git checkout -b fitur-baru`).
3. Commit perubahan (`git commit -m 'Menambahkan fitur baru'`).
4. Push ke branch (`git push origin fitur-baru`).
5. Buat Pull Request.

## Lisensi
Proyek ini dilisensikan di bawah MIT License.

