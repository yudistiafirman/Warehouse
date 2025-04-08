# Dokumentasi API Warehouse

API ini digunakan untuk mengelola data barang di gudang. API ini dilindungi dengan autentikasi JWT, kecuali endpoint login.

## Autentikasi

Untuk mengakses endpoint yang dilindungi, Anda perlu menyertakan token JWT dalam header `Authorization` dengan format:
`Authorization: Bearer <token>`
