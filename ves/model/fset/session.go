package fset

import (
	"fmt"
	"github.com/HyperService-Consortium/NSB/contract/isc/TxState"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/lib/encoding"
	"github.com/Myriad-Dreamin/go-ves/lib/upstream"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
)

// SessionFSet is the collection of functions related to model.session
type SessionFSet struct {
	AccountDB *model.SessionAccountDB
	Index     types.Index
}

func NewSessionFSet(provider *model.Provider, index types.Index) *SessionFSet {
	return &SessionFSet{AccountDB: provider.SessionAccountDB(), Index: index}
}

func (s SessionFSet) GetAccounts(ses *model.Session) ([]uiptypes.Account, error) {
	accounts, err := s.AccountDB.ID(ses.ISCAddress)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSessionAccountFindError, err)
	}
	return model.SessionAccountsToUIPAccounts(accounts), nil
}

func (s SessionFSet) GetAckCount(ses *model.Session) (int64, error) {
	return s.AccountDB.GetAcknowledged(ses.ISCAddress)
}

func (s SessionFSet) FindTransaction(
	iscAddress []byte,
	transactionID int64) (b []byte, err error) {
	tx := new(model.Transaction)
	has, err := tx.FindSessionIndex(model.EncodeAddress(iscAddress), transactionID)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSelectError, err)
	} else if !has {
		return nil, wrapper.WrapCode(types.CodeNotFound)
	}

	return model.DecodeContent(tx.Content), nil
}

//?
func (s SessionFSet) GetTransactingTransaction(ses *model.Session) ([]byte, error) {
	//return ses.UnderTransacting
	return s.FindTransaction(ses.GetGUID(), ses.UnderTransacting)
}

func (s SessionFSet) InitSessionInfo(
	iscAddress []byte, intents []*opintent.TransactionIntent, accounts []*model.SessionAccount) (
	ses *model.Session, err error) {
	ses = model.NewSession(iscAddress)
	for i := range accounts {
		accounts[i].SessionID = ses.ISCAddress
		if _, err = accounts[i].Create(); err != nil {
			err = wrapper.Wrap(types.CodeSessionInsertAccountError, err)
			return
		}
	}

	ses.AccountsCount, err = s.AccountDB.GetTotal(ses.ISCAddress)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSessionAccountGetTotalError, err)
	}

	ses.TransactionCount = int64(len(intents))
	for i := range intents {
		b, err := upstream.Serializer.TransactionIntent.Marshal(intents[i])
		if err != nil {
			return nil, wrapper.Wrap(types.CodeTransactionIntentSerializeError, err)
		}

		if _, err = (&model.Transaction{
			SessionID: ses.ISCAddress,
			Index:     int64(i),
			Content:   model.EncodeContent(b),
		}).Create(); err != nil {
			return nil, wrapper.Wrap(types.CodeSessionInsertTransactionError, err)
		}
	}

	_, err = ses.Create()
	return
}

func (s SessionFSet) AckForInit(ses *model.Session, acc uiptypes.Account, signature uiptypes.Signature) error {
	var (
		sac = &model.SessionAccount{
			SessionID: ses.ISCAddress,
			ChainID:   acc.GetChainId(),
			Address:   encoding.EncodeBase64(acc.GetAddress()),
		}
		has bool
		err error
	)
	//verifySignature

	if has, err = sac.Find(); err != nil {
		return wrapper.Wrap(types.CodeSessionAccountFindError, err)
	} else if !has {
		return wrapper.Wrap(types.CodeSessionAccountNotFound, err)
	}

	sac.Acknowledged = true
	if aff, err := sac.UpdateAcknowledged(); err != nil {
		return wrapper.Wrap(types.CodeUpdateError, err)
	} else if aff == 0 {
		return wrapper.WrapCode(types.CodeUpdateNoEffect)
	}
	return nil
}

type NotCurrentTransactionError struct {
	Requiring int64 `json:"requiring"`
	Current   int64 `json:"current"`
}

func (n NotCurrentTransactionError) Error() string {
	return fmt.Sprintf("requiring:%v,current:%v", n.Requiring, n.Current)
}

func (s SessionFSet) NotifyAttestation(
	ses *model.Session, nsb types.NSBInterface,
	bn uiptypes.BlockChainInterface, attestation uiptypes.Attestation) (err error) {
	if attestation.GetTid() != uint64(ses.UnderTransacting) {
		return wrapper.Wrap(types.CodeSessionNotCurrentTransaction, NotCurrentTransactionError{
			Requiring: int64(attestation.GetTid()),
			Current:   ses.UnderTransacting,
		})
	}
	switch attestation.GetAid() {
	case TxState.Unknown, TxState.Initing, TxState.Inited:
		return nil
	case TxState.Instantiating, TxState.Instantiated, TxState.Open, TxState.Opened:
		return nil
	case TxState.Closed:
		ses.UnderTransacting++
		if ses.UnderTransacting == ses.TransactionCount {
			err = nsb.SettleContract(ses.GetGUID())
			if err != nil {
				return wrapper.Wrap(types.CodeSettleContractError, err)
			}
		}
		return nil
	default:
		return wrapper.WrapCode(types.CodeTransactionStateNotFound)
	}
}

func (s SessionFSet) ProcessAttestation(types.NSBInterface, uiptypes.BlockChainInterface, uiptypes.Attestation) (interface{}, interface{}, error) {
	panic("implement me")
}

func (s SessionFSet) SyncFromISC() error {
	panic("implement me")
}
