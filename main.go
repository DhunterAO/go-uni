package main

import (
    "fmt"
    "log"
	
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
	
	"./contracts/router"
)

func main() {
	client, err := ethclient.Dial("/data/polygon/.bor/bor.ipc")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")

	const uniV3RouterAddr = common.HexToAddress("0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45")

	instance, err := router.NewRouter(uniV3RouterAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	fatory, err := instance.Factory(&bind.CallOpts{}, uniV3RouterAddr)
	fmt.Printf("wei: %s\n", bal)
}
