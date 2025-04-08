package helpers

import (
	"errors"
	"warehouse-backend/models"

	"gorm.io/gorm"
)

// ValidateBarangInput mengecek apakah input Barang valid.
// Parameter isUpdate digunakan untuk membedakan antara create (isUpdate==false)
// dan update (isUpdate==true). Untuk update, validasi referensial hanya dilakukan jika
// nilai IDJenis atau IDSatuan tidak nol.
func ValidateBarangInput(db *gorm.DB, input models.BarangInput, isUpdate bool) error {
	// Validasi nilai numeric
	if input.Stok < 0 {
		return errors.New("stok tidak boleh negatif")
	}
	if input.StokMinimum < 0 {
		return errors.New("stok minimum tidak boleh negatif")
	}

	// Untuk create, pastikan IDJenis dan IDSatuan tidak nol
	if !isUpdate {
		if input.IDJenis == 0 {
			return errors.New("id_jenis harus diisi")
		}
		if input.IDSatuan == 0 {
			return errors.New("id_satuan harus diisi")
		}
	}

	// Validasi referensial untuk id_jenis
	if input.IDJenis != 0 {
		var jenis models.JenisBarang
		if err := db.First(&jenis, input.IDJenis).Error; err != nil {
			return errors.New("id_jenis tidak valid")
		}
	}

	// Validasi referensial untuk id_satuan
	if input.IDSatuan != 0 {
		var satuan models.Satuan
		if err := db.First(&satuan, input.IDSatuan).Error; err != nil {
			return errors.New("id_satuan tidak valid")
		}
	}

	return nil
}
