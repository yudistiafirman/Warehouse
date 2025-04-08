package models

type Barang struct {
	IDBarang    uint `gorm:"primaryKey;column:id_barang"`
	IDJenis     uint
	JenisBarang JenisBarang `gorm:"foreignKey:IDJenis"`
	IDSatuan    uint
	Satuan      Satuan `gorm:"foreignKey:IDSatuan"`
	NamaBarang  string
	Stok        int
	StokMinimum int
}

type BarangInput struct {
	IDJenis     uint   `json:"id_jenis"`
	IDSatuan    uint   `json:"id_satuan"`
	NamaBarang  string `json:"nama_barang"`
	Stok        int    `json:"stok"`
	StokMinimum int    `json:"stok_minimum"`
}

func (Barang) TableName() string {
	return "barang"
}
