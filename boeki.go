package main

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// control trading from Boeki smart contract
func (u *Util) cmdBoeki(args []string) error {
	if len(args) >= 1 {
		switch args[0] {
		case "deposit":
			if len(args) == 2 {
				amount, _ := ethToWei(args[1])
				tx, err := u.DepositBoeki(amount)
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				u.logf("%s", tx)
				u.AddPending(tx)

			}
		case "deploy":
			addr, tx, err := u.DeployBoeki()
			if err != nil {
				u.logf("%s", err)
				return nil
			}
			u.logf("%s", tx)
			u.logf("SmartContract Address: %s", addr.String())
			u.AddPending(tx)
		}
	}
	return nil
}

func (u *Util) DepositBoeki(amount *big.Int) (*types.Transaction, error) {
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

	// deposit Amount
	opts.Value = amount

	return u.boeki.Deposit(opts)
}

func (u *Util) DeployBoeki() (common.Address, *types.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pk, err := u.privateKey()
	if err != nil {
		return common.Address{}, nil, err
	}
	opts := bind.NewKeyedTransactor(pk)

	opts.GasPrice = u.gasPrice
	opts.GasLimit = u.gasLimit
	opts.Context = ctx
	opts.From = u.ma.Address

	addr, tx, _, err := DeployBoekismart(opts, u.c, DELTA_CONTRACT, u.ma.Address)
	return addr, tx, err
}
