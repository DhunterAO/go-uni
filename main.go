package main

import (
	"fmt"
	"log"
	"math/big"

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

	usdc := common.HexToAddress("0x2791bca1f2de4661ed88a30c99a7a9449aa84174")
	weth := common.HexToAddress("0x7ceb23fd6bc0add59e62ac25578270cff1b9f619")
	fee := big.NewInt(500)
	receiptor := common.HexToAddress("0xF5c12b5b6aB5aFbD87a5BE34f1f7a6473b7eAb0F")
	amountIn := big.NewInt(100000000000)
	amountOutMin := big.NewInt(0)
	sqrtPrice := big.NewInt(0)
	params := router.IV3SwapRouterExactInputSingleParams{
		TokenIn:           usdc,
		TokenOut:          weth,
		Fee:               fee,
		Recipient:         receiptor,
		AmountIn:          amountIn,
		AmountOutMinimum:  amountOutMin,
		SqrtPriceLimitX96: sqrtPrice,
	}
	amountOut, err := instance.ExactInputSingle(&bind.CallOpts{}, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("amountOut: %s\n", amountOut)
}
