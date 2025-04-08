package models

type JenisBarang struct {
	IDJenis   uint   `gorm:"primaryKey;column:id_jenis"`
	NamaJenis string `gorm:"unique"`
}

func (JenisBarang) TableName() string {
	return "jenis_barang"
}
