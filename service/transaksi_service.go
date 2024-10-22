package service

import (
	"github.com/adieos/ets-pweb-be/dto"
	"github.com/adieos/ets-pweb-be/entity"
	"github.com/adieos/ets-pweb-be/repository"
)

type (
	TransaksiService interface {
		CreateTransaksi(userId string, req dto.CreateTransaksiRequest) (dto.CreateTransaksiResponse, error)
		GetAllTransaksi(ownerId string, role string) (dto.GetAllTransaksiResponse, error)
		GetDetailTransaksi(transaksiId string) (dto.GetDetailTransaksiResponse, error)
		UpdateTransaksi(transaksiId string, req dto.UpdateTransaksiRequest) (dto.UpdateTransaksiResponse, error)
		DeleteTransaksi(transaksiId string) error
	}

	transaksiService struct {
		transaksiRepo repository.TransaksiRepository
	}
)

func NewTransaksiService(transaksiRepo repository.TransaksiRepository, jwtService JWTService) TransaksiService {
	return &transaksiService{
		transaksiRepo: transaksiRepo,
	}
}

func (s *transaksiService) CreateTransaksi(userId string, req dto.CreateTransaksiRequest) (dto.CreateTransaksiResponse, error) {
	transaksi := entity.Transaksi{
		UserID: userId,
		Name:   req.Name,
		Type:   req.Type,
		Amount: req.Amount,
		Notes:  req.Notes,
	}

	transReg, err := s.transaksiRepo.CreateTransaksi(transaksi)
	if err != nil {
		return dto.CreateTransaksiResponse{}, err
	}

	res := dto.CreateTransaksiResponse{
		ID:     transReg.ID.String(),
		Owner:  transReg.UserID,
		Name:   transReg.Name,
		Type:   transReg.Type,
		Amount: transReg.Amount,
		Notes:  transReg.Notes,
	}

	return res, nil
}

// flis jangan ngebug
func (s *transaksiService) GetAllTransaksi(userId string, role string) (dto.GetAllTransaksiResponse, error) {
	var ownerId string
	if role == "ADMIN" {
		ownerId = "ADMIN"
	} else {
		ownerId = userId
	}

	data, err := s.transaksiRepo.GetAllTransaksi(ownerId)
	if err != nil {
		return dto.GetAllTransaksiResponse{}, err
	}

	var res []dto.GetDetailTransaksiResponse
	for _, transaksi := range data {
		res = append(res, dto.GetDetailTransaksiResponse{
			ID:     transaksi.ID.String(),
			Owner:  transaksi.UserID,
			Name:   transaksi.Name,
			Type:   transaksi.Type,
			Amount: transaksi.Amount,
			Notes:  transaksi.Notes,
		})
	}

	return dto.GetAllTransaksiResponse{
		Count: len(res),
		Data:  res,
	}, nil
}

func (s *transaksiService) GetDetailTransaksi(transaksiId string) (dto.GetDetailTransaksiResponse, error) {
	transaksi, err := s.transaksiRepo.GetTransaksiById(transaksiId)
	if err != nil {
		return dto.GetDetailTransaksiResponse{}, err
	}

	return dto.GetDetailTransaksiResponse{
		ID:     transaksi.ID.String(),
		Owner:  transaksi.UserID,
		Name:   transaksi.Name,
		Type:   transaksi.Type,
		Amount: transaksi.Amount,
		Notes:  transaksi.Notes,
	}, nil
}

func (s *transaksiService) UpdateTransaksi(transaksiId string, req dto.UpdateTransaksiRequest) (dto.UpdateTransaksiResponse, error) {

	transaksiOld, err := s.transaksiRepo.GetTransaksiById(transaksiId)
	if err != nil {
		return dto.UpdateTransaksiResponse{}, err
	}

	transaksiNew := entity.Transaksi{
		ID:     transaksiOld.ID,
		UserID: transaksiOld.UserID,
		Name:   req.Name,
		Type:   req.Type,
		Amount: req.Amount,
		Notes:  req.Notes,
	}

	result, err := s.transaksiRepo.UpdateTransaksi(transaksiNew)
	if err != nil {
		return dto.UpdateTransaksiResponse{}, err
	}

	return dto.UpdateTransaksiResponse{
		ID:     result.ID.String(),
		Owner:  result.UserID,
		Name:   result.Name,
		Type:   result.Type,
		Amount: result.Amount,
		Notes:  result.Notes,
	}, nil
}

func (s *transaksiService) DeleteTransaksi(transaksiId string) error {
	err := s.transaksiRepo.DeleteTransaksi(transaksiId)

	// ini kode apa lol
	if err != nil {
		return err
	}

	return nil
}
