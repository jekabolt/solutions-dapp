package teststore

import (
	"context"
	"fmt"
	"testing"

	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/matryer/is"
)

func TestXxx(t *testing.T) {
	is := is.New(t)
	s := NewTestStore(30)

	status := pb_nft.Status_Unknown
	s.AddMockData([]pb_nft.Status{
		status,
	}, 39)

	mrs, err := s.GetPaged(context.Background(), &pb_nft.ListPagedRequest{
		Status: status,
		Page:   1,
	})
	is.NoErr(err)
	fmt.Println("mrs", len(mrs))
	// for _, mr := range mrs {
	// 	fmt.Println("- mr ", mr)
	// }
}
