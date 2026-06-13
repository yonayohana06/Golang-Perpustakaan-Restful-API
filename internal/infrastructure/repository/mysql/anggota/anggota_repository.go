package anggota

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"gorm.io/gorm"
)

type AnggotaRepository struct {
	db *gorm.DB
}

func NewAnggotaRepository(db *gorm.DB) *AnggotaRepository {
	return &AnggotaRepository{db: db}
}

func (r *AnggotaRepository) GetAll() ([]*domain.Anggota, error) {
	var anggota []*domain.Anggota
	err := r.db.Find(&anggota).Error
	return anggota, err
}
