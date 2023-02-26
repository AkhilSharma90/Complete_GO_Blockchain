package cmd

import (
	"fmt"
	"log"

	core "github.com/akhil/proper_blockchain/core"
	blockchain "github.com/akhil/proper_blockchain/core/blockchain"
	transactions "github.com/akhil/proper_blockchain/core/transactions"
	wallets "github.com/akhil/proper_blockchain/core/wallets"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !wallets.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !wallets.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{Blockchain: bc}
	defer bc.Close()

	wallets, err := wallets.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := core.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := transactions.NewCoinbaseTX(from, "")
		txs := []*transactions.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		core.SendTx(core.KnownNodes[0], tx)
	}

	fmt.Println("Success!")
}
