package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"k.veselin.org/boeki/chain"
	"k.veselin.org/boeki/delta"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (u *Util) Sha256(digit *big.Int, other common.Address) []byte {
	h := sha256.New()
	// contract
	h.Write(GREET_CONTRACT.Bytes())
	// admin
	h.Write(u.ma.Address.Bytes())
	// digit
	h.Write(digit.Bytes())
	// other
	h.Write(other.Bytes())

	return h.Sum(nil)
}

func digitHash(digit *big.Int) []byte {
	h := sha256.New()
	h.Write(digit.Bytes())
	return h.Sum(nil)
}

func (u *Util) cmdDelta(args []string) error {
	if len(args) >= 1 {
		switch args[0] {
		case "filled":
			filled, err := u.d.AmountFilled(u.order.PureTradeParams())
			if err != nil {
				u.logf("%s", err)
				return nil
			}
			u.logf("Amount Filled: %s", filled.String())
		case "hash":
			if len(args) == 3 {
				digit, _ := new(big.Int).SetString(args[1], 10)
				other := common.HexToAddress(args[2])
				chainHash, _ := u.greet.ComplexHash(nil, digit, other)
				u.logf("chain hash: %x", chainHash)

				localHash := u.Sha256(digit, other)
				u.logf("local hash: %x", localHash)

				chainDigit, _ := u.greet.SimpleHashDigit(nil, digit)
				u.logf("chain digit %x", chainDigit)

				localDigit := digitHash(digit)
				u.logf("local digit %x", localDigit)

			}
		case "trade":
			o := u.order

			if len(args) == 2 {
				switch args[1] {
				case "sell":
					o = u.orderSell
				case "buy":
					o = u.orderBuy
				}
			}

			amount := o.GetAmountGet()
			u.logf("%s", amount.String())
			txs, err := u.Trade(amount, o)
			if err != nil {
				u.logf("%s", err)
				return nil
			}
			u.AddPending(txs)
			u.logf("%s", txs)
			u.logf("%s", txs.Hash().String())

		case "make":
			if len(args) == 5 {
				switch args[1] {
				case "sell":
					// amountGive amountGet
					// i want to sell my tokens
					tgi := u.taddr
					d, _ := u.token.Decimals(nil)
					agi, _ := ethToWeiDecimal(args[2], int(d.Int64()))
					tge := ETHER_TOKEN
					age, _ := ethToWei(args[3])
					exp, ok := new(big.Int).SetString(args[4], 10)
					if !ok {
						u.logf("Exp convert problem.")
						return nil
					}
					nonce := new(big.Int).SetInt64(rand.Int63n(100000))
					order, err := u.MakeOrder(tge, age, tgi, agi, exp, nonce, u.ma.Address)
					if err != nil {
						u.logf("Make Order: %s", err)
						return nil
					}

					u.logf("AmountGet: %d", age)
					u.logf("AmountGive: %d", agi)

					if order == nil {
						u.logf("Nil order.")
						return nil
					}

					//u.logf("%+v", order)
					u.order = order
					u.orderSell = order
					u.logf("%+v\n%+v", u.order, u.order.GetC())

					ke, sha := u.CheckSignatures(tge, age, tgi, agi, exp, nonce, u.ma.Address)
					u.logf("k: %s, s: %s", ke.String(), sha.String())

				case "buy":
					// i wan to buy token

					tge := u.taddr
					d, _ := u.token.Decimals(nil)
					u.logf("Token deciaml: %s", d.String())
					age, _ := ethToWeiDecimal(args[2], int(d.Int64()))
					tgi := ETHER_TOKEN
					agi, _ := ethToWei(args[3])

					// Expires
					exp, _ := new(big.Int).SetString(args[4], 10)
					nonce := new(big.Int).SetInt64(rand.Int63n(100000))
					order, err := u.MakeOrder(tge, age, tgi, agi, exp, nonce, u.ma.Address)
					if err != nil {
						u.logf("Make Order: %s", err)
						return nil
					}

					u.order = order
					u.orderBuy = order
					u.logf("%+v\n%+v", u.order, u.order.GetC())

					ke, sha := u.CheckSignatures(tge, age, tgi, agi, exp, nonce, u.ma.Address)
					u.logf("k: %s, s: %s", ke.String(), sha.String())

				}
			}
		case "admin":
			admin := u.d.Admin()
			u.logf("Delta Admin: %s", admin.String())
		case "balance":
			ba := common.Address{}
			var other common.Address
			if len(args) >= 2 {
				switch args[1] {
				case "boeki":
					ba = BOEKI_CONTRACT
				case "proxy":
					ba = PROXY_CONTRACT
				default:
					ba = common.HexToAddress(args[1])
				}
			} else {
				ba = u.ma.Address
			}
			if len(args) >= 3 {
				other = common.HexToAddress(args[2])
			}
			ether := u.BalanceOf(ETHER_TOKEN, ba).String()
			if *mainNetwork {
				u.logf("delta balances (%s):", ba.String())
				u.logf("ETH: %s", ether)
				if len(args) >= 3 {
					u.logf("%s: %s", other.String(), u.BalanceOf(other, ba).String())
				}
				return nil
			}
			kusari := u.BalanceOf(KUSARI_TOKEN, ba).String()
			chisai := u.BalanceOf(CHISAI_TOKEN, ba).String()
			u.logf("delta balances (%s):", ba.String())
			u.logf("ETH: %s | KSR: %s | CIS: %s", ether, kusari, chisai)
		case "deposit":
			if len(args) == 2 {
				amount, err := ethToWei(args[1])
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				txs, err := u.Deposit(amount)
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				u.logf("%s", txs)
				u.logf("%s", txs.Hash().String())
			}
		case "withdraw":
			u.SetGasLimit(40000)
			u.logf("gasLimit set to 40 000 for this operation.")
			if len(args) == 2 {
				amount, err := ethToWei(args[1])
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				txs, err := u.Withdraw(amount)
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				u.AddPending(txs)
				u.logf("%s", txs)
				u.logf("%s", txs.Hash().String())
			}
		case "depositToken":
			if len(args) == 3 {
				taddr := common.HexToAddress(args[1])
				amount, err := ethToWei(args[2])
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				txs, err := u.DepositToken(taddr, amount)
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				u.AddPending(txs)
				u.logf("%s", txs)
				u.logf("%s", txs.Hash().String())
			}

		case "approve":
			if len(args) == 2 {
				amount, err := ethToWei(args[1])
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				txs, err := u.Approve(DELTA_CONTRACT, amount)
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				u.AddPending(txs)
				u.logf("%s", txs)
				u.logf("%s", txs.Hash().String())
			}
		case "":
		}
	}

	return nil
}

