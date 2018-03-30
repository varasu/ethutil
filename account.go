package main

import (
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func (u *Util) cmdAccount(args []string) error {
	if len(args) >= 1 {
		switch args[0] {
		case "unlock":
			if len(args) == 1 {
				u.logf("use: unlock [minutes:timeout]")
			} else if len(args) >= 2 {
				minutes, err := strconv.Atoi(args[1])
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				u.logf("unlock %s for %d minutes", u.MainAddress().String(), minutes)
				p, err := u.readSecret()
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				err = u.ks.TimedUnlock(*u.ma, p, time.Minute*time.Duration(minutes))
				if err != nil {
					u.logf("%s", err)
				} else {
					u.logf("%s unlocked for %d minutes.", u.MainAddress().String(), minutes)
				}
			}

		case "list":
			for i, a := range u.ks.Accounts() {
				u.logf("%d: %s", i+1, a.Address.String())
			}

		case "set":
			if len(args) >= 2 {
				n, err := strconv.Atoi(args[1])
				if err != nil {
					u.logf("%s", err)
					return nil
				}
				if len(u.ks.Accounts()) >= n {
					u.SetMainAccount(u.ks.Accounts()[n-1])
				} else {
					u.logf("Out of boundary.")
				}
			}

		case "new":
			p, err := u.readSecret()
			if err != nil {
				u.logf("%s", err)
				return nil
			}
			if p == "" {
				u.logf("Do not use empty passphrase.")
				return nil
			}
			a, err := u.ks.NewAccount(p)
			if err != nil {
				u.logf("%s", err)
				return nil
			}
			u.logf("New account created at %s", a.Address.String())
			if len(u.ks.Accounts()) == 1 {
				u.SetMainAccount(a)
			}

		case "import":
			if len(args) >= 3 {
				switch args[1] {
				case "hex":
					s := args[2]

					pkey, err := crypto.HexToECDSA(s)
					if err != nil {
						u.logf("%s", err)
						return nil
					}

					p, err := u.readSecret()
					if err != nil {
						u.logf("%s", err)
						return nil
					}
					a, err := u.ks.ImportECDSA(pkey, p)
					if err != nil {
						u.logf("%s", err)
						return nil
					}
					u.logf("%s is imported.", a.Address.String())

				default:
					u.logf("%s not implemented.", args[1])
				}
			} else {
				u.logf("import hex|json input[string|file]")
			}
		}
	} else {
		u.logf("commands: list new import")
	}
	return nil
}
