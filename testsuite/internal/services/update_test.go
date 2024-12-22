package services_test

import (
	"errors"
	"markitos-service-boilerplate/internal/domain"
	"markitos-service-boilerplate/internal/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanUpdateABoiler(t *testing.T) {
	createdBoiler, err := boilerCreateService.Execute(services.BoilerCreateRequest{
		Message: VALID_MESSAGE,
	})
	require.NoError(t, err)

	model, err := boilerUpdateService.Execute(services.BoilerUpdateRequest{
		Id:      createdBoiler.Id,
		Message: createdBoiler.Message + " updated",
	})

	require.NoError(t, err)
	require.Equal(t, createdBoiler.Id, model.Id)
	require.NotEqual(t, createdBoiler.Message, model.Message)
	require.Equal(t, createdBoiler.Message+" updated", model.Message)

	require.True(t, boilerMockSpyRepository.(*MockSpyBoilerRepository).UpdateHaveBeenCalledWithMessage(model))
}

func TestCantUpdatOneBoilerWithEmptyMessage(t *testing.T) {
	var request services.BoilerUpdateRequest = services.BoilerUpdateRequest{
		Id:      VALID_UUIDV4,
		Message: "",
	}

	response, err := boilerUpdateService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.BoilerInvalidMessageError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}

func TestCantUpdatOneBoilerWithEmptyId(t *testing.T) {
	var request services.BoilerUpdateRequest = services.BoilerUpdateRequest{
		Id:      "",
		Message: VALID_MESSAGE,
	}

	response, err := boilerUpdateService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.BoilerInvalidIdError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}

func TestCantUpdatOneBoilerWithInvalidId(t *testing.T) {
	var request services.BoilerUpdateRequest = services.BoilerUpdateRequest{
		Id:      "invalid-id",
		Message: VALID_MESSAGE,
	}

	response, err := boilerUpdateService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.BoilerInvalidIdFormatError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}