func (u *Util) Receipt(txHash common.Hash) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return u.c.TransactionReceipt(ctx, txHash)
}

func (u *Util) BalanceOf(taddr common.Address, user common.Address) *big.Float {
	b, _ := u.d.BalanceOf(taddr, user)
	d := new(big.Int).SetInt64(18)
	if !reflect.DeepEqual(ETHER_TOKEN, taddr) {
		token, _ := chain.NewStandardToken(taddr, u.c)
		d, _ = token.Decimals(nil)
	} else {
	}
	return weiToEthDecimal(b, int(d.Int64()))
}

func (u *Util) cmdToken(args []string) error {
	if len(args) >= 1 {
		switch args[0] {
		case "add":
			if len(args) >= 2 {
				var err error
				taddr := common.HexToAddress(args[1])
				u.taddr = taddr
				u.token, err = chain.NewStandardToken(taddr, u.c)
				if err != nil {
					u.logf("AddToken: %s", err)
					return nil
				}
				u.logf("Token is added.")
			}
		case "balance":
			if len(args) >= 2 {
			} else {
				if u.token != nil {
					n, _ := u.token.Name(nil)
					b, err := u.token.BalanceOf(nil, u.ma.Address)
					if err != nil {
						u.logf("%s", err)
						return nil
					}
					d, _ := u.token.Decimals(nil)
					balance := weiToEthDecimal(b, int(d.Int64()))
					u.logf("%s: %s", n, balance.String())
				} else {
					u.logf("Add token please")
				}
			}
		case "deposit":
			if len(args) == 2 {
				amount, err := ethToWei(args[1])
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				txs, err := u.DepositToken(u.taddr, amount)
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				u.AddPending(txs)
				u.logf("%s", txs)
				u.logf("%s", txs.Hash().String())

			}

		case "withdraw":
			if len(args) == 2 {
				amount, err := ethToWei(args[1])
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				txs, err := u.WithdrawToken(u.taddr, amount)
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				u.AddPending(txs)
				u.logf("%s", txs)
				u.logf("%s", txs.Hash().String())

			}

		case "transfer":
			// from to amount || to amount
			if len(args) == 3 {
				to := common.HexToAddress(args[1])
				d, _ := u.token.Decimals(nil)
				amount, err := ethToWeiDecimal(args[2], int(d.Int64()))
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				if u.token == nil {
					u.logf("Missing Token")
					return nil
				}
				txs, err := u.Transfer(to, amount)
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				u.logf("%s", txs)
			}
		}
	}
	return nil
}

