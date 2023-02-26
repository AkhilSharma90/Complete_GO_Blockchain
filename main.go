package main

import (
	cmd "github.com/akhil/proper_blockchain/cmd"
)

func main() {
	cmd := cmd.CLI{}
	cmd.Run()
}
