package main

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (u *Util) cmdLogid(args []string) error {
	if len(args) >= 1 {
		switch args[0] {
		case "deploy":
			minEther := new(big.Int).SetUint64(1700000000000000000)
			owner := u.ma.Address
			if len(args) == 2 {
				minEther, _ = ethToWei(args[1])
			}
			if len(args) == 3 {
				owner = common.HexToAddress(args[2])
			}
			addr, tx, lg, err := u.DeployLogid(owner, minEther)
			u.logid = lg
			if err != nil {
				u.logf("%s", err)
				return nil
			}
			u.logf("%s", tx)
			u.logf("SmartContract Address: %s", addr.String())
			u.AddPending(tx)

		case "minether":
			mineth, _ := u.logid.MinEther(nil)
			u.logf("minEther: %d", mineth)

		case "ban":
			user := common.Address{}
			if len(args) == 2 {
				user = common.HexToAddress(args[1])
			} else {
				u.logf("User?")
				return nil
			}
			tx, err := u.BanUser(user)
			if err != nil {
				u.logf("%s", err)
				return nil
			}
			u.AddPending(tx)
			u.logf("%s", tx)
			u.logf("%s", tx.Hash().String())
		}
	}
	return nil
}

func (u *Util) BanUser(user common.Address) (*types.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pk, err := u.privateKey()
	if err != nil {
		return nil, err
	}
	opts := bind.NewKeyedTransactor(pk)

	opts.GasPrice = u.gasPrice
	opts.GasLimit = u.gasLimit
	opts.Context = ctx
	opts.From = u.ma.Address

	return u.logid.Ban(opts, user)
}
