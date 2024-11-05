package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Arbitrum Sepolia
	client, err := ethclient.Dial("https://sepolia-rollup.arbitrum.io/rpc")
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress("0x543FF5baFD7fcD727711900A48F040B4405D4618")

	// read array length in slot 0
	lengthSlot := common.Hash{}
	length, err := client.StorageAt(context.Background(), contractAddr, lengthSlot, nil)
	if err != nil {
		log.Fatal(err)
	}

	arrayLength := new(big.Int).SetBytes(length).Uint64()
	fmt.Println("arrayLength:", arrayLength)

	// calculate keccak256 hash of array start position
	arraySlot := common.BigToHash(big.NewInt(0))
	baseSlot := common.BytesToHash(crypto.Keccak256(arraySlot.Bytes()))

	// read each array element
	for i := uint64(0); i < arrayLength; i++ {
		// calculate current element storage position
		currentSlot := new(big.Int).Add(new(big.Int).SetBytes(baseSlot.Bytes()), big.NewInt(int64(i*2)))

		// read first slot (user + startTime)
		data, err := client.StorageAt(context.Background(), contractAddr, common.BigToHash(currentSlot), nil)
		if err != nil {
			log.Fatal(err)
		}

		// read second slot (amount)
		nextSlot := new(big.Int).Add(currentSlot, big.NewInt(1))
		amountData, err := client.StorageAt(context.Background(), contractAddr, common.BigToHash(nextSlot), nil)
		if err != nil {
			log.Fatal(err)
		}

		// parse data (reverse order)
		/*
			struct LockInfo {
				address user;      // 20 bytes
				uint64 startTime;  // 8 bytes
				uint256 amount;    // 32 bytes
			}
		*/
		user := common.BytesToAddress(data[12:32])     // last 20 bytes is user address
		startTime := new(big.Int).SetBytes(data[4:12]) // middle 8 bytes is startTime
		amount := new(big.Int).SetBytes(amountData)    // 32 bytes is amount

		fmt.Printf("locks[%d]:\n", i)
		fmt.Printf("  user: %s\n", user.Hex())
		fmt.Printf("  startTime: %s (%d)\n",
			time.Unix(startTime.Int64(), 0).Format("2006-01-02 15:04:05"),
			startTime,
		)
		fmt.Printf("  amount: %.2f ETH\n",
			new(big.Float).Quo(
				new(big.Float).SetInt(amount),
				new(big.Float).SetFloat64(1e18),
			),
		)

		fmt.Printf("  Debug - slot1 data: %x\n", data)
		fmt.Printf("  	Offset: 0 - padding: %x\n", data[0:4])    // first 4 bytes (padding)
		fmt.Printf("  	Offset: 4 - startTime: %x\n", data[4:12]) // middle 8 bytes
		fmt.Printf("  	Offset: 12 - user: %x\n", data[12:32])    // last 20 bytes
		fmt.Printf("  Debug - slot2 data (amount): %x\n", amountData)
		fmt.Println("----------------------------------------")
	}
}
