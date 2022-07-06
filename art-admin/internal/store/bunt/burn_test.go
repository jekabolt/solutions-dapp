package bunt

import (
	"fmt"
	"testing"

	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/matryer/is"
)

func TestBurn(t *testing.T) {
	is := is.New(t)

	c := Config{
		DBPath: ":memory:",
	}

	bdb, err := c.InitDB()
	is.NoErr(err)

	bs := bdb.BurnStore()

	err = bs.BurnNft(&pb_nft.BurnRequest{
		Txid:               "0x0",
		Address:            "0x0",
		MintSequenceNumber: 1,
		Shipping: &pb_nft.ShippingTo{
			Email: "test@mail.com",
		},
	})
	is.NoErr(err)

	bsis, err := bs.GetBurned()
	is.NoErr(err)
	is.True(len(bsis) == 1)

	bsis, err = bs.GetBurnedErrors()
	is.NoErr(err)
	is.True(len(bsis) == 0)

	bsis, err = bs.GetBurnedPending()
	is.NoErr(err)
	is.True(len(bsis) == 1)

	err = bs.UpdateShippingStatus(&pb_nft.ShippingStatusUpdateRequest{
		Id: fmt.Sprint(bsis[0].Id),
		Status: &pb_nft.ShippingStatus{
			Error: "test err",
		},
	})
	is.NoErr(err)

	bsis, err = bs.GetBurnedErrors()
	is.NoErr(err)
	is.True(len(bsis) == 1)

	err = bs.UpdateShippingStatus(&pb_nft.ShippingStatusUpdateRequest{
		Id: fmt.Sprint(bsis[0].Id),
		Status: &pb_nft.ShippingStatus{
			Error: "",
		},
	})
	is.NoErr(err)

	bsis, err = bs.GetBurnedPending()
	is.NoErr(err)
	is.True(len(bsis) == 1)

	err = bs.UpdateShippingStatus(&pb_nft.ShippingStatusUpdateRequest{
		Id: fmt.Sprint(bsis[0].Id),
		Status: &pb_nft.ShippingStatus{
			Success: true,
		},
	})
	is.NoErr(err)

	bsis, err = bs.GetBurnedPending()
	is.NoErr(err)
	is.True(len(bsis) == 0)
}
