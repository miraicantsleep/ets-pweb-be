package repository

import (
	"github.com/adieos/ets-pweb-be/entity"
	"gorm.io/gorm"
)

type (
	TransaksiRepository interface {
		CreateTransaksi(transaksi entity.Transaksi) (entity.Transaksi, error)
		GetTransaksiById(transaksiId string) (entity.Transaksi, error)
		GetAllTransaksi(ownerId string) ([]entity.Transaksi, error) // can be used for komunal, if admin ownerId = NULL
		UpdateTransaksi(transaksi entity.Transaksi) (entity.Transaksi, error)
		DeleteTransaksi(transaksiId entity.Transaksi) error
	}

	transaksiRepository struct {
		db *gorm.DB
	}
)

func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiRepository{
		db: db,
	}
}

func (r *transaksiRepository) CreateTransaksi(transaksi entity.Transaksi) (entity.Transaksi, error) {
	if err := r.db.Create(&transaksi).Error; err != nil {
		return entity.Transaksi{}, err
	}

	return transaksi, nil
}

func (r *transaksiRepository) GetTransaksiById(transaksiId string) (entity.Transaksi, error) {
	var transaksi entity.Transaksi
	if err := r.db.Where("id = ?", transaksiId).Take(&transaksi).Error; err != nil {
		return entity.Transaksi{}, err
	}

	return transaksi, nil
}

// get all transaksi here
func (r *transaksiRepository) GetAllTransaksi(ownerId string) ([]entity.Transaksi, error) {
	var transaksis []entity.Transaksi
	if ownerId == "ADMIN" {
		if err := r.db.Find(&transaksis).Error; err != nil {
			return nil, err
		}
		return transaksis, nil
	}

	if err := r.db.Where("user_id = ?", ownerId).Find(&transaksis).Error; err != nil {
		return nil, err
	}

	return transaksis, nil
}

// get trans id dulu
func (r *transaksiRepository) UpdateTransaksi(transaksi entity.Transaksi) (entity.Transaksi, error) {
	if err := r.db.Save(&transaksi).Error; err != nil {
		return entity.Transaksi{}, err
	}

	return transaksi, nil
}

// MUST TEST! gorm docs stinks man
func (r *transaksiRepository) DeleteTransaksi(transaksiId entity.Transaksi) error {
	if err := r.db.Delete(&transaksiId).Error; err != nil {
		// if err := r.db.Delete(&entity.Transaksi{}, transaksiId).Error; err != nil {
		return err
	}

	return nil
}
