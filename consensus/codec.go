package consensus

import (
	go_amino "github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
)

func RegisterCodec(cdc *go_amino.Codec) {
	cdc.RegisterConcrete(&abci.ConsensusParams{}, "abci/consensus/ConsensusParams", nil)
}
