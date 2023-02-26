package cmd

import (
	"fmt"
	"log"

	"github.com/akhil/proper_blockchain/core"
	wallets "github.com/akhil/proper_blockchain/core/wallets"
)

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)
	if len(minerAddress) > 0 {
		if wallets.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}
	core.StartServer(nodeID, minerAddress)
}
