package systoken

import (
	"math/big"
	"strings"

	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
)

// isPaid checks if the mint is paid
func (sys *Systoken) IsPaid(mr *pb_nft.NFTMintRequestWithStatus) (bool, error) {
	_, err := sys.OwnerById(nil, big.NewInt(int64(mr.NftMintRequest.GetMintSequenceNumber())))
	if err != nil {
		if strings.Contains(err.Error(), "invalid token ID") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
