package services_test

import (
	"errors"
	"markitos-service-boilerplate/internal/domain"
	"markitos-service-boilerplate/internal/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanGetABoiler(t *testing.T) {
	response, _ := boilerCreateService.Execute(services.BoilerCreateRequest{
		Message: VALID_MESSAGE,
	})
	require.True(t, domain.IsUUIDv4(response.Id))

	model, err := boilerOneService.Execute(services.NewBoilerOneRequest(response.Id))
	require.NoError(t, err)
	require.True(t, domain.IsUUIDv4(model.Id))
	require.True(t, boilerMockSpyRepository.(*MockSpyBoilerRepository).OneHaveBeenCalledWithMessage(model))
}

func TestCantGetOneBoilerWithEmptyId(t *testing.T) {
	var request services.BoilerOneRequest = services.BoilerOneRequest{
		Id: "",
	}

	response, err := boilerOneService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.BoilerInvalidIdError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}
