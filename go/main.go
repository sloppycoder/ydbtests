package main

import (
	"fmt"
	"lang.yottadb.com/go/yottadb"
	"log"
)

const (
	Global = "^acc"
)

func main() {
	defer yottadb.Exit()

	var errstr yottadb.BufferT

	tptoken := yottadb.NOTTP
	for i := 1; i < 1000; i++ {
		id := fmt.Sprintf("%05d", i)
		err := yottadb.SetValE(tptoken, &errstr, "300.00", Global, []string{id})
		if err != nil {
			log.Fatalf("error during save %v, errstr %v", err.Error(), errstr)
		}

		val, err := yottadb.ValE(tptoken, &errstr, Global, []string{id})
		if err != nil {
			log.Fatalf("err during read %v, errstr %v", err.Error(), errstr)
		}
		fmt.Printf("===>%s\n", val)
	}
}
