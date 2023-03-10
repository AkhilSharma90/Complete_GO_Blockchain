# A Complete Blockchain Using GO

Adbanced blockchain implementation in Go.
Using all sorts of concepts like Wallets, Nodes, UTXOs, transactions, block iterators etc.

## Usage

```bash
Usage:
    `createblockchain -address <ADDRESS>` - Create a blockchain and send genesis block reward to ADDRESS
    `createwallet` - Generates a new key-pair and saves it into the wallet file
    `getbalance -address <ADDRESS>` - Get balance of ADDRESS
    `listaddresses` - Lists all addresses from the wallet file
    `printchain` - Print all the blocks of the blockchain
    `reindexutxo` - Rebuilds the UTXO set
    `send -from <FROM> -to <TO> -amount <AMOUNT> -mine` - Send AMOUNT of coins from FROM address to TO. Mine on the same node, when -mine is set.
    `startnode -miner <ADDRESS>` - Start a node with ID specified in NODE_ID env. var. -miner enables mining
```
