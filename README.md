# Dokumentasi API Warehouse

API ini digunakan untuk mengelola data barang di gudang. API ini dilindungi dengan autentikasi JWT, kecuali endpoint login.

## Autentikasi

Untuk mengakses endpoint yang dilindungi, Anda perlu menyertakan token JWT dalam header `Authorization` dengan format:
`Authorization: Bearer <token>` Token diperoleh melalui endpoint `/login`.

## Endpoint

### 1. Login

- **Path**: `/login`
- **Metode HTTP**: `POST`
- **Autentikasi**: Tidak diperlukan
- **Header**: 
  - `Content-Type: application/json`
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string"
  }

- ```diff
  - username: Nama pengguna (wajib)
  - password: Kata sandi (wajib)

- **Response**:
  - **Status 200 OK**:
     ```json
     {
        "token": "string"
     }
  - **Status 400 Bad Request**: Jika format input salah
  - **Status 401 Unauthorized**: Jika username tidak ditemukan atau password salah
  - Status 500 Internal Server Error: Jika terjadi kesalahan server

### 1. Membuat Barang Baru

- **Path**: `/barang/create`
- **Metode HTTP**: `POST`
- **Autentikasi**: Diperlukan (JWT)
- **Header**: 
  - `Content-Type: application/json`
  - `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
    "id_jenis": "integer",
    "id_satuan": "integer",
    "nama_barang": "string",
    "stok": "integer",
    "stok_minimum": "integer"
  }

- ```diff
  - id_jenis: ID jenis barang (wajib, harus ada di tabel jenis_barang)
  - id_satuan: ID satuan barang (wajib, harus ada di tabel satuan)
  - nama_barang: Nama barang (wajib)
  - stok: Jumlah stok awal (wajib, tidak boleh negatif)
  - stok_minimum: Stok minimum (wajib, tidak boleh negatif)

- **Response**:
  - **Status 201 Created**:
     ```json
     {
        "message": "barang berhasil dibuat",
        "barang": {
          "idBarang": "integer",
          "idJenis": "integer",
          "jenisBarang": {
            "idJenis": "integer",
            "namaJenis": "string"
          },
        "idSatuan": "integer",
        "satuan": {
          "idSatuan": "integer",
          "namaSatuan": "string"
        },
        "namaBarang": "string",
        "stok": "integer",
        "stokMinimum": "integer"
      }
    }
  - **Status 400 Bad Request**: Jika input tidak valid atau validasi gagal
  - **Status 401 Unauthorized**: Jika token tidak valid atau hilang
  - Status 500 Internal Server Error: Jika terjadi kesalahan server
 

