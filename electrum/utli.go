package electrum

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

func main() {
	log.Fatal(btcutil.AmountBTC)
}

func AddressToScriptHex(addressStr string) (string, error) {
	address, err := btcutil.DecodeAddress(addressStr, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}
	script, err := txscript.PayToAddrScript(address)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", script), nil
}

func scriptHexToSHA256(scriptHex string) ([]byte, error) {
	sha256 := sha256.New()
	s, err := hex.DecodeString(scriptHex)
	if err != nil {
		return nil, err
	}
	sha256.Write(s)
	return sha256.Sum(nil), err
}

func ScriptHexToSHA256Reversed(scriptHex string) (string, error) {
	var sb strings.Builder
	sha, err := scriptHexToSHA256(scriptHex)
	if err != nil {
		return "", err
	}
	for i := len(sha) - 1; i != -1; i-- {
		sb.WriteByte(sha[i])
	}

	return fmt.Sprintf("%x", sb.String()), nil
}

func BtcAddressToElectrumxFormat(address string) (string, error) {
	scriptHex, err := AddressToScriptHex(address)
	if err != nil {
		return "", err
	}
	result, err := ScriptHexToSHA256Reversed(scriptHex)
	return result, err
}
