package main

import (
	"log"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (u *Util) AddPending(txs *types.Transaction) {
	u.pending = txs
}

func (u *Util) Pending() *types.Transaction {
	return u.pending
}

func (u *Util) HeadWriter() {
	name := "HeadWriter"
	ch := make(chan *types.Header)
	s, err := u.d.SubscribeNewHead(ch)
	if err != nil {
		log.Printf("HeadWriter err: %s\n", err)
		return
	}
	defer s.Unsubscribe()
	for {
		select {
		case <-u.done:
			u.logf("%s ended.", name)
			return
		case h := <-ch:
			if u.pending == nil {
				continue
			}
			block, _ := u.d.BlockNumber(h.Number)
			for _, tx := range block.Transactions() {
				if reflect.DeepEqual(tx.Hash(), u.Pending().Hash()) {
					if tx == nil {
						u.logf("Nil transaction")
						continue
					}
					r, _ := u.Receipt(tx.Hash())
					u.logf("status: %d, cumulativeGasUsed: %d, gasUsed: %d", r.Status, r.CumulativeGasUsed, r.GasUsed)
					if !reflect.DeepEqual(r.ContractAddress, common.Address{}) {
						u.logf("Contract Address: %s", r.ContractAddress.String())
					}
					u.pending = nil
					continue
				}
			}
		}
	}
}
