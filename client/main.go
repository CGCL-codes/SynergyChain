package main

import (
	"ShareChain-Client/block"
	"ShareChain-Client/front"
	"log"
	"time"
)

var URL = "http://127.0.0.1:7999/GetSharingData"

func main() {
	// 初始化区块数据库
	block.InitBlockChainDB()
	// 开启同步区块头线程
	go sync()

	go func() {
		log.Fatal(front.HttpApiRun("9000"))
	}()

	c := time.Tick(3 * time.Second)
	for {
		<- c

	}

}

func sync() {
	c := time.Tick(3 * time.Second)
	for {
		<- c
		block.SyncBlockHead()
	}
}