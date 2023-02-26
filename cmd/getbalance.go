package cmd

import (
	"fmt"
	"log"

	"github.com/akhil/proper_blockchain/core"

	utils "github.com/akhil/proper_blockchain/utils"
	blockchain "github.com/akhil/proper_blockchain/core/blockchain"
	wallets "github.com/akhil/proper_blockchain/core/wallets"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !wallets.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{Blockchain: bc}
	defer bc.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
