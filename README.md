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

  ```diff
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
 

