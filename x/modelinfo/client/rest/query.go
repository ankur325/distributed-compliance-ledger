package rest

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"net/http"
)

func getModelsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		params, err := restCtx.ParsePaginationParams()
		if err != nil {
			return
		}

		restCtx.QueryList(fmt.Sprintf("custom/%s/all_models", storeName), params)
	}
}

func getModelHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		vid := vars[vid]
		pid := vars[pid]

		res, height, err := restCtx.QueryStore(keeper.ModelInfoId(vid, pid), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound, types.ErrModelInfoDoesNotExist(vid, pid).Error())
			return
		}

		var modelInfo types.ModelInfo
		cliCtx.Codec.MustUnmarshalBinaryBare(res, &modelInfo)

		restCtx.EncodeAndRespondWithHeight(modelInfo, height)
	}
}

func getVendorsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		params, err := restCtx.ParsePaginationParams()
		if err != nil {
			return
		}

		restCtx.QueryList(fmt.Sprintf("custom/%s/vendors", storeName), params)
	}
}

func getVendorModelsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		vid := vars[vid]

		res, height, err := restCtx.QueryStore(keeper.VendorProductsId(vid), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound, types.ErrVendorProductsDoNotExist(vid).Error())
			return
		}

		var vendorProducts types.VendorProducts
		cliCtx.Codec.MustUnmarshalBinaryBare(res, &vendorProducts)

		restCtx.EncodeAndRespondWithHeight(vendorProducts, height)
	}
}