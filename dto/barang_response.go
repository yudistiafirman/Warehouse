package dto

type JenisBarangResponse struct {
	IDJenis   uint   `json:"idJenis"`
	NamaJenis string `json:"namaJenis"`
}

type SatuanResponse struct {
	IDSatuan   uint   `json:"idSatuan"`
	NamaSatuan string `json:"namaSatuan"`
}

type BarangResponse struct {
	IDBarang    uint                `json:"idBarang"`
	IDJenis     uint                `json:"idJenis"`
	JenisBarang JenisBarangResponse `json:"jenisBarang"`
	IDSatuan    uint                `json:"idSatuan"`
	Satuan      SatuanResponse      `json:"satuan"`
	NamaBarang  string              `json:"namaBarang"`
	Stok        int                 `json:"stok"`
	StokMinimum int                 `json:"stokMinimum"`
}
