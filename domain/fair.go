package domain

import (
	"fmt"

	"github.com/Deivisson/feiras_livres/utils/errs"
)

type Fair struct {
	Id                string     `gorm:"size:8;not null;primaryKey" json:"id"`
	Longitude         string     `gorm:"size:10;not null" json:"longitude"`
	Latitude          string     `gorm:"size:10;not null" json:"latitude"`
	Sector            string     `gorm:"size:15;not null" json:"setorCensitario"`
	Area              string     `gorm:"size:13;not null" json:"areaPonderacao"`
	DistrictCode      string     `gorm:"size:9;not null" json:"codigoDistrito"`
	District          string     `gorm:"size:18;not null" json:"distrito"`
	SubprefectureCode string     `gorm:"size:2;not null" json:"codigoSubprefeitura"`
	SubprefectureName string     `gorm:"size:25;not null" json:"nomeSubPrefeitura"`
	Region5           string     `gorm:"size:6;not null" json:"regiao5"`
	Region8           string     `gorm:"size:7;not null" json:"regiao8"`
	FairName          string     `gorm:"size:30;not null" json:"nomeFeira"`
	Registry          string     `gorm:"size:6;not null" json:"registro"`
	Address           string     `gorm:"size:34;not null" json:"logradouro"`
	Number            string     `gorm:"size:5;not null" json:"numero"`
	Neighborhood      string     `gorm:"size:20;not null" json:"bairro"`
	Reference         string     `gorm:"size:50;not null" json:"referencia"`
	Validation        Validation `gorm:"-" json:"-"`
}

type FairSearchRequestDTO struct {
	District     string `json:"distrito"`
	Region5      string `json:"regiao5"`
	FairName     string ` json:"nomeFeira"`
	Neighborhood string `json:"bairro"`
}

//go:generate mockgen -destination=../mocks/domain/mockFairRepository.go -package=domain github.com/Deivisson/feiras_livres/domain FairRepository
type FairRepository interface {
	Create(*Fair) *errs.AppError
	BulkCreate([]Fair) *errs.AppError
	Update(*Fair) *errs.AppError
	Search(*FairSearchRequestDTO) ([]Fair, *errs.AppError)
	FindById(id string) (*Fair, *errs.AppError)
	Delete(id string) *errs.AppError
	HasAny() (bool, *errs.AppError)
}

func (f *Fair) Validate() *errs.AppError {
	f.Validation.AddErrorField("longitude", validateField(f.Longitude, true, 10))
	f.Validation.AddErrorField("latitude", validateField(f.Latitude, true, 10))
	f.Validation.AddErrorField("setorCensitario", validateField(f.Sector, true, 15))
	f.Validation.AddErrorField("areaPonderacao", validateField(f.Area, true, 13))
	f.Validation.AddErrorField("codigoDistrito", validateField(f.DistrictCode, true, 9))
	f.Validation.AddErrorField("distrito", validateField(f.District, true, 18))
	f.Validation.AddErrorField("codigoSubprefeitura", validateField(f.SubprefectureCode, true, 2))
	f.Validation.AddErrorField("nomeSubPrefeitura", validateField(f.SubprefectureName, true, 25))
	f.Validation.AddErrorField("regiao5", validateField(f.Region5, true, 6))
	f.Validation.AddErrorField("regiao8", validateField(f.Region8, true, 7))
	f.Validation.AddErrorField("registro", validateField(f.Registry, true, 6))
	f.Validation.AddErrorField("nomeFeira", validateField(f.FairName, true, 30))
	f.Validation.AddErrorField("logradouro", validateField(f.Address, true, 34))
	f.Validation.AddErrorField("numero", validateField(f.Number, true, 5))
	f.Validation.AddErrorField("bairro", validateField(f.Neighborhood, true, 20))
	f.Validation.AddErrorField("referencia", validateField(f.Reference, true, 50))

	if f.Validation.HasError() {
		return errs.NewValidationError(f.Validation.Errors)
	}
	return nil
}

func validateField(value string, required bool, length int) []string {
	var errors []string
	if value == "" && required {
		errors = append(errors, "Can't be blank!")
	}

	if len(value) > length {
		errors = append(errors, fmt.Sprintf("Max length is %d.", length))
	}
	return errors
}
