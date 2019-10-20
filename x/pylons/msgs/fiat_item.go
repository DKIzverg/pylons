package msgs

import (
	"encoding/json"

	"github.com/MikeSofaer/pylons/x/pylons/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgFiatItem struct {
	// TODO: ID of item should be available in this Msg
	// TODO: MsgFiatItem & types.Item is same, do we need to use these differently? discuss discord
	Doubles    []types.DoubleKeyValue
	Longs      []types.LongKeyValue
	Strings    []types.StringKeyValue
	CookbookID string
	Sender     sdk.AccAddress
}

// NewMsgFiatItem a constructor for MsgFiatItem msg
func NewMsgFiatItem(cookbookID string, doubles []types.DoubleKeyValue, longs []types.LongKeyValue, strings []types.StringKeyValue, sender sdk.AccAddress) MsgFiatItem {
	return MsgFiatItem{
		CookbookID: cookbookID,
		Doubles:    doubles,
		Longs:      longs,
		Strings:    strings,
		Sender:     sender,
	}
}

// Route should return the name of the module
func (msg MsgFiatItem) Route() string { return "pylons" }

// Type should return the action
func (msg MsgFiatItem) Type() string { return "fiat_item" }

// ValidateBasic validates the Msg
func (msg MsgFiatItem) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress(msg.Sender.String())
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgFiatItem) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners gets the signer who should have signed the message
func (msg MsgFiatItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
