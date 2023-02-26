package cmd

import (
	"fmt"

	"github.com/akhil/proper_blockchain/core"
	blockchain "github.com/akhil/proper_blockchain/core/blockchain"
)

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
