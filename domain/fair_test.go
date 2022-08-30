package domain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_should_return_errors_when_required_data_is_empty(t *testing.T) {
	fair := Fair{}
	errs := fair.Validate()
	assertRequiredFields(t, errs.ValidationErrors)
}

func Test_should_return_errors_when_max_length_is_greater_than_allowed(t *testing.T) {
	// TODO: Use a fake Generator
	fair := Fair{
		Longitude:         strings.Repeat("x", 11),
		Latitude:          strings.Repeat("x", 11),
		Sector:            strings.Repeat("x", 16),
		Area:              strings.Repeat("x", 14),
		DistrictCode:      strings.Repeat("x", 10),
		District:          strings.Repeat("x", 19),
		SubprefectureCode: strings.Repeat("x", 3),
		SubprefectureName: strings.Repeat("x", 26),
		Region5:           strings.Repeat("x", 7),
		Region8:           strings.Repeat("x", 8),
		FairName:          strings.Repeat("x", 31),
		Registry:          strings.Repeat("x", 7),
		Address:           strings.Repeat("x", 35),
		Number:            strings.Repeat("x", 6),
		Neighborhood:      strings.Repeat("x", 21),
		Reference:         strings.Repeat("x", 51),
	}
	errs := fair.Validate()
	assertMaxLengthFields(t, errs.ValidationErrors)
}

func assertRequiredFields(t *testing.T, errors map[string][]string) {
	msg := "Can't be blank!"
	assert.Equal(t, errors["longitude"][0], msg)
	assert.Equal(t, errors["latitude"][0], msg)
	assert.Equal(t, errors["setorCensitario"][0], msg)
	assert.Equal(t, errors["areaPonderacao"][0], msg)
	assert.Equal(t, errors["codigoDistrito"][0], msg)
	assert.Equal(t, errors["distrito"][0], msg)
	assert.Equal(t, errors["codigoSubprefeitura"][0], msg)
	assert.Equal(t, errors["nomeSubPrefeitura"][0], msg)
	assert.Equal(t, errors["regiao5"][0], msg)
	assert.Equal(t, errors["regiao8"][0], msg)
	assert.Equal(t, errors["nomeFeira"][0], msg)
	assert.Equal(t, errors["registro"][0], msg)
	assert.Equal(t, errors["logradouro"][0], msg)
	assert.Equal(t, errors["numero"][0], msg)
	assert.Equal(t, errors["bairro"][0], msg)
}

func assertMaxLengthFields(t *testing.T, errors map[string][]string) {
	assert.Equal(t, errors["longitude"][0], "Max length is 10.")
	assert.Equal(t, errors["latitude"][0], "Max length is 10.")
	assert.Equal(t, errors["setorCensitario"][0], "Max length is 15.")
	assert.Equal(t, errors["areaPonderacao"][0], "Max length is 13.")
	assert.Equal(t, errors["codigoDistrito"][0], "Max length is 9.")
	assert.Equal(t, errors["distrito"][0], "Max length is 18.")
	assert.Equal(t, errors["codigoSubprefeitura"][0], "Max length is 2.")
	assert.Equal(t, errors["nomeSubPrefeitura"][0], "Max length is 25.")
	assert.Equal(t, errors["regiao5"][0], "Max length is 6.")
	assert.Equal(t, errors["regiao8"][0], "Max length is 7.")
	assert.Equal(t, errors["nomeFeira"][0], "Max length is 30.")
	assert.Equal(t, errors["registro"][0], "Max length is 6.")
	assert.Equal(t, errors["logradouro"][0], "Max length is 34.")
	assert.Equal(t, errors["numero"][0], "Max length is 5.")
	assert.Equal(t, errors["bairro"][0], "Max length is 20.")
	assert.Equal(t, errors["referencia"][0], "Max length is 50.")
}
