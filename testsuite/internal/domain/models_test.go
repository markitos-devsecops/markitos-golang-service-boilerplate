package domain_test

import (
	"errors"
	"markitos-service-boilerplate/internal/domain"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateBoilerWithEmptyId(t *testing.T) {
	boiler, err := domain.NewBoiler("", "Hello, World!")

	var invalidIdErr *domain.BoilerInvalidIdError
	require.True(t, errors.As(err, &invalidIdErr))
	require.Equal(t, domain.NewBoilerInvalidIdError("").Error(), err.Error())
	require.True(t, strings.HasPrefix(err.Error(), domain.BOILER_INVALID_ID_PREFIX))
	require.Error(t, err)
	require.Empty(t, boiler)
}

func TestCreateBoilerWithEmptyMessage(t *testing.T) {
	boiler, err := domain.NewBoiler(VALID_UUIDV4, "")

	var invalidErr *domain.BoilerInvalidMessageError
	require.True(t, errors.As(err, &invalidErr))
	require.Equal(t, domain.NewBoilerInvalidMessageError("").Error(), err.Error())
	require.True(t, strings.HasPrefix(err.Error(), domain.BOILER_INVALID_MESSAGE_PREFIX))
	require.Error(t, err)
	require.Empty(t, boiler)
}
