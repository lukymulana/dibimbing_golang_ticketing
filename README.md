# Ticketing System API

Sistem manajemen event dan tiket berbasis Golang. Mendukung registrasi user, login, CRUD event, pembelian tiket, laporan penjualan, serta role-based access control (RBAC).

---

## 🚀 Fitur Utama
- **User Management:** Register & login, autentikasi JWT, role admin/user.
- **Event Management:** CRUD event (admin), validasi kapasitas & harga, status event.
- **Ticket Management:** Beli tiket, lihat/cancel tiket sendiri (user), validasi kapasitas.
- **Reporting:** Laporan summary & per event (admin).
- **RBAC:** Endpoint dibatasi sesuai role.
- **Pagination & Filter:** Daftar event/tiket bisa dipaging & filter.
- **Dockerized:** Mudah dijalankan dengan Docker Compose.

---

## 📂 Struktur Folder

```
.
├── config/         # Konfigurasi DB & env
├── controller/     # Handler endpoint
├── dto/            # Data transfer object (validasi input/output)
├── entity/         # Model database
├── middleware/     # JWT & RBAC
├── repository/     # Query ke database
├── routes/         # Routing endpoint
├── service/        # Bisnis logic
├── .env            # Konfigurasi environment
├── Dockerfile      # Build container
├── docker-compose.yml
├── go.mod, go.sum
├── main.go         # Entry point
└── README.md
```

---

## ⚙️ Instalasi & Menjalankan

### 1. **Clone & Setup**
```sh
git clone <repo-url>
cd dibimbing_golang_ticketing
```

### 2. **Atur file .env**
Contoh:
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=ticketing
JWT_SECRET=your_jwt_secret
APP_PORT=8080
```

### 3. **Jalankan dengan Docker Compose**
```sh
docker-compose up --build
```

### 4. **Akses API**
- http://localhost:8080

---

## 🛠️ Endpoint Utama

### User
- `POST /register` — Register user baru
- `POST /login` — Login user

### Event (Admin)
- `GET /events` — Lihat daftar event (public)
- `POST /events` — Tambah event baru
- `PUT /events/:id` — Update event
- `DELETE /events/:id` — Hapus event

### Ticket (User)
- `GET /tickets` — Lihat tiket sendiri
- `POST /tickets` — Beli tiket
- `GET /tickets/:id` — Detail tiket
- `PATCH /tickets/:id` — Cancel tiket

### Report (Admin)
- `GET /reports/summary` — Laporan ringkasan
- `GET /reports/event/:id` — Laporan per event

---

## 🔐 Role & Autentikasi
- Gunakan header `Authorization: Bearer <token>` untuk endpoint yang butuh login.
- Admin hanya bisa mengakses endpoint event & report tertentu.
- User hanya bisa mengakses/mengelola tiket miliknya sendiri.

---

## 📄 Dokumentasi API
- Import file `ticketing-system.postman_collection.json` ke Postman untuk contoh request lengkap.
- Set variabel `user_token` dan `admin_token` di Postman setelah login.

---

## 📝 Catatan
- Pastikan DB MySQL sudah berjalan & kredensial sesuai dengan `.env`.
- Untuk development, bisa pakai MySQL lokal atau dari docker-compose.
- Untuk production, pastikan environment variable aman dan tidak commit `.env` ke repository.

---

## 👨‍💻 Kontribusi & Lisensi
Pull request sangat diterima. Silakan fork dan modifikasi sesuai kebutuhan!
