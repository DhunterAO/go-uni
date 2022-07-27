package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/DhunterAO/go-uni/contracts/curve/vyper"
	"github.com/DhunterAO/go-uni/contracts/univ3/router"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	headers := make(chan *types.Header)

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	amount := int64(1000000000)

	uniV3RouterAddr := common.HexToAddress("0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45")
	curveVyperAddr := common.HexToAddress("0x1d8b86e3d88cdb2d34688e87e72f388cb541b7c8")

	account := common.HexToAddress("0xF5c12b5b6aB5aFbD87a5BE34f1f7a6473b7eAb0F")

	usdc := common.HexToAddress("0x2791bca1f2de4661ed88a30c99a7a9449aa84174")
	weth := common.HexToAddress("0x7ceb23fd6bc0add59e62ac25578270cff1b9f619")
	exactInputParams := router.IV3SwapRouterExactInputSingleParams{
		TokenIn:           usdc,
		TokenOut:          weth,
		Fee:               big.NewInt(500),
		Recipient:         account,
		AmountIn:          big.NewInt(amount),
		AmountOutMinimum:  big.NewInt(0),
		SqrtPriceLimitX96: big.NewInt(0),
	}

	exactOutputParams := router.IV3SwapRouterExactOutputSingleParams{
		TokenIn:           weth,
		TokenOut:          usdc,
		Fee:               big.NewInt(500),
		Recipient:         account,
		AmountOut:         big.NewInt(amount),
		AmountInMaximum:   big.NewInt(1000000000000000000),
		SqrtPriceLimitX96: big.NewInt(0),
	}

	routerInstance, err := router.NewRouterCaller(uniV3RouterAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	vyperInstance, err := vyper.NewVyperCaller(curveVyperAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Number, time.Now()) // pointer to event log
			amountOut, err := routerInstance.ExactInputSingle(&bind.CallOpts{From: account}, exactInputParams)
			if err != nil {
				log.Fatal("error:", err)
			}
			fmt.Printf("univ3 amountOut: %s\n", amountOut)

			amountIn, err := routerInstance.ExactOutputSingle(&bind.CallOpts{From: account}, exactOutputParams)
			if err != nil {
				log.Fatal("error:", err)
			}
			fmt.Printf("univ3 amountIn: %s\n", amountIn)
			fmt.Println(time.Now())

			amountOut, err = vyperInstance.GetDyUnderlying(&bind.CallOpts{From: account}, big.NewInt(1), big.NewInt(4), big.NewInt(amount))
			if err != nil {
				log.Fatal("error:", err)
			}
			fmt.Printf("curve amountOut: %s\n", amountOut)

			// amountIn, err = vyperInstance.GetDyUnderlying(&bind.CallOpts{From: account}, big.NewInt(4), big.NewInt(1), big.NewInt(amount))
			// if err != nil {
			// 	log.Fatal("error:", err)
			// }
			// fmt.Printf("curve amountIn: %s\n", amountIn)
		}
	}
}
