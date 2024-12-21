package services_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanListABoilers(t *testing.T) {
	models, err := boilerListService.Execute()

	require.NoError(t, err)
	require.NotNil(t, models)
	require.True(t, boilerMockSpyRepository.(*MockSpyBoilerRepository).ListHaveBeenCalled())
}
