// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//nolint:testpackage
package keeper

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/types"
)

func TestKeeper_ModelInfoGetSet(t *testing.T) {
	setup := Setup()

	// check if model present
	require.False(t, setup.ModelinfoKeeper.IsModelInfoPresent(setup.Ctx, testconstants.VID,
		testconstants.PID, testconstants.SoftwareVersion, testconstants.HardwareVersion))

	// no model before its created
	require.Panics(t, func() {
		setup.ModelinfoKeeper.GetModelInfo(setup.Ctx, testconstants.VID, testconstants.PID,
			testconstants.SoftwareVersion, testconstants.SoftwareVersion)
	})

	// create model
	setup.ModelinfoKeeper.SetModelInfo(setup.Ctx, DefaultModelInfo())

	// check if model present
	require.True(t, setup.ModelinfoKeeper.IsModelInfoPresent(setup.Ctx, testconstants.VID, testconstants.PID,
		testconstants.SoftwareVersion, testconstants.HardwareVersion))

	// get model info
	modelInfo := setup.ModelinfoKeeper.GetModelInfo(setup.Ctx, testconstants.VID, testconstants.PID,
		testconstants.SoftwareVersion, testconstants.HardwareVersion)
	require.NotNil(t, modelInfo)
	require.Equal(t, testconstants.ProductName, modelInfo.Model.ProductName)
	require.Equal(t, testconstants.Owner, modelInfo.Owner)
}

func TestKeeper_ModelInfoIterator(t *testing.T) {
	setup := Setup()

	count := 10

	// add 10 models infos with same VID and check associated products
	PopulateStoreWithModelsHavingSameVendor(setup, count)

	// get total count
	totalModes := setup.ModelinfoKeeper.CountTotalModelInfos(setup.Ctx)
	require.Equal(t, count, totalModes)

	// get iterator
	var expectedRecords []types.ModelInfo

	setup.ModelinfoKeeper.IterateModelInfos(setup.Ctx, func(modelInfo types.ModelInfo) (stop bool) {
		expectedRecords = append(expectedRecords, modelInfo)

		return false
	})
	require.Equal(t, count, len(expectedRecords))
}

func TestKeeper_ModelInfoDelete(t *testing.T) {
	setup := Setup()

	// create model
	setup.ModelinfoKeeper.SetModelInfo(setup.Ctx, DefaultModelInfo())

	// check if model present
	require.True(t, setup.ModelinfoKeeper.IsModelInfoPresent(setup.Ctx, testconstants.VID, testconstants.PID,
		testconstants.SoftwareVersion, testconstants.HardwareVersion))

	// delete model
	setup.ModelinfoKeeper.DeleteModelInfo(setup.Ctx, testconstants.VID, testconstants.PID,
		testconstants.SoftwareVersion, testconstants.HardwareVersion)

	// check if model present
	require.False(t, setup.ModelinfoKeeper.IsModelInfoPresent(setup.Ctx, testconstants.VID,
		testconstants.PID, testconstants.SoftwareVersion, testconstants.HardwareVersion))

	// try to get model info
	require.Panics(t, func() {
		setup.ModelinfoKeeper.GetModelInfo(setup.Ctx, testconstants.VID, testconstants.PID,
			testconstants.SoftwareVersion, testconstants.HardwareVersion)
	})
}

