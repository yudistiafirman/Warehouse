package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"warehouse-backend/dto"
	"warehouse-backend/helpers"
	"warehouse-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BarangHandler struct {
	DB *gorm.DB
}

func NewBarangHandler(db *gorm.DB) *BarangHandler {
	return &BarangHandler{DB: db}
}

// GetAllBarang menangani permintaan GET semua barang
func (h *BarangHandler) GetAllBarang(c *gin.Context) {
	var barangList []models.Barang
	if err := h.DB.Preload("JenisBarang").Preload("Satuan").Find(&barangList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data barang"})
		return
	}

	var response []dto.BarangResponse
	for _, b := range barangList {
		response = append(response, dto.BarangResponse{
			IDBarang:    b.IDBarang,
			IDJenis:     b.IDJenis,
			IDSatuan:    b.IDSatuan,
			NamaBarang:  b.NamaBarang,
			Stok:        b.Stok,
			StokMinimum: b.StokMinimum,
			JenisBarang: dto.JenisBarangResponse{
				IDJenis:   b.JenisBarang.IDJenis,
				NamaJenis: b.JenisBarang.NamaJenis,
			},
			Satuan: dto.SatuanResponse{
				IDSatuan:   b.Satuan.IDSatuan,
				NamaSatuan: b.Satuan.NamaSatuan,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// CreateBarang menangani pembuatan barang baru
func (h *BarangHandler) CreateBarang(c *gin.Context) {
	var input models.BarangInput

	// Bind JSON request ke struct BarangInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input tidak valid"})
		return
	}

	// Validasi input (create: isUpdate = false)
	if err := helpers.ValidateBarangInput(h.DB, input, false); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil data referensial untuk response
	var jenis models.JenisBarang
	h.DB.First(&jenis, input.IDJenis)

	var satuan models.Satuan
	h.DB.First(&satuan, input.IDSatuan)

	// Buat instance barang baru
	barang := models.Barang{
		IDJenis:     input.IDJenis,
		IDSatuan:    input.IDSatuan,
		NamaBarang:  input.NamaBarang,
		Stok:        input.Stok,
		StokMinimum: input.StokMinimum,
	}

	if err := h.DB.Create(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membuat barang"})
		return
	}

	// Buat response menggunakan DTO
	response := dto.BarangResponse{
		IDBarang:    barang.IDBarang,
		IDJenis:     barang.IDJenis,
		JenisBarang: dto.JenisBarangResponse{IDJenis: jenis.IDJenis, NamaJenis: jenis.NamaJenis},
		IDSatuan:    barang.IDSatuan,
		Satuan:      dto.SatuanResponse{IDSatuan: satuan.IDSatuan, NamaSatuan: satuan.NamaSatuan},
		NamaBarang:  barang.NamaBarang,
		Stok:        barang.Stok,
		StokMinimum: barang.StokMinimum,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "barang berhasil dibuat",
		"barang":  response,
	})
}

// UpdateBarang menangani permintaan UPDATE barang
func (h *BarangHandler) UpdateBarang(c *gin.Context) {
	// Ambil ID dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var existingBarang models.Barang
	result := h.DB.First(&existingBarang, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "barang tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mencari barang"})
		return
	}

	var input models.BarangInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input tidak valid"})
		return
	}

	// Validasi input (update: isUpdate = true)
	if err := helpers.ValidateBarangInput(h.DB, input, true); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateData := map[string]interface{}{
		"id_jenis":     input.IDJenis,
		"id_satuan":    input.IDSatuan,
		"nama_barang":  input.NamaBarang,
		"stok":         input.Stok,
		"stok_minimum": input.StokMinimum,
	}

	if err := h.DB.Model(&existingBarang).Select("*").Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal update barang"})
		return
	}

	// Ambil data terbaru dengan relasi
	h.DB.Preload("JenisBarang").Preload("Satuan").First(&existingBarang, id)

	// Buat DTO response
	response := dto.BarangResponse{
		IDBarang: existingBarang.IDBarang,
		IDJenis:  existingBarang.IDJenis,
		JenisBarang: dto.JenisBarangResponse{
			IDJenis:   existingBarang.JenisBarang.IDJenis,
			NamaJenis: existingBarang.JenisBarang.NamaJenis,
		},
		IDSatuan: existingBarang.IDSatuan,
		Satuan: dto.SatuanResponse{
			IDSatuan:   existingBarang.Satuan.IDSatuan,
			NamaSatuan: existingBarang.Satuan.NamaSatuan,
		},
		NamaBarang:  existingBarang.NamaBarang,
		Stok:        existingBarang.Stok,
		StokMinimum: existingBarang.StokMinimum,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "barang berhasil diupdate",
		"barang":  response,
	})
}

// DeleteBarang menangani permintaan DELETE barang
func (h *BarangHandler) DeleteBarang(c *gin.Context) {
	// Ambil ID dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// Cek apakah barang ada
	var barang models.Barang
	result := h.DB.First(&barang, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Barang tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencari barang"})
		return
	}

	// Hapus barang
	if err := h.DB.Delete(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus barang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Barang berhasil dihapus",
		"id":      id,
	})
}
