package services_test

import (
	"errors"
	"markitos-golang-service-boilerplate/internal/domain"
	"markitos-golang-service-boilerplate/internal/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCreateABoiler(t *testing.T) {
	var request services.BoilerCreateRequest = services.BoilerCreateRequest{
		Message: VALID_MESSAGE,
	}
	response, err := boilerCreateService.Execute(request)

	require.NoError(t, err)
	require.True(t, domain.IsUUIDv4(response.Id))
	require.Equal(t, VALID_MESSAGE, response.Message)
	require.True(t, boilerMockSpyRepository.(*MockSpyBoilerRepository).CreateHaveBeenCalledWithMessage(response))
}

func TestCantCreateABoilerWithEmptyMessage(t *testing.T) {
	var request services.BoilerCreateRequest = services.BoilerCreateRequest{
		Message: "",
	}

	response, err := boilerCreateService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.BoilerInvalidMessageError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}
