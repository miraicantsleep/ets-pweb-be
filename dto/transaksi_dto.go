package dto

import (
	"errors"
)

const (
	// Failed
	MESSAGE_FAILED_CREATE_TRANSAKSI  = "failed create transaksi"
	MESSAGE_FAILED_GET_TRANSAKSI     = "failed get transaksi"
	MESSAGE_FAILED_GET_ALL_TRANSAKSI = "failed get all transaksi"
	MESSAGE_FAILED_UPDATE_TRANSAKSI  = "failed update transaksi"
	MESSAGE_FAILED_DELETE_TRANSAKSI  = "failed delete transaksi"

	// Success
	MESSAGE_SUCCESS_CREATE_TRANSAKSI  = "success create transaksi"
	MESSAGE_SUCCESS_GET_TRANSAKSI     = "success get transaksi"
	MESSAGE_SUCCESS_GET_ALL_TRANSAKSI = "success get all transaksi"
	MESSAGE_SUCCESS_UPDATE_TRANSAKSI  = "success update transaksi"
	MESSAGE_SUCCESS_DELETE_TRANSAKSI  = "success delete transaksi"
)

var (
	ErrInvalidTransaksiType = errors.New("invalid transaksi type")
	ErrCreateTransaksi      = errors.New("failed create transaksi")
	ErrGetTransaksiById     = errors.New("failed get transaksi by id")
	ErrTransaksiNotFound    = errors.New("transaksi not found")
	ErrUpdateTransaksi      = errors.New("failed update transaksi")
	ErrDeleteTransaksi      = errors.New("failed delete transaksi")
)

type (
	CreateTransaksiRequest struct {
		Name   string `json:"name" binding:"required"`
		Type   string `json:"type" binding:"required"`
		Amount int    `json:"amount" binding:"required"`
		Notes  string `json:"notes"`
	}

	CreateTransaksiResponse struct {
		ID     string `json:"id"`
		Owner  string `json:"owner"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Amount int    `json:"amount"`
		Notes  string `json:"notes"`
	}

	GetDetailTransaksiResponse struct {
		ID     string `json:"id"`
		Owner  string `json:"owner"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Amount int    `json:"amount"`
		Notes  string `json:"notes"`
	}

	GetAllTransaksiResponse struct {
		Count int                          `json:"count"`
		Data  []GetDetailTransaksiResponse `json:"data"`
	}

	UpdateTransaksiRequest struct {
		Name   string `json:"name" binding:"required"`
		Type   string `json:"type" binding:"required"`
		Amount int    `json:"amount" binding:"required"`
		Notes  string `json:"notes"`
	}

	UpdateTransaksiResponse struct {
		ID     string `json:"id"`
		Owner  string `json:"owner"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Amount int    `json:"amount"`
		Notes  string `json:"notes"`
	}
)
