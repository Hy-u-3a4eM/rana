package transaction

import (
	"github.com/hy-u-3a4em/rana/core/source/amount"
	"github.com/hy-u-3a4em/rana/core/source/time_instant"
)

type Transaction struct {
	fee      *amount.Amount
	deadline time_instant.TimeInstant
}
