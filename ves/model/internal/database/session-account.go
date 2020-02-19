package database

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
)

type SessionAccount struct {
	SessionID    string                         `dorm:"session_id" gorm:"column:session_id;not_null" json:"-"`
	ChainID      uiptypes.ChainIDUnderlyingType `dorm:"chain_id" gorm:"column:chain_id;not_null" json:"chain_id"`
	Address      string                         `dorm:"address" gorm:"column:address;not_null" json:"address"`
	Acknowledged bool                           `dorm:"acknowledged" gorm:"column:acknowledged;not_null" json:"acknowledged"`
}

func NewSessionAccount(chainID uiptypes.ChainIDUnderlyingType, address []byte) *SessionAccount {
	return &SessionAccount{
		ChainID: chainID,
		Address: EncodeAddress(address),
	}
}

// TableName specification
func (SessionAccount) TableName() string {
	return "session_account"
}

func (sa SessionAccount) GetID() uint {
	panic("aborted")
}

func (sa SessionAccount) GetChainId() uiptypes.ChainID {
	return sa.ChainID
}

// todo
func (sa SessionAccount) GetAddress() uiptypes.Address {
	//if sa.decodedAddress == nil {
	//	sa.decodedAddress = decodeBase64(sa.Address)
	//}
	return sugar.HandlerError(DecodeAddress(sa.Address)).([]byte)
}

type SessionAccountFilter = Filter

func SessionAccountsToUIPAccounts(accounts []SessionAccount) (uipAccounts []uiptypes.Account) {
	uipAccounts = make([]uiptypes.Account, len(accounts))
	for i := range accounts {
		uipAccounts[i] = accounts[i]
	}
	return uipAccounts
}