func (u *Util) privateKey() (*ecdsa.PrivateKey, error) {
	pass := ""
	if *mainNetwork {
		pass = ""
	}
	jkey, err := u.ks.Export(*u.ma, pass, "")
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(jkey, "")
	if err != nil {
		return nil, err
	}
	return key.PrivateKey, nil
}

func (u *Util) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
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

	return u.token.Transfer(opts, to, amount)
}

func (u *Util) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
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

	return u.token.Approve(opts, spender, amount)
}

func (u *Util) IsAllowed(amount *big.Int, owner, spender common.Address) bool {
	remaining, err := u.token.Allowance(nil, owner, spender)
	if err != nil {
		return false
	}
	if amount.Cmp(remaining) == 0 || amount.Cmp(remaining) == -1 {
		return true
	} else {
		if remaining.Uint64() != 0 {
			u.logf("Not enough. remaining: %s", remaining)
		}
		return false
	}
}

func (u *Util) DepositToken(taddr common.Address, amount *big.Int) (*types.Transaction, error) {
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

	if !u.IsAllowed(amount, u.ma.Address, DELTA_CONTRACT) {
		return nil, errors.New("Approve this transfer first.")
	}

	return u.d.DepositToken(opts, taddr, amount)
}

func (u *Util) WithdrawToken(taddr common.Address, amount *big.Int) (*types.Transaction, error) {
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

	return u.d.WithdrawToken(opts, taddr, amount)
}

func (u *Util) Deposit(amount *big.Int) (*types.Transaction, error) {
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

	// amount of Ether to deposit to delta smart contract
	opts.Value = amount

	return u.d.Deposit(opts)
}

func (u *Util) Withdraw(amount *big.Int) (*types.Transaction, error) {
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

	return u.d.Withdraw(opts, amount)
}

func (u *Util) cmdEth(args []string) error {
	if len(args) >= 1 {
		switch args[0] {
		case "gas":
			//
		case "set":
			if len(args) == 3 {
				switch args[1] {
				case "gasPrice":
					u.SetGasPriceGwei(args[2])
					u.logf("Done.")
				case "gasLimit":
					u.SetGasLimitString(args[2])
					u.logf("Done.")
				default:
					u.logf("help: set gasPrice|gasLimit [value]")
				}
			}
		case "balance":
			a := u.ma.Address
			if len(args) == 2 {
				if args[1] == "boeki" {
					a = BOEKI_CONTRACT
				}
			}
			b, _ := u.BalanceEth(a)
			u.logf("%s: %s", a.String(), b.String())
		case "deploy":
			if len(args) >= 2 {
				file := args[1]
				b, err := ioutil.ReadFile(file)
				if err != nil {
					u.logf("ReadFile: %s", err)
					return nil
				}
				hex := strings.TrimSpace(string(b))

				binary, err := hexutil.Decode("0x" + hex)
				if err != nil {
					u.logf("hexutil.Decode: %s", err)
					return nil
				}

				nonce := u.Nonce()
				if len(args) >= 3 {
					u.SetGasPriceGwei(args[2])
				}
				if len(args) >= 4 {
					gl, err := strconv.Atoi(args[3])
					if err != nil {
						u.logf("Atoi: %s", err)
						return nil
					}
					u.SetGasLimit(uint64(gl))
				}

				tx := types.NewContractCreation(nonce, nil, u.GasLimit(), u.GasPrice(), binary)

				stx, err := u.ks.SignTx(*u.ma, tx, u.NetworkID())
				if err != nil {
					u.logf("signtx: %s", err)
				} else {
					ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
					defer cancel()
					tx = stx
					u.AddPending(tx)
					err = u.c.SendTransaction(ctx, tx)
					if err != nil {
						u.logf("sendtransaction: %s", err)
						return nil
					}
					u.logf("%s", tx)
				}

			}
		case "send":
			// @address @value
			if len(args) >= 3 {
				to := common.HexToAddress(args[1])
				amount, _ := ethToWei(args[2])
				nonce := u.Nonce()

				if len(args) >= 4 {
					u.SetGasPriceGwei(args[3])
				}

				u.logf("Send %s ether to %s", amount.String(), to.String())

				tx := types.NewTransaction(nonce, to, amount, u.GasLimit(), u.GasPrice(), make([]byte, 0))

				stx, err := u.ks.SignTx(*u.ma, tx, u.NetworkID())
				if err != nil {
					u.logf("signtx: %s", err)
				} else {
					ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
					defer cancel()
					tx = stx
					err = u.c.SendTransaction(ctx, tx)
					if err != nil {
						u.logf("sendtransaction: %s", err)
						return nil
					}
					u.AddPending(tx)
					u.logf("%s", tx)
					u.logf("%s", tx.Hash().String())
				}
			}
		default:
			u.logf("...")
		}
	} else {
		u.logf("commands: balance")
	}
	return nil
}

