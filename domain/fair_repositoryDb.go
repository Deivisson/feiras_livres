package domain

import (
	"errors"
	"fmt"

	"github.com/Deivisson/feiras_livres/utils/errs"
	"gorm.io/gorm"
)

type FairRepositoryDb struct {
	dbClient *gorm.DB
}

func (r FairRepositoryDb) Create(fair *Fair) *errs.AppError {
	err := r.dbClient.Debug().Create(&fair).Error
	if err != nil {
		if fair.Validation.HasError() {
			return errs.NewValidationError(fair.Validation.Errors)
		} else {
			return errs.NewUnexpectedError(err.Error())
		}
	}
	return nil
}

func (r FairRepositoryDb) Update(fair *Fair) *errs.AppError {
	err := r.dbClient.Debug().Updates(&fair).Error
	if err != nil {
		if fair.Validation.HasError() {
			return errs.NewValidationError(fair.Validation.Errors)
		} else {
			return errs.NewUnexpectedError(err.Error())
		}
	}
	return nil
}

func (r FairRepositoryDb) Search(filter *FairSearchRequestDTO) ([]Fair, *errs.AppError) {
	fairs := []Fair{}
	query := r.dbClient.Debug().Model(Fair{})
	applyFilter(query, "fair_name", filter.FairName)
	applyFilter(query, "district", filter.District)
	applyFilter(query, "neighborhood", filter.Neighborhood)
	applyFilter(query, "region5", filter.Region5)

	err := query.Scan(&fairs).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return fairs, nil
}

func applyFilter(query *gorm.DB, fieldName, value string) {
	if value != "" {
		query.Where(fieldName+" ILike ? ", fmt.Sprint("%", value, "%"))
	}
}

func (r FairRepositoryDb) FindById(id string) (*Fair, *errs.AppError) {
	fair := Fair{}
	err := r.dbClient.Debug().Model(Fair{}).Where("id = ?", id).Take(&fair).Error
	if err != nil {
		return nil, errs.NewNotFoundError("Fair not found to then given ID")
	}
	return &fair, nil
}

func (r FairRepositoryDb) Delete(id string) *errs.AppError {
	fair, err := r.FindById(id)
	if err != nil {
		return err
	}

	if err := r.dbClient.Delete(&fair).Error; err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}

func (r FairRepositoryDb) BulkCreate(fairs []Fair) *errs.AppError {
	if err := r.dbClient.Debug().Create(fairs).Error; err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}

func (s FairRepositoryDb) HasAny() (bool, *errs.AppError) {
	var fair Fair
	err := s.dbClient.First(&fair).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errs.NewUnexpectedError(err.Error())
	}
	return true, nil
}

func NewFairRepositoryDb(dbClient *gorm.DB) FairRepositoryDb {
	return FairRepositoryDb{dbClient}
}
