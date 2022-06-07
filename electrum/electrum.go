package electrum

import (
	"context"
	"log"
	"time"
)

func GetBitcoinConfirmedBalance(ctx context.Context, btcAddress string) (float64, error) {
	scripthash, err := BtcAddressToElectrumxFormat(btcAddress)
	if err != nil {
		return 0, err
	}
	// Establishing a new SSL connection to an ElectrumX server
	server := NewServer(&ServerOptions{ConnTimeout: 15 * time.Second, ReqTimeout: 15 * time.Second})
	if err := server.ConnectTCP("bch.imaginary.cash:50001"); err != nil {
		log.Fatal(err)
	}

	// Asking the server for the balance of address 1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa
	// We must use scripthash of the address now as explained in ElectrumX docs
	balance, err := server.GetBalance(scripthash)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Scripthash:   %+v", scripthash)
	log.Printf("Address confirmed balance:   %+v", balance.Confirmed)
	log.Printf("Address unconfirmed balance: %+v", balance.Unconfirmed)
	return balance.Confirmed, nil
}