func (u *Util) Nonce() uint64 {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	nonce, _ := u.c.NonceAt(ctx, u.MainAddress(), nil)
	return nonce
}

func (u *Util) NetworkID() *big.Int {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	id, _ := u.c.NetworkID(ctx)
	return id

}

func (u *Util) BalanceEth(addr common.Address) (*big.Float, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	b, err := u.c.BalanceAt(ctx, addr, nil)
	if err != nil {
		return nil, err
	}
	return weiToEth(b), nil
}

func (u *Util) MakeOrder(tge common.Address, age *big.Int, tgi common.Address, agi *big.Int, exp, nonce *big.Int, usr common.Address) (*delta.Order, error) {
	order := delta.MakeOrder(tge, age, tgi, agi, exp, nonce, usr)
	_hash, _ := u.greet.HashTrade(nil, DELTA_CONTRACT, tge, age, tgi, agi, exp, nonce)
	h := make([]byte, 0)
	for _, v := range _hash {
		h = append(h, v)
	}
	u.logf("trade hash: %x", h)
	// append signHash before Sign the message.
	// Sign calculates an Ethereum ECDSA signature for:
	// keccack256("\x19Ethereum Signed Message:\n" + len(message) + message))
	// Pure sign function
	sig, err := u.ks.SignHash(*u.ma, signHash(h))
	if err != nil {
		return nil, err
	}
	u.logf("sign hash: %x", sig)
	if !order.AddSignature(sig) {
		return nil, errors.New("Signature err.")
	}
	order.Updated = time.Now()
	return order, nil
}

func (u *Util) Trade(amount *big.Int, order *delta.Order) (*types.Transaction, error) {
	u.d.SetGasPriceWei(u.GasPrice())
	u.d.SetGasLimit("900000")
	u.d.SetAccount(*u.ma)

	avalVolume, err := u.d.AvailableVolume(order.PureTradeParams())
	if err != nil {
		return nil, err
	}
	u.logf("Available Volume: %d", avalVolume)

	/*
		valid, err := u.d.TestTrade(order.TestTradeParams(amount, u.ma.Address))
		if err != nil {
			return nil, err
		}
		if !valid {
			return nil, errors.New("Do not trade this order.")
		}*/

	return u.d.Trade(order.TradeParams(nil, amount))
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func (u *Util) CheckSignatures(tge common.Address, age *big.Int, tgi common.Address, agi *big.Int, exp, nonce *big.Int, usr common.Address) (common.Address, common.Address) {
	_hash256, _ := u.greet.HashTrade(nil, DELTA_CONTRACT, tge, age, tgi, agi, exp, nonce)
	h := make([]byte, 0)
	for _, v := range _hash256 {
		h = append(h, v)
	}
	u.logf("trade hash: %x", h)
	sig, err := u.ks.SignHash(*u.ma, signHash(h))
	if err != nil {
		u.logf("%s", err)
		return common.Address{}, common.Address{}
	}
	u.logf("sign hash: %x", sig)
	var r [32]byte
	for i := 0; i <= 31; i++ {
		r[i] = sig[i]
	}
	var s [32]byte
	for i := 32; i <= 63; i++ {
		s[i-32] = sig[i]
	}
	var v uint8 = sig[64]
	v += 27

	first, _ := u.greet.CheckSignatureKeccak(nil, DELTA_CONTRACT, tge, age, tgi, agi, exp, nonce, v, r, s)
	second, _ := u.greet.CheckSignatureSha3(nil, DELTA_CONTRACT, tge, age, tgi, agi, exp, nonce, v, r, s)
	return first, second
}

func weiToEth(wei *big.Int) *big.Float {
	base := new(big.Float).SetFloat64(math.Pow10(18))
	nwei := new(big.Float).SetInt(wei)
	rwei := new(big.Float).Quo(nwei, base)
	return rwei
}

func weiToEthDecimal(wei *big.Int, decimal int) *big.Float {
	base := new(big.Float).SetFloat64(math.Pow10(decimal))
	nwei := new(big.Float).SetInt(wei)
	rwei := new(big.Float).Quo(nwei, base)
	return rwei
}

func ethToWei(seth string) (*big.Int, error) {
	return ethToWeiDecimal(seth, 18)
}

func ethToWeiDecimal(seth string, decimals int) (*big.Int, error) {
	r := new(big.Int)
	eth := new(big.Float)
	_, ok := eth.SetString(seth)
	if !ok {
		return r, errors.New("ethToWei Failed.")
	}
	base := new(big.Float).SetFloat64(math.Pow10(decimals))
	wei := new(big.Float).Mul(eth, base)
	wei.Int(r)
	return r, nil

}
