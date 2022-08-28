package domain

import (
	"fmt"
)

type FairRepositoryStub struct {
	fairs []Fair
}

func (s FairRepositoryStub) Search() ([]Fair, error) {
	return s.fairs, nil
}

func (s FairRepositoryStub) FindById(id string) (*Fair, error) {
	for _, v := range s.fairs {
		if v.Id == id {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("not found by Id: %s", id)
}

func (s FairRepositoryStub) BulkCreate(fairs []Fair) error {
	s.fairs = fairs
	return nil
}

func (s FairRepositoryStub) HasAny() (bool, error) {
	return false, nil
}

func NewFairRepositoryStub() FairRepositoryStub {
	fairs := []Fair{
		{
			Id: "1",
			FairCommon: FairCommon{
				Longitude:         "-46550164",
				Latitude:          "-23558733",
				Sector:            "355030885000091",
				Area:              "3550308005040",
				DistrictCode:      "87",
				District:          "VILA FORMOSA",
				SubprefectureCode: "26",
				SubprefectureName: "ARICANDUVA-FORMOSA-CARRAO",
				Region5:           "Leste",
				Region8:           "Leste 1",
				FairName:          "VILA FORMOSA",
				Registry:          "4041-0",
				Address:           "RUA MARAGOJIPE",
				Number:            "S/N",
				Neighborhood:      "VL FORMOSA",
				Reference:         "TV RUA PRETORIA",
			},
		},
		{
			Id: "741",
			FairCommon: FairCommon{
				Longitude:         "-46515046",
				Latitude:          "-23583043",
				Sector:            "355030804000097",
				Area:              "3550308005151",
				DistrictCode:      "4",
				District:          "ARICANDUVA",
				SubprefectureCode: "26",
				SubprefectureName: "ARICANDUVA-FORMOSA-CARRAO",
				Region5:           "Leste",
				Region8:           "Leste 1",
				FairName:          "VILA RICA",
				Registry:          "1049-9",
				Address:           "RUA PROF ALZIRA DE O GILIOLI",
				Number:            "1817",
				Neighborhood:      "VL Rica",
				Reference:         "CENTRO ESPORTIVO MUNICIPAL",
			},
		},
	}
	return FairRepositoryStub{fairs}
}
