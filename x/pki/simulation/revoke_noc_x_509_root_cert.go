package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func SimulateMsgRevokeNocRootX509Cert(
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRevokeNocX509RootCert{
			Signer: simAccount.Address.String(),
		}

		// TODO: Handling the RevokeNocRootX509Cert simulation

		return simtypes.NoOpMsg(pkitypes.ModuleName, msg.Type(), "RevokeNocRootX509Cert simulation not implemented"), nil, nil
	}
}