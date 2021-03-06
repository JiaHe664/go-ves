package bni

import (
	"fmt"
	"github.com/HyperService-Consortium/NSB/application/nsb_proto"
	nsbcli "github.com/HyperService-Consortium/NSB/lib/nsb-client"
	"github.com/HyperService-Consortium/NSB/merkmap"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	merkleproof "github.com/HyperService-Consortium/go-uip/merkle-proof"
	"github.com/HyperService-Consortium/go-uip/storage"
	"github.com/HyperService-Consortium/go-uip/uip"
)

func (bn *BN) GetTransactionProof(chainID uint64, blockID []byte, additional []byte) (uip.MerkleProof, error) {
	ci, err := bn.dns.GetChainInfo(chainID)

	if err != nil {
		return nil, err
	}

	resp, err := nsbcli.NewNSBClient(ci.GetChainHost()).GetProof(additional, `prove_on_tx_trie`)

	if err != nil {
		return nil, err
	}

	return merkleproof.NewMPTUsingKeccak256(resp.Proof, resp.Key, resp.Value), nil
}

func (bn *BN) GetStorageAt(chainID uip.ChainID, typeID uip.TypeID, contractAddress uip.ContractAddress, pos []byte, description []byte) (uip.Variable, error) {
	ci, err := bn.dns.GetChainInfo(chainID)
	if err != nil {
		return nil, err
	}

	return newTendermintNSBStorageHandler(ci.GetChainHost()).GetStorageAt(
		typeID, contractAddress, pos, description)
}

type NSBClient interface {
	GetStorageProofAt(
		args nsb_proto.ArgsGetStorageAt) (*merkmap.ProofJson, error)
}


func newTendermintNSBStorageHandler(host string) tendermintNSBStorageHandler {
	return tendermintNSBStorageHandler{
		handler: nsbcli.NewNSBClient(host),
	}
}

type tendermintNSBStorageHandler struct {
	handler NSBClient
}

func (t tendermintNSBStorageHandler) GetTransactionProof(blockID uip.BlockID, color []byte) (uip.MerkleProof, error) {
	panic("implement me")
}

func (t tendermintNSBStorageHandler) GetStorageAt(typeID uip.TypeID, contractAddress uip.ContractAddress, pos []byte, description []byte) (uip.Variable, error) {
	resp, err := t.handler.GetStorageProofAt(
		nsb_proto.ArgsGetStorageAt{
			Address: contractAddress,
			Key:     description,
			Slot:    pos,
		})
	if err != nil {
		return nil, err
	}

	return tendermintNSBStorageBytesToValue(resp.Value, typeID)
}

func tendermintNSBStorageBytesToValue(value []byte, id uip.TypeID) (uip.Variable, error) {
	var decoder interface {
		Decode([]byte) (interface{}, error)
	}
	switch id {
	case value_type.Uint8:
		decoder = storage.Uint8
	case value_type.Uint16:
		decoder = storage.Uint16
	case value_type.Uint32:
		decoder = storage.Uint32
	case value_type.Uint64:
		decoder = storage.Uint32
	case value_type.Uint128:
		decoder = storage.Uint128
	case value_type.Uint256:
		decoder = storage.Uint256
	case value_type.Int8:
		decoder = storage.Int8
	case value_type.Int16:
		decoder = storage.Int16
	case value_type.Int32:
		decoder = storage.Int32
	case value_type.Int64:
		decoder = storage.Int32
	case value_type.Int128:
		decoder = storage.Int128
	case value_type.Int256:
		decoder = storage.Int256
	case value_type.String:
		decoder = storage.String
	case value_type.Bytes:
		decoder = storage.Bytes
	case value_type.Bool:
		decoder = storage.Bool
	case value_type.Unknown:
		if len(value) == 0 {
			return &uip.VariableImpl{Type: id, Value: nil}, nil
		}
	}
	if decoder == nil {
		return nil, fmt.Errorf("invalid value type: %v", id)
	}
	return refWrapper(id).wrapValue(decoder.Decode(value))
}

type refWrapper uip.TypeID

func (id refWrapper) wrapValue(value interface{}, err error) (uip.Variable, error) {
	if err != nil {
		return nil, err
	}
	return uip.VariableImpl{Value: value, Type: uip.TypeID(id)}, nil
}

