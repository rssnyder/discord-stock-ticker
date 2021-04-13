package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAMCPrice(t *testing.T) {
	res, err := GetStockPrice("AMC")
	require.NoError(t, err)
	assert.Equal(t, 1, len(res.QuoteSummary.Results))
	assert.Equal(t, "AMC", res.QuoteSummary.Results[0].Price.Symbol)
	assert.Equal(t, "USD", res.QuoteSummary.Results[0].Price.Currency)
}