func TestKeeper_VendorProductsUpdatesWithModelInfo(t *testing.T) {
	setup := Setup()
	count := 10

	// check if vendor products present
	require.False(t, setup.ModelinfoKeeper.IsVendorProductsPresent(setup.Ctx, testconstants.VID))

	// get vendor products
	vendorProducts := setup.ModelinfoKeeper.GetVendorProducts(setup.Ctx, testconstants.VID)
	require.True(t, vendorProducts.IsEmpty())

	var PIDs []types.Product

	// add 10 model infos with same VID and check associated products
	for i := 0; i < count; i++ {
		// add model info
		modelInfo := DefaultModelInfo()
		modelInfo.Model.PID = uint16(i)
		setup.ModelinfoKeeper.SetModelInfo(setup.Ctx, modelInfo)

		vendorProduct := types.Product{
			PID:             modelInfo.Model.PID,
			Name:            modelInfo.Model.ProductName,
			SoftwareVersion: modelInfo.Model.SoftwareVersion,
			HardwareVersion: modelInfo.Model.HardwareVersion,
			SKU:             modelInfo.Model.SKU,
			Owner:           modelInfo.Owner,
		}
		PIDs = append(PIDs, vendorProduct)

		// check vendor products
		vendorProducts = setup.ModelinfoKeeper.GetVendorProducts(setup.Ctx, testconstants.VID)
		require.Equal(t, PIDs, vendorProducts.Products)
	}

	// remove all model infos in a random way and check associated products
	for i := count; i > 0; i-- {
		//nolint:gosec
		index := uint16(rand.Intn(i))
		pid := PIDs[index]

		PIDs = append(PIDs[:index], PIDs[index+1:]...)

		// remove second model
		setup.ModelinfoKeeper.DeleteModelInfo(setup.Ctx, testconstants.VID, pid.PID,
			testconstants.SoftwareVersion, testconstants.HardwareVersion)

		// check vendor products
		vendorProducts = setup.ModelinfoKeeper.GetVendorProducts(setup.Ctx, testconstants.VID)
		require.Equal(t, PIDs, vendorProducts.Products)
	}

	// check if vendor products present
	require.False(t, setup.ModelinfoKeeper.IsVendorProductsPresent(setup.Ctx, testconstants.VID))
}

func TestKeeper_VendorProductsOverTwoModelsWithDifferentVendor(t *testing.T) {
	setup := Setup()

	PopulateStoreWithModelsHavingDifferentVendor(setup, 2)

	// check vendor products for first device
	vendorProductsForModel1 := setup.ModelinfoKeeper.GetVendorProducts(setup.Ctx, 1)
	require.Equal(t, 1, len(vendorProductsForModel1.Products))
	require.Equal(t, uint16(1), vendorProductsForModel1.Products[0].PID)

	// check vendor products for second device
	vendorProductsForModel2 := setup.ModelinfoKeeper.GetVendorProducts(setup.Ctx, 2)
	require.Equal(t, 1, len(vendorProductsForModel2.Products))
	require.Equal(t, uint16(2), vendorProductsForModel2.Products[0].PID)
}

func TestKeeper_VendorProductsIteratorOverOneVendor(t *testing.T) {
	setup := Setup()

	// add 10 model infos with same Vendor
	expectedLen := 1

	PopulateStoreWithModelsHavingSameVendor(setup, 10)

	// get total count
	totalVendorProducts := setup.ModelinfoKeeper.CountTotalVendorProducts(setup.Ctx)
	require.Equal(t, expectedLen, totalVendorProducts)

	// get iterator
	var expectedRecords []types.VendorProducts

	setup.ModelinfoKeeper.IterateVendorProducts(setup.Ctx, func(vendorProducts types.VendorProducts) (stop bool) {
		expectedRecords = append(expectedRecords, vendorProducts)

		return false
	})
	require.Equal(t, expectedLen, len(expectedRecords))
}

func TestKeeper_VendorProductsIteratorOverDifferentVendors(t *testing.T) {
	setup := Setup()

	// add 10 model infos with different Vendors
	count := 10
	PopulateStoreWithModelsHavingDifferentVendor(setup, count)

	// get total count
	totalVendorProducts := setup.ModelinfoKeeper.CountTotalVendorProducts(setup.Ctx)
	require.Equal(t, count, totalVendorProducts)

	// get iterator
	var expectedRecords []types.VendorProducts

	setup.ModelinfoKeeper.IterateVendorProducts(setup.Ctx, func(vendorProducts types.VendorProducts) (stop bool) {
		expectedRecords = append(expectedRecords, vendorProducts)

		return false
	})
	require.Equal(t, count, len(expectedRecords))
}
