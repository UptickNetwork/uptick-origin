package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Supply return the amount of the denom
func (idc IDCollection) Supply() int {
	return len(idc.TokenIds)
}

// AddID adds an tokenID to the idCollection
func (idc IDCollection) AddID(tokenID string) IDCollection {
	idc.TokenIds = append(idc.TokenIds, tokenID)
	return idc
}

// ----------------------------------------------------------------------------
// IDCollections is an array of Id Collections
type IDCollections []IDCollection

// Add adds an Id to the idCollection
func (idcs IDCollections) Add(DenomId, tokenID string) IDCollections {
	for i, idc := range idcs {
		if idc.DenomId == DenomId {
			idcs[i] = idc.AddID(tokenID)
			return idcs
		}
	}
	return append(idcs, IDCollection{
		DenomId:  DenomId,
		TokenIds: []string{tokenID},
	})
}

// String follows stringer interface
func (idcs IDCollections) String() string {
	if len(idcs) == 0 {
		return ""
	}

	var buf bytes.Buffer
	for _, idCollection := range idcs {
		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(idCollection.String())
	}
	return buf.String()
}

// NewOwner creates a new Owner
func NewOwner(owner sdk.AccAddress, idCollections ...IDCollection) Owner {
	return Owner{
		Address:       owner.String(),
		IDCollections: idCollections,
	}
}

type Owners []Owner

// NewOwner creates a new Owner
func NewOwners(owner ...Owner) Owners {
	return append([]Owner{}, owner...)
}

// String follows stringer interface
func (owners Owners) String() string {
	var buf bytes.Buffer
	for _, owner := range owners {
		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(owner.String())
	}
	return buf.String()
}
