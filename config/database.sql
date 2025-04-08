-- Setup Struktur Database Warehosue
CREATE DATABASE IF NOT EXISTS warehouse;
USE warehouse;

-- Tabel `users`
CREATE TABLE users (
    id_user INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    nama_lengkap VARCHAR(255),
    level ENUM('admin', 'admin_gudang', 'kepala_gudang') NOT NULL
);

-- Tabel `jenis_barang`
CREATE TABLE jenis_barang (
    id_jenis INT AUTO_INCREMENT PRIMARY KEY,
    nama_jenis VARCHAR(255) UNIQUE NOT NULL
);

-- Tabel `satuan`
CREATE TABLE satuan (
    id_satuan INT AUTO_INCREMENT PRIMARY KEY,
    nama_satuan VARCHAR(255) UNIQUE NOT NULL
);

-- Tabel `barang`
CREATE TABLE barang (
    id_barang INT AUTO_INCREMENT PRIMARY KEY,
    id_jenis INT NOT NULL,
    id_satuan INT NOT NULL,
    nama_barang VARCHAR(255) NOT NULL,
    stok INT NOT NULL DEFAULT 0,
    stok_minimum INT NOT NULL DEFAULT 0,
    FOREIGN KEY (id_jenis) REFERENCES jenis_barang(id_jenis) ON DELETE CASCADE,
    FOREIGN KEY (id_satuan) REFERENCES satuan(id_satuan) ON DELETE CASCADE
);


-- Trigger untuk Update Stok pada Barang
DELIMITER $$



DELIMITER ;

-- Insert Data Awal
INSERT INTO users (username, password, nama_lengkap, level)
VALUES
    -- password : Admin1234
    ('admin', '$2y$10$R3IqiZSysEAveFSBGBKlvuxfCZ3397ZkCWr.6aHgrboSST60zukpG', 'Administrator', 'admin'), 
     -- password : Admingudang1234
    ('admin_gudang', '$2y$10$yQ1D9GSQ0dVOLeKdWgIQkuDbI2wKPZMTOSJUVibkHn7wopg3NTHHC', 'Admin Gudang', 'adming_udang'),
     -- password : Kepalagudang1234
    ('kepala_gudang', '$2y$10$ce1DNMyxFZQOosX/ZhEgJO9xoEp4V.PIIx25MGortv4S8mnghtW66', 'Kepala Gudang', 'kepala_gudang');

INSERT INTO jenis_barang (nama_jenis)
VALUES ('Elektronik'), ('Pakaian'), ('Makanan');

INSERT INTO satuan (nama_satuan)
VALUES ('pcs'), ('kg'), ('meter');
