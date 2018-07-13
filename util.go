package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"k.veselin.org/boeki/chain"
	"k.veselin.org/boeki/delta"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"golang.org/x/crypto/ssh/terminal"
)

var keystoreDir = flag.String("ksd", "/data/eth/keystore", "keystore dir")
var provider = flag.String("p", "", "provider url")
var mainNetwork = flag.Bool("main", false, "enable main network")

type command struct {
	run      func(*Util, []string) error
	complete func() []string
}

var commands = map[string]command{
	"logid":   command{run: (*Util).cmdLogid},
	"boeki":   command{run: (*Util).cmdBoeki},
	"token":   command{run: (*Util).cmdToken},
	"delta":   command{run: (*Util).cmdDelta},
	"account": command{run: (*Util).cmdAccount},
	"eth":     command{run: (*Util).cmdEth},
	"help":    command{run: (*Util).cmdHelp},
	"quit":    command{run: (*Util).cmdQuit},
}

func (app *Util) cmdQuit(args []string) error {
	if len(args) > 0 {
		app.logf("the QUIT command takes no argument")
		return nil
	}
	app.logf("Exit program: %s", time.Now())
	return errExitApp
}

func (app *Util) cmdHelp(args []string) error {
	app.logf("eth: %s", time.Now())
	app.logf("help")
	return nil
}

type Util struct {
	term    *terminal.Terminal
	prompt  string
	outpref string
	done    chan bool

	// EhtClinet
	c *ethclient.Client
	d *chain.DeltaChain
	// KeyStore
	ks *keystore.KeyStore
	// Main Account
	ma *accounts.Account

	// order for trading
	order *delta.Order

	orderSell *delta.Order
	orderBuy  *delta.Order

	// testing Greeter smart contract
	greet *Greet

	// Boeki Smart contract
	boeki *Boekismart
	proxy *Boekiproxy

	logid *Logid

	token *chain.StandardToken
	taddr common.Address

	// pending transaction
	pending *types.Transaction

	// ethereum properties
	// set global gasLimit
	gasLimit uint64
	// set global gasPrice
	gasPrice *big.Int
}

func (u *Util) MainAddress() common.Address {
	return u.ma.Address
}

func (u *Util) SetGasLimit(gl uint64) {
	u.gasLimit = gl
}

func (u *Util) SetGasLimitString(sgl string) {
	gl, _ := strconv.Atoi(sgl)
	u.gasLimit = uint64(gl)
}

func (u *Util) SetGasPriceGwei(gp string) {
	p, _ := gweiToWei(gp)
	u.gasPrice = p
}

func (u *Util) GasLimit() uint64 {
	return u.gasLimit
}

func (u *Util) GasPrice() *big.Int {
	return u.gasPrice
}

func gweiToWei(sgwei string) (*big.Int, error) {
	r := new(big.Int)
	gwei, ok := new(big.Int).SetString(sgwei, 10)
	if !ok {
		return r, errors.New("gweiToWei Failed.")
	}
	base := new(big.Int)
	new(big.Float).SetFloat64(math.Pow10(9)).Int(base)
	r.Mul(gwei, base)
	return r, nil
}

func (u *Util) SetKeyStore(ksd string) error {
	u.ks = keystore.NewKeyStore(ksd, 16, 2)
	if u.ks == nil {
		return errors.New("KeyStore err.")
	} else {
		if len(u.ks.Accounts()) > 0 {
			// Set Account
			a := u.ks.Accounts()[0]
			u.SetMainAccount(a)
		} else {
			u.prompt += "XXXX): "
		}

		return nil
	}
}

func (u *Util) SetMainAccount(a accounts.Account) {
	// set prompt for clarity
	if u.term != nil {
		u.term.SetPrompt(fmt.Sprintf("eth(%s): ", a.Address.String()[:6]))
	} else {
		u.prompt += a.Address.String()[:6] + "): "
	}
	u.ma = &a

	//u.logf("Set %s as main account.", u.ma.Address.String())
}

var errExitApp = errors.New("internal sentinel error value to quit the console reading loop")

func main() {
	var err error
	flag.Parse()

	// enable main network
	if *mainNetwork {
		log.Printf("Main Network enabled.\n")
	}

	util := &Util{}

	//ch
	util.done = make(chan bool, 1)

	// Set KeyStore
	err = util.SetKeyStore(*keystoreDir)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	// ethclient provide
	util.c, err = ethclient.Dial(*provider)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	// setup DeltaChain
	util.d, err = chain.NewDeltaChain(*keystoreDir, *provider, DELTA_CONTRACT)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	// new logid smart
	util.logid, err = NewLogid(LOGID_ADDR, util.c)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	// Boeki Smart contract
	util.boeki, err = NewBoekismart(BOEKI_CONTRACT, util.c)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	// Boeki Proxy
	util.proxy, err = NewBoekiproxy(PROXY_CONTRACT, util.c)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	// setup Greeter
	util.greet, err = NewGreet(GREET_CONTRACT, util.c)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	// test token
	util.taddr = KUSARI_TOKEN
	util.token, err = chain.NewStandardToken(KUSARI_TOKEN, util.c)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	// defaults for sending value: ether
	util.SetGasLimit(21000)
	util.SetGasPriceGwei("1")

	// NewHead
	go util.HeadWriter()

	if err = util.Main(); err != nil {
		if util.term != nil {
			util.logf("%s", err.Error())
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
		os.Exit(1)
	}

}

func (app *Util) Main() error {
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		return err
	}
	defer terminal.Restore(0, oldState)

	var screen = struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}

	prompt := "eth(" + app.prompt
	app.term = terminal.NewTerminal(screen, prompt)

	errc := make(chan error, 2)
	go func() { errc <- app.readConsole() }()
	return <-errc
}

func (app *Util) logf(format string, args ...interface{}) {
	fmt.Fprintf(app.term, app.outpref+format+"\n", args...)
}

func lookupCommand(prefix string) (name string, c command, ok bool) {
	prefix = strings.ToLower(prefix)
	if c, ok = commands[prefix]; ok {
		return prefix, c, ok
	}

	for full, candidate := range commands {
		if strings.HasPrefix(full, prefix) {
			if c.run != nil {
				return "", command{}, false
			}
			c = candidate
			name = full
		}
	}
	return name, c, c.run != nil
}

func (app *Util) readConsole() error {
	for {
		line, err := app.term.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("termial.ReadLine: %s", err.Error())
		}
		f := strings.Fields(line)
		if len(f) == 0 {
			continue
		}
		cmd, args := f[0], f[1:]
		if _, c, ok := lookupCommand(cmd); ok {
			err = c.run(app, args)
		} else {
			app.logf("Unknown command %q", line)
		}
		if err == errExitApp {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

func (app *Util) readSecret() (string, error) {
	for {
		secret, err := app.term.ReadPassword("type your secret: ")
		if err != nil {
			return "", fmt.Errorf("termial.ReadPassword: %s", err.Error())
		}
		return secret, nil
	}
}
