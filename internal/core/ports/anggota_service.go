package ports

import domain "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"

type IAnggotaRepository interface {
	GetAll() ([]*domain.Anggota, error)
}

type IAnggotaService interface {
	GetAll() ([]*domain.Anggota, error)
}
