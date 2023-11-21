package address

import "github.com/hy-u-3a4em/rana/core/source/public_key"

type Address struct {
	encoded   string
	publicKey public_key.PublicKey
	isValid   bool
}
