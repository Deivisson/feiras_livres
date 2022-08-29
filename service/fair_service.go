package service

import (
	"encoding/json"

	"github.com/Deivisson/feiras_livres/domain"
	"github.com/Deivisson/feiras_livres/utils/errs"
)

type FairService interface {
	Create(params []byte) (*domain.Fair, *errs.AppError)
	Update(params []byte, id string) (*domain.Fair, *errs.AppError)
	BulkCreate(fairs []domain.Fair) *errs.AppError
	Delete(id string) *errs.AppError
	Search(params []byte) ([]domain.Fair, *errs.AppError)
	GetById(id string) (*domain.Fair, *errs.AppError)
}

type DefaultFairService struct {
	repo domain.FairRepository
}

func (s DefaultFairService) Create(params []byte) (*domain.Fair, *errs.AppError) {
	fair := domain.Fair{}
	if err := json.Unmarshal(params, &fair); err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	if err := fair.Validate(); err != nil {
		return nil, err
	}

	err := s.repo.Create(&fair)
	return &fair, err
}

func (s DefaultFairService) BulkCreate(fairs []domain.Fair) *errs.AppError {
	return s.repo.BulkCreate(fairs)
}

func (s DefaultFairService) Update(params []byte, id string) (*domain.Fair, *errs.AppError) {
	fair, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(params, &fair); err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	if err := fair.Validate(); err != nil {
		return nil, err
	}

	err = s.repo.Update(fair)
	return fair, err
}

func (s DefaultFairService) Search(params []byte) ([]domain.Fair, *errs.AppError) {
	dto := domain.FairSearchRequestDTO{}

	if err := json.Unmarshal(params, &dto); err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	fairs, err := s.repo.Search(&dto)
	if err != nil {
		return nil, err
	}
	return fairs, nil
}

func (s DefaultFairService) Delete(id string) *errs.AppError {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultFairService) GetById(id string) (*domain.Fair, *errs.AppError) {
	fair, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return fair, nil
}

func NewFairService(repository domain.FairRepository) DefaultFairService {
	return DefaultFairService{repository}
}
