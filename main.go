package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/sunshuo3436/solidity_go/ethwallet"
)

func main() {
	// Connecting to an Ethereum node
	client, err := ethclient.Dial("https://holesky.infura.io/v3/c0fd902e8c32475f85909278c7352314")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddress := common.HexToAddress("0xcefc94f685afbcf773f5d7305ceddd9a0b4c9e21")
	privateKeyHex := "661260045b66da8aa4b5e6ac127f07021bee203136ee769fcaebb97c2eeb3708"

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// Read ABI file
	abiBytes, err := ioutil.ReadFile("ethwallet.abi")
	if err != nil {
		log.Fatalf("Failed to read ABI file: %v", err)
	}

	// Parse ABI
	contractAbi, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	// Create a new contract instancez
	contractInstance, err := ethwallet.NewEthwallet(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a smart contract: %v", err)
	}

	// Create a new transactor
	auth := bind.NewKeyedTransactor(privateKey)

	// Get balance
	balance, err := contractInstance.GetBalance(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}
	fmt.Printf("Balance: %d\n", balance)

	// Send Transaction
	amount := big.NewInt(1000000000000000000)
	tx, err := contractInstance.Withdraw(auth, amount)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}
	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())

	// Signing transactions
	txData, err := contractAbi.Pack("Withdraw", amount)
	if err != nil {
		log.Fatalf("Failed to pack transaction data: %v", err)
	}
	txHash := crypto.Keccak256(txData)
	signature, err := crypto.Sign(txHash, privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}
	fmt.Printf("Signature: %x\n", signature)
}
