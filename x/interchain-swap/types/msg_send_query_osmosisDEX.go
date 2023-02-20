package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendQueryOsmosisDex = "send_query_to_OsmosisDex"

var _ sdk.Msg = &QuerySpotPriceStruct{}

func NewMsgSendQueryAllBalances(creator, channelId string, poolID string, baseAssetDenom string, quoteAssetDenom string) *QuerySpotPriceStruct {
	return &QuerySpotPriceStruct{
		Creater:         creator,
		ChannelID:       channelId,
		PoolID:          poolID,
		BaseAssetDenom:  baseAssetDenom,
		QuoteAssetDenom: quoteAssetDenom,
	}
}

func (msg *QuerySpotPriceStruct) Route() string {
	return RouterKey
}

func (msg *QuerySpotPriceStruct) Type() string {
	return TypeMsgSendQueryOsmosisDex
}

func (msg *QuerySpotPriceStruct) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creater)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *QuerySpotPriceStruct) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *QuerySpotPriceStruct) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creater)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
