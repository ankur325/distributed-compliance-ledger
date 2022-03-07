package keeper

import (
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/base/types"
)

var _ types.QueryServer = Keeper{}
