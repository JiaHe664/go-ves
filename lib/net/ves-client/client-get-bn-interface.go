package vesclient

import (
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"
	"github.com/HyperService-Consortium/go-uip/uiptypes"

	ethbni "github.com/Myriad-Dreamin/go-ves/lib/bni/eth"
	nsbbni "github.com/Myriad-Dreamin/go-ves/lib/bni/ten"
)

func (vc *VesClient) getRouter(chainID uint64) (uiptypes.Router, error) {
	if ci, err := vc.dns.GetChainInfo(chainID); err != nil {
		return nil, err
	} else {
		switch ci.GetChainType() {
		case ChainType.Ethereum:
			return ethbni.NewBN(vc.dns), nil
		case ChainType.TendermintNSB:
			return nsbbni.NewBN(vc.dns), nil
		default:
			return nil, wrapCode(CodeUnknownChainID)
		}
	}
}

func (vc *VesClient) getBlockStorage(chainID uint64) (uiptypes.Storage, error) {
	if ci, err := vc.dns.GetChainInfo(chainID); err != nil {
		return nil, err
	} else {
		switch ci.GetChainType() {
		case ChainType.Ethereum:
			return ethbni.NewBN(vc.dns), nil
		case ChainType.TendermintNSB:
			return nsbbni.NewBN(vc.dns), nil
		default:
			return nil, wrapCode(CodeUnknownChainID)
		}
	}
}
