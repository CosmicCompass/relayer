package cmd

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"gitlab.ctlabs.in/saikrishna/freeflix-media-hub/modules/ibc/xnfts"
)

func RegisterCode(cdc *codec.Codec) {
	xnfts.RegisterCodec(cdc)
}
