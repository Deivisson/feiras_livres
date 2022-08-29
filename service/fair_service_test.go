package service

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Deivisson/feiras_livres/domain"
	mockdomain "github.com/Deivisson/feiras_livres/mocks/domain"
	"github.com/Deivisson/feiras_livres/utils/errs"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockRepo *mockdomain.MockFairRepository
var service FairService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = mockdomain.NewMockFairRepository(ctrl)
	service = NewFairService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_new_fair_when_is_created_successfully(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	fairReq := buildFair(buildDefaultPayload())
	mockRepo.EXPECT().Create(fairReq).Return(nil)

	// Act
	data, _ := json.Marshal(fairReq)
	createdFair, appError := service.Create(data)

	// Assert
	if appError != nil || createdFair == nil {
		t.Error("Test failed while validating creation of new fair")
	}
}

func Test_should_return_an_updated_fair_when_updated_successfully(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// Arrange
	currentFair := buildFair(buildDefaultPayload())
	mockRepo.EXPECT().FindById("1").Return(currentFair, nil)
	mockRepo.EXPECT().Update(currentFair).Return(nil)

	// Act
	updatedFair, err := service.Update(buildPayload("1", "Funny Fair"), "1")

	// Assert
	if err != nil {
		t.Error("Test failed while validating error for new fair")
	}
	assert.Equal(t, updatedFair.Id, "1")
	assert.Equal(t, updatedFair.FairName, "Funny Fair")
}

func Test_should_return_errors_on_update_when_the_request_is_not_valid(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// Arrange
	fairReq := &domain.Fair{Id: "1"}
	mockRepo.EXPECT().FindById("1").Return(&domain.Fair{Id: "1"}, nil)
	data, _ := json.Marshal(fairReq)

	// Act
	_, err := service.Update(data, "1")

	// Assert
	if err == nil {
		t.Error("Test failed while validating error for new fair")
	}

	if err != nil && err.Code != http.StatusBadRequest {
		t.Error("Test should return a error code 400")
	}
}

func Test_should_not_find_when_id_nonexistent(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// Arrange
	mockRepo.EXPECT().FindById("1").Return(nil, errs.NewNotFoundError("Fair not found to then given ID"))

	// Act
	_, err := service.GetById("1")

	// Assert
	if err == nil {
		t.Error("Test failed while try to find nonexistent record")
		return
	}
	assert.Equal(t, err.Code, 404)
	assert.Equal(t, err.Message, "Fair not found to then given ID")
}

func buildFair(payload []byte) *domain.Fair {
	var fair domain.Fair
	json.Unmarshal(payload, &fair)
	return &fair
}

func buildDefaultPayload() []byte {
	return buildPayload("1", "FEIRA VILA FORMOSA")
}
func buildPayload(id, nomeFeira string) []byte {
	data, _ := json.Marshal(
		map[string]string{
			"id":                  id,
			"nomeFeira":           nomeFeira,
			"longitude":           "-99999999",
			"latitude":            "-99999999",
			"setorCensitario":     "355030885000091",
			"areaPonderacao":      "3550308005040",
			"codigoDistrito":      "87",
			"distrito":            "VILA FORMOSA",
			"codigoSubprefeitura": "26",
			"nomeSubPrefeitura":   "ARICANDUVA-FORMOSA-CARRAO",
			"regiao5":             "Leste",
			"regiao8":             "Leste 1",
			"registro":            "4041-0",
			"logradouro":          "RUA MARAGOJIPE",
			"numero":              "S/N",
			"bairro":              "VL FORMOSA",
			"referencia":          "TV RUA PRETORIA",
		},
	)
	return data
}
