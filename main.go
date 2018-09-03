package main

import (
	"context"
	"fmt"
	"go-eth/eth"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	client, err := eth.Connect("http://localhost:8545")

	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	blockNumber, err := client.GetBlockNumber(context.TODO())

	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	fmt.Printf("Latest block number: %s\n", blockNumber.String())
	// fmt.Println(client.EthClient.BlockByNumber(context.TODO(), blockNumber))
	// fmt.Println(client.EthClient.HeaderByNumber(context.TODO(), big.NewInt(0)))
	from := common.HexToAddress("30b82c8694b59695d78f33a7ba1c2a55dfa618d5")
	to := common.HexToAddress("5e0f92917d632f7cdb7564a67644ca45430b524c")
	// nonce, err := client.EthClient.NonceAt(context.TODO(), from, nil)
	amount := big.NewInt(1)
	gasLimit := big.NewInt(90000)
	gasPrice := big.NewInt(0)
	data := []byte{}
	message := eth.NewMessage(from, &to, amount, gasLimit, gasPrice, data)
	var txHash common.Hash
	txHash, err = client.SendTransaction(context.TODO(), &message)

	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	fmt.Printf("Transaction has been sent, transaction hash: %s\n", txHash.String())

	tx, isPending, _ := client.EthClient.TransactionByHash(context.TODO(), txHash)
	fmt.Println(tx.Nonce())
	fmt.Println(isPending)

	time.Sleep(10 * time.Second)

	receipt, _ := client.EthClient.TransactionReceipt(context.TODO(), txHash)
	fmt.Println(receipt.Status)
}
