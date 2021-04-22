package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBitcoinPrice(t *testing.T) {
	res, err := GetCryptoPrice("bitcoin")
	require.NoError(t, err)
	assert.Equal(t, "bitcoin", len(res.ID))
	assert.Equal(t, "btc", res.Symbol)
	assert.Equal(t, "Bitcoin", res.Name)
	// TODO: make this test more when there struct is all filled out
}
