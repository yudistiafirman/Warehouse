package models

type Satuan struct {
	IDSatuan   uint   `gorm:"primaryKey;column:id_satuan"`
	NamaSatuan string `gorm:"unique"`
}

func (Satuan) TableName() string {
	return "satuan"
}
