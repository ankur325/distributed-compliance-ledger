package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestQuerier_QueryValidator(t *testing.T) {
	setup := Setup()

	// add validator
	validator := DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)

	// query validator
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryValidator, validator.Address.String()},
		abci.RequestQuery{},
	)

	var receivedValidator types.Validator
	setup.Cdc.MustUnmarshalJSON(result, &receivedValidator)

	// check
	require.Equal(t, validator, receivedValidator)
}

func TestQuerier_QueryValidator_ForUnknown(t *testing.T) {
	setup := Setup()

	// query validator
	result, err := setup.Querier(
		setup.Ctx,
		[]string{QueryValidator, test_constants.ConsensusAddress1.String()},
		abci.RequestQuery{},
	)

	// check
	require.Nil(t, result)
	require.NotNil(t, err)
	require.Equal(t, types.CodeValidatorDoesNotExist, err.Code())
}

func TestQuerier_QueryValidators(t *testing.T) {
	setup := Setup()

	// add 2 validators
	validator1, validator2 := StoreTwoValidators(setup)

	// query all validators
	listValidators := getValidators(setup, types.All)

	// check
	require.Equal(t, 2, listValidators.Total)
	require.Equal(t, 2, len(listValidators.Items))
	require.Equal(t, validator1, listValidators.Items[0])
	require.Equal(t, validator2, listValidators.Items[1])

	// query Active validators
	activeValidators := getValidators(setup, types.Active)

	// check
	require.Equal(t, 2, activeValidators.Total)
	require.Equal(t, 2, len(activeValidators.Items))
	require.Equal(t, validator1, activeValidators.Items[0])
	require.Equal(t, validator2, activeValidators.Items[1])

	// query Jailed validators
	jailedValidators := getValidators(setup, types.Jailed)

	// check
	require.Equal(t, 0, jailedValidators.Total)
	require.Equal(t, 0, len(jailedValidators.Items))
}

func getValidators(setup TestSetup, state types.ValidatorState) types.ListValidatorItems {
	paginationParams := pagination.NewPaginationParams(0, 0)
	params, _ := setup.Cdc.MarshalJSON(types.NewListValidatorsParams(paginationParams, state))

	// query all validators
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryValidators},
		abci.RequestQuery{Data: params},
	)

	var listValidators types.ListValidatorItems
	setup.Cdc.MustUnmarshalJSON(result, &listValidators)

	return listValidators
}
