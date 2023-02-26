package cmd

import (
	"fmt"

	wallets "github.com/akhil/proper_blockchain/core/wallets"
)

func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := wallets.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your new address: %s\n", address)
}
