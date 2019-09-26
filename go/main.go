package main

import (
	"fmt"
	"lang.yottadb.com/go/yottadb"
	"log"
)

const (
	Global = "^acc"
	Id = "10001"
)

func main() {
	defer yottadb.Exit()

	var errstr yottadb.BufferT

	tptoken := yottadb.NOTTP
	err := yottadb.SetValE(tptoken, &errstr, "300.00", Global, []string {Id})
	if err != nil {
		log.Fatalf("error during save %v, errstr %v", err.Error(), errstr)
	}

	val, err := yottadb.ValE(tptoken, &errstr, Global, []string{Id})
	if err != nil {
		log.Fatalf("err=%v, errstr=%s", err.Error(), errstr)
	}
	fmt.Printf("===>%s\n",val)
}
