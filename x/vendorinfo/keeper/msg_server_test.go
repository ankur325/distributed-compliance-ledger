package keeper_test

// import (
// 	"context"
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
// 	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/keeper"
// 	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
// )

// func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
// 	dclauthKeeper := &DclauthKeeperMock{}
// 	k, ctx := keepertest.VendorinfoKeeper(t, dclauthKeeper)
// 	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
// }