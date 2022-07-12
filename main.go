package main

import (
    "fmt"
    "log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
	
	"github.com/DhunterAO/go-uni/contracts/router"
)

func main() {
	client, err := ethclient.Dial("/data/polygon/.bor/bor.ipc")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")

	uniV3RouterAddr := common.HexToAddress("0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45")

	instance, err := router.NewRouterCaller(uniV3RouterAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	fatory, err := instance.Factory(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("fatory: %s\n", fatory)
	
	params := router.IV3SwapRouterExactInputSingleParams{

	}
	amountOut, err := instance.ExactInputSingle()
}
