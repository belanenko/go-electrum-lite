package electrum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBtcutil_AddressToScriptHex(t *testing.T) {
	address := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
	expected := "76a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1888ac"

	actual, _ := AddressToScriptHex(address)

	assert.Equal(t, expected, actual)
}

func TestBtcutil_scriptHexToSHA256(t *testing.T) {
	scripthash := "76a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1888ac"
	expected := "6191c3b590bfcfa0475e877c302da1e323497acf3b42c08d8fa28e364edf018b"

	actualBytes, _ := scriptHexToSHA256(scripthash)
	actual := fmt.Sprintf("%x", actualBytes)

	assert.Equal(t, expected, actual)
}

func TestBtcutil_ScriptHexToSHA256Reversed(t *testing.T) {
	scripthash := "76a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1888ac"
	expected := "8b01df4e368ea28f8dc0423bcf7a4923e3a12d307c875e47a0cfbf90b5c39161"

	actual, _ := ScriptHexToSHA256Reversed(scripthash)

	assert.Equal(t, expected, actual)
}

func TestBtcutil_BtcAddressToElectrumxFormat(t *testing.T) {
	address := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
	expected := "8b01df4e368ea28f8dc0423bcf7a4923e3a12d307c875e47a0cfbf90b5c39161"

	actual, _ := BtcAddressToElectrumxFormat(address)

	assert.Equal(t, expected, actual)
}
