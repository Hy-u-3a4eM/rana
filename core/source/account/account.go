package account

import (
	"github.com/hy-u-3a4em/rana/core/source/address"
	"github.com/hy-u-3a4em/rana/core/source/key_pair"
)

type Account struct {
	keyPair key_pair.KeyPair
	address address.Address
}
