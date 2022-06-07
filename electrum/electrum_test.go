package electrum

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestElectrum_GetBalance(t *testing.T) {
	DebugMode = true
	b, err := GetBitcoinConfirmedBalance(context.Background(), "bitcoin.aranguren.org:50001", "bc1qjsn4d8cqnvhsw4q6m456wxg54lanu4jnq2lypc")
	log.Fatal(b)
	b, err = GetBitcoinConfirmedBalance(context.Background(), "btc.electroncash.dk:60001", "bc1qjsn4d8cqnvhsw4q6m456wxg54lanu4jnq2lypc")
	require.NoError(t, err)
	b, err = GetBitcoinConfirmedBalance(context.Background(), "fullnode.titanconnect.ca:50001", "bc1qjsn4d8cqnvhsw4q6m456wxg54lanu4jnq2lypc")
	require.NoError(t, err)
	assert.True(t, b > 0)
}
