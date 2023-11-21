package key_pair

import (
	"github.com/hy-u-3a4em/rana/core/source/private_key"
	"github.com/hy-u-3a4em/rana/core/source/public_key"
)

type KeyPair struct {
	privateKey private_key.PrivateKey
	publicKey  public_key.PublicKey
}
