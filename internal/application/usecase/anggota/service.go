package anggota

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"
)

type AnggotaService struct {
	repo ports.IAnggotaRepository
}

func NewAnggotaUseCase(repo ports.IAnggotaRepository) *AnggotaService {
	return &AnggotaService{repo}
}

func (s *AnggotaService) GetAll() ([]*domain.Anggota, error) {

	return s.repo.GetAll()
}
