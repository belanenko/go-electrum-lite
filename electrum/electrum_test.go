package electrum

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestXxx(t *testing.T) {
	DebugMode = true
	b, err := GetBitcoinConfirmedBalance(context.Background(), "bch.imaginary.cash:50001", "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")
	require.NoError(t, err)
	assert.True(t, b > 0)
}
