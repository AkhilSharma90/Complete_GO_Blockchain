package cmd

import (
	"fmt"
	"log"

	"github.com/akhil/proper_blockchain/core"
	blockchain "github.com/akhil/proper_blockchain/core/blockchain"
	wallets "github.com/akhil/proper_blockchain/core/wallets"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !wallets.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.CreateBlockchain(address, nodeID)
	defer bc.Close()

	UTXOSet := core.UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
