package main

import (
	"context"
	"errors"
	"math/big"
	"time"

	"k.veselin.org/boeki/delta"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

// control trading from Boeki smart contract
func (u *Util) cmdBoeki(args []string) error {
	if len(args) >= 1 {
		switch args[0] {
		case "operator":
			if len(args) == 2 {
				tx, err := u.ChangeOperator(common.HexToAddress(args[1]))
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				u.AddPending(tx)
				u.logf("%s", tx)
				u.logf("%s", tx.Hash().String())
				return nil
			}
			operator, err := u.proxy.Operator(nil)
			if err != nil {
				u.logf("%s", err)
				return nil
			}

			u.logf("Current Boeki operator: %s", operator.String())
		case "sky":
			digit := new(big.Int).SetUint64(13)
			addr := CHISAI_TOKEN
			if len(args) >= 2 {
				digit, _ = new(big.Int).SetString(args[1], 10)
			}
			if len(args) >= 3 {
				addr = common.HexToAddress(args[2])
			}
			txs, err := u.Sky("0x69231804", digit, addr)
			if err != nil {
				u.logf("%s", err)
				return nil
			}
			u.AddPending(txs)
			u.logf("%s", txs)
			u.logf("%s", txs.Hash().String())

		case "tradeAmass":
			if u.orderSell != nil && u.orderBuy != nil {
				tx, err := u.BoekiCompactSellBuy("0x78a31e77", nil, nil, u.orderSell, u.orderBuy)
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				u.AddPending(tx)
				u.logf("%s", tx)
				u.logf("%s", tx.Hash().String())

			} else {
				u.logf("make one Sell and one Buy order first")
			}
		case "tradeFake":
			o := u.order
			side, _ := o.GetSide()
			hexfunc := ""
			switch side {
			case "sell":
				hexfunc = "0x641fb55d"
			case "buy":
				hexfunc = "0x8b3dcb18"
			}
			amount := o.GetAmountGet()
			if len(args) == 2 {
				amount, _ = ethToWei(args[1])
			}
			u.logf("%s", amount.String())
			txs, err := u.BoekiCompactTrade(hexfunc, amount, o)
			if err != nil {
				u.logf("%s", err)
				return nil
			}
			u.AddPending(txs)
			u.logf("%s", txs)
			u.logf("%s", txs.Hash().String())

		case "tradeC":
			o := u.order
			side, _ := o.GetSide()
			hexfunc := ""
			switch side {
			case "sell":
				hexfunc = "0x641fb55d"
			case "buy":
				hexfunc = "0xb1d64d4b"
				//hexfunc = "0x00000007"
			}
			amount := o.GetAmountGet()
			if len(args) == 2 {
				amount, _ = ethToWei(args[1])
			}
			u.logf("%s", amount.String())
			txs, err := u.BoekiCompactTrade(hexfunc, amount, o)
			if err != nil {
				u.logf("%s", err)
				return nil
			}
			u.AddPending(txs)
			u.logf("%s", txs)
			u.logf("%s", txs.Hash().String())
		case "test":
			if len(args) >= 2 {
				switch args[1] {
				case "selluint":
					digit := new(big.Int).SetUint64(13)
					addr := CHISAI_TOKEN
					if len(args) == 3 {
						digit, _ = new(big.Int).SetString(args[2], 10)
					}
					if len(args) == 4 {
						addr = common.HexToAddress(args[3])
					}
					tx, err := u.SellUint(digit, addr, "0xfac8636c")
					if err != nil {
						u.logf("%s", err)
						return nil
					}
					u.logf("%s", tx)
					u.AddPending(tx)
				case "sellhigh":
					tx, err := u.SellHigh(new(big.Int).SetUint64(5101650), CHISAI_TOKEN, "0x00448366")
					if err != nil {
						u.logf("%s", err)
						return nil
					}
					u.logf("%s", tx)
					u.AddPending(tx)
				}
			}
		case "withdraw":
			u.SetGasLimit(56000)
			u.logf("gasLimit set to 56 000 for this operation.")

			if len(args) == 2 {
				amount, err := ethToWei(args[1])
				if err != nil {
					u.logf("%s", err)
					return nil
				}

				tx, err := u.BoekiCompactWithdraw(amount)
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				u.logf("%s", tx)
				u.AddPending(tx)
			}
		case "deposit":
			u.SetGasLimit(56000)
			u.logf("gasLimit set to 56 000 for this operation.")

			if len(args) == 3 {
				switch args[1] {
				case "proxy":
					amount, _ := ethToWei(args[2])
					tx, err := u.DepositBoekiProxy(amount)
					if err != nil {
						u.logf("%s", err)
						return nil
					}
					u.logf("%s", tx)
					u.AddPending(tx)

				case "normal":
					amount, _ := ethToWei(args[2])
					tx, err := u.DepositBoeki(amount)
					if err != nil {
						u.logf("%s", err)
						return nil
					}
					u.logf("%s", tx)
					u.AddPending(tx)
				}

			}
		case "deploy":
			bot := u.ma.Address
			if len(args) == 2 {
				if args[1] == "proxy" {

					addr, tx, bp, err := u.DeployBoekiProxy()
					u.proxy = bp
					if err != nil {
						u.logf("%s", err)
						return nil
					}
					u.logf("%s", tx)
					u.logf("(Proxy) SmartContract Address: %s", addr.String())
					u.AddPending(tx)
					return nil
				} else {
					bot = common.HexToAddress(args[1])
				}
			}
			// useProxy = false
			addr, tx, bs, err := u.DeployBoeki(DELTA_CONTRACT, bot, false)
			u.boeki = bs
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

func (u *Util) unpack(width int, b []byte) []byte {
	blen := len(b)
	if blen > width {
		return make([]byte, 0)
	}
	padding := width - blen
	out := make([]byte, padding)
	out = append(out, b...)
	return out
}

func (u *Util) SellUint(expires *big.Int, addr common.Address, hexfunc string) (*types.Transaction, error) {
	//TODO

	data := make([]byte, 0)

	method, err := hexutil.Decode(hexfunc)
	if err != nil {
		return nil, err
	}

	bexp := u.unpack(12, expires.Bytes())
	if len(bexp) != 12 {
		return nil, errors.New("[]byte wrong len")
	}

	u.logf("%+v %x", bexp, bexp)

	baddr := addr.Bytes()

	u.logf("m: %d, e: %d, a: %d", len(method), len(bexp), len(baddr))

	data = append(data, method...)
	data = append(data, bexp...)
	data = append(data, baddr...)

	u.logf("input: %x", data)

	nonce := u.Nonce()
	tx := types.NewTransaction(nonce, BOEKI_CONTRACT, nil, u.GasLimit(), u.GasPrice(), data)

	stx, err := u.ks.SignTx(*u.ma, tx, u.NetworkID())
	if err != nil {
		return nil, err
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
		defer cancel()
		tx = stx
		err = u.c.SendTransaction(ctx, tx)
		if err != nil {
			return nil, err
		}
		return tx, nil
	}

}

func (u *Util) Sky(hexfunc string, digit *big.Int, addr common.Address) (*types.Transaction, error) {
	//TODO

	data := make([]byte, 0)

	method, err := hexutil.Decode(hexfunc)

	bdigit := u.unpack(12, digit.Bytes())
	ba := addr.Bytes()

	data = append(data, method...)
	data = append(data, bdigit...)
	data = append(data, ba...)

	u.logf("input: %x", data)

	nonce := u.Nonce()
	tx := types.NewTransaction(nonce, BOEKI_CONTRACT, nil, u.GasLimit(), u.GasPrice(), data)

	stx, err := u.ks.SignTx(*u.ma, tx, u.NetworkID())
	if err != nil {
		return nil, err
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
		defer cancel()
		tx = stx
		err = u.c.SendTransaction(ctx, tx)
		if err != nil {
			return nil, err
		}
		return tx, nil
	}

}

func (u *Util) BoekiCompactWithdraw(_amount *big.Int) (*types.Transaction, error) {
	data := make([]byte, 0)

	hexfunc := "0x2e1a7d4d"
	//hexfunc := "0x66cff89a"
	method, err := hexutil.Decode(hexfunc)
	if err != nil {
		return nil, err
	}

	amount := u.unpack(32, _amount.Bytes())
	data = append(data, method...)
	data = append(data, amount...)

	nonce := u.Nonce()
	tx := types.NewTransaction(nonce, BOEKI_CONTRACT, nil, u.GasLimit(), u.GasPrice(), data)

	stx, err := u.ks.SignTx(*u.ma, tx, u.NetworkID())
	if err != nil {
		return nil, err
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
		defer cancel()
		tx = stx
		err = u.c.SendTransaction(ctx, tx)
		if err != nil {
			return nil, err
		}
		return tx, nil
	}
}

func (u *Util) BoekiCompactTrade(hexfunc string, _amount *big.Int, o *delta.Order) (*types.Transaction, error) {
	data := make([]byte, 0)

	method, err := hexutil.Decode(hexfunc)
	if err != nil {
		return nil, err
	}

	argpack := u.ToPack(o, _amount)
	data = append(data, method...)
	data = append(data, argpack...)

	u.logf("input(%d): %x", len(data), data)

	/*
		gasLimit, err := u.d.EstimateGas(u.ma.Address, &BOEKI_CONTRACT, data)
		if err != nil {
			return nil, err
		}*/

	gasLimit := uint64(500000)

	nonce := u.Nonce()
	tx := types.NewTransaction(nonce, BOEKI_CONTRACT, nil, gasLimit, u.GasPrice(), data)

	stx, err := u.ks.SignTx(*u.ma, tx, u.NetworkID())
	if err != nil {
		return nil, err
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
		defer cancel()

		tx = stx
		err = u.c.SendTransaction(ctx, tx)
		if err != nil {
			return nil, err
		}
		return tx, nil
	}

}

func (u *Util) ToPack(o *delta.Order, _amount *big.Int) []byte {
	data := make([]byte, 0)

	c := o.GetC()
	side, _ := o.GetSide()

	age := u.unpack(12, c.AmountGet.Bytes())
	agi := u.unpack(12, c.AmountGive.Bytes())
	amount := age
	if _amount != nil {
		amount = u.unpack(12, _amount.Bytes())
	}
	exp := u.unpack(8, c.Expires.Bytes())
	n := u.unpack(8, c.Nonce.Bytes())
	user := c.User.Bytes()
	var token []byte
	switch side {
	case "buy":
		token = c.TokenGet.Bytes()
	case "sell":
		token = c.TokenGive.Bytes()
	}

	r, s := o.GetRS()

	data = append(data, byte(c.V))
	data = append(data, age...)
	data = append(data, agi...)
	data = append(data, amount...)
	data = append(data, exp...)
	data = append(data, n...)
	data = append(data, user...)
	data = append(data, token...)
	data = append(data, r...)
	data = append(data, s...)

	return data
}

func (u *Util) BoekiCompactSellBuy(hexfunc string, s_amount, b_amount *big.Int, sello, buyo *delta.Order) (*types.Transaction, error) {
	data := make([]byte, 0)

	method, err := hexutil.Decode(hexfunc)
	if err != nil {
		return nil, err
	}

	argpackSell := u.ToPack(sello, s_amount)
	argpackBuy := u.ToPack(buyo, b_amount)
	data = append(data, method...)
	data = append(data, argpackSell...)
	data = append(data, argpackBuy...)

	gasLimit, err := u.d.EstimateGas(u.ma.Address, &BOEKI_CONTRACT, data)
	if err != nil {
		return nil, err
	}

	u.logf("EsimateGas: %d", gasLimit)

	nonce := u.Nonce()
	tx := types.NewTransaction(nonce, BOEKI_CONTRACT, nil, gasLimit, u.GasPrice(), data)

	stx, err := u.ks.SignTx(*u.ma, tx, u.NetworkID())
	if err != nil {
		return nil, err
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
		defer cancel()
		tx = stx
		err = u.c.SendTransaction(ctx, tx)
		if err != nil {
			return nil, err
		}
		return tx, nil
	}

}

func (u *Util) SellHigh(expires *big.Int, addr common.Address, hexfunc string) (*types.Transaction, error) {
	//TODO

	data := make([]byte, 0)

	method, err := hexutil.Decode(hexfunc)
	if err != nil {
		return nil, err
	}
	bexp := u.unpack(16, expires.Bytes())
	baddr := addr.Bytes()

	data = append(data, method...)
	data = append(data, bexp...)
	data = append(data, baddr...)

	nonce := u.Nonce()
	tx := types.NewTransaction(nonce, BOEKI_CONTRACT, nil, u.GasLimit(), u.GasPrice(), data)

	stx, err := u.ks.SignTx(*u.ma, tx, u.NetworkID())
	if err != nil {
		return nil, err
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
		defer cancel()
		tx = stx
		err = u.c.SendTransaction(ctx, tx)
		if err != nil {
			return nil, err
		}
		return tx, nil
	}

}

func (u *Util) ChangeOperator(operator common.Address) (*types.Transaction, error) {
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

	return u.proxy.ChangeOperator(opts, operator)
}

func (u *Util) WithdrawBoeki(amount *big.Int) (*types.Transaction, error) {
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

	return u.boeki.Withdraw(opts, amount)
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

func (u *Util) DepositBoekiProxy(amount *big.Int) (*types.Transaction, error) {
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

	return u.proxy.Deposit_BOEKI(opts)
}

func (u *Util) DeployBoeki(target common.Address, bot common.Address, useProxy bool) (common.Address, *types.Transaction, *Boekismart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pk, err := u.privateKey()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	opts := bind.NewKeyedTransactor(pk)

	opts.GasPrice = u.gasPrice
	opts.GasLimit = u.gasLimit
	opts.Context = ctx
	opts.From = u.ma.Address

	addr, tx, bs, err := DeployBoekismart(opts, u.c, target, bot, useProxy)
	return addr, tx, bs, err
}

func (u *Util) DeployLogid(owner common.Address, minEther *big.Int) (common.Address, *types.Transaction, *Logid, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pk, err := u.privateKey()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	opts := bind.NewKeyedTransactor(pk)

	opts.GasPrice = u.gasPrice
	opts.GasLimit = u.gasLimit
	opts.Context = ctx
	opts.From = u.ma.Address

	addr, tx, bs, err := DeployLogid(opts, u.c, minEther, owner)
	return addr, tx, bs, err
}

func (u *Util) DeployBoekiProxy() (common.Address, *types.Transaction, *Boekiproxy, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pk, err := u.privateKey()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	opts := bind.NewKeyedTransactor(pk)

	opts.GasPrice = u.gasPrice
	opts.GasLimit = u.gasLimit
	opts.Context = ctx
	opts.From = u.ma.Address

	addr, tx, bs, err := DeployBoekiproxy(opts, u.c, BOEKI_CONTRACT, DELTA_CONTRACT)
	return addr, tx, bs, err
}
