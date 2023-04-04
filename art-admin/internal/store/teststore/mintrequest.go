package teststore

import (
	"context"
	"fmt"

	"github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
)

func remove(slice []*pb_nft.NFTMintRequestWithStatus, i int) []*pb_nft.NFTMintRequestWithStatus {
	return append(slice[:i], slice[i+1:]...)
}

func paginate(pageNum int, pageSize int, sliceLength int) (int, int) {
	pageNum = pageNum - 1
	start := pageNum * pageSize

	if start > sliceLength {
		start = sliceLength
	}

	end := start + pageSize
	if end > sliceLength {
		end = sliceLength
	}

	return start, end
}

func (ts *testStore) New(ctx context.Context, mr *pb_nft.NFTMintRequestToUpload, il []*pb_nft.ImageList) (*pb_nft.NFTMintRequestWithStatus, error) {
	oid := getId()
	mws := &pb_nft.NFTMintRequestWithStatus{
		OffchainUrl: "",
		OnchainUrl:  "",
		Status:      pb_nft.Status_Unknown,
		NftMintRequest: &pb_nft.NFTMintRequest{
			Id:                 oid,
			EthAddress:         mr.EthAddress,
			Description:        mr.Description,
			MintSequenceNumber: 0,
		},
		SampleImages: il,
		Collection:   "",
		Duration:     "",
		Shipping: &pb_nft.Shipping{
			Shipping: &pb_nft.ShippingTo{},
		},
	}
	if mr == nil {
		return nil, fmt.Errorf("New:mr is nil")
	}
	ts.mintRequest = append(ts.mintRequest, mws)

	return mws, nil
}

func (ts *testStore) GetById(ctx context.Context, id string) (*pb_nft.NFTMintRequestWithStatus, error) {
	ok := false
	m := &pb_nft.NFTMintRequestWithStatus{}
	for _, mr := range ts.mintRequest {
		if mr.NftMintRequest.Id == id {
			ok = true
			m = mr
			break
		}
	}
	if !ok {
		return nil, fmt.Errorf("no such id")
	}
	return m, nil
}

func (ts *testStore) GetAll(ctx context.Context, status pb_nft.Status) ([]*pb_nft.NFTMintRequestWithStatus, error) {
	mrs := []*pb_nft.NFTMintRequestWithStatus{}
	for _, mr := range ts.mintRequest {
		if mr.Status == status {
			mrs = append(mrs, mr)
		}
	}
	return mrs, nil
}

func (ts *testStore) GetPaged(ctx context.Context, listOpts *pb_nft.ListPagedRequest) ([]*pb_nft.NFTMintRequestWithStatus, error) {
	temp := ts.mintRequest

	if listOpts.Status != pb_nft.Status_Any {
		temp = []*pb_nft.NFTMintRequestWithStatus{}
		for _, mr := range ts.mintRequest {
			if mr.Status.String() == listOpts.Status.String() {
				temp = append(temp, mr)
			}
		}
	}

	start, end := paginate(int(listOpts.Page), ts.pageSize, len(temp))

	return temp[start:end], nil
}

func (ts *testStore) DeleteMintById(ctx context.Context, id string) error {
	for i, mr := range ts.mintRequest {
		if mr.NftMintRequest.Id == id {
			ts.mintRequest = remove(ts.mintRequest, i)
			break
		}
	}
	return nil
}

func (ts *testStore) GetAllToUpload(ctx context.Context) (map[int32]*nft.NFTMintRequestWithStatus, error) {
	mrs := []*pb_nft.NFTMintRequestWithStatus{}
	for _, mr := range ts.mintRequest {
		if mr.Status == pb_nft.Status_Uploaded ||
			mr.Status == pb_nft.Status_Burned ||
			mr.Status == pb_nft.Status_Shipped {
			mrs = append(mrs, mr)
		}
	}
	toUpload := map[int32]*nft.NFTMintRequestWithStatus{}
	for _, mr := range mrs {
		toUpload[mr.NftMintRequest.MintSequenceNumber] = mr
	}
	return toUpload, nil
}

func (ts *testStore) UpdateIpfsUrl(ctx context.Context, id string, ipfsUrl string) (*pb_nft.NFTMintRequestWithStatus, error) {
	num := 0
	mr := &pb_nft.NFTMintRequestWithStatus{}
	for i, m := range ts.mintRequest {
		if m.NftMintRequest.Id == id {
			num = i
			mr = m
		}
	}

	if ipfsUrl == "" {
		return nil, fmt.Errorf("UpdateIpfsUrl:offchainUrl is empty")
	}

	if !((mr.Status == pb_nft.Status_UploadedOffchain) ||
		(mr.Status == pb_nft.Status_Uploaded)) {
		return nil, fmt.Errorf("UpdateIpfsUrl:can change url only for uploaded onchain & offchain")
	}

	mr.OnchainUrl = ipfsUrl
	mr.Status = pb_nft.Status_Uploaded

	ts.mintRequest[num] = mr
	return mr, nil

}
func (ts *testStore) UpdateOffchainUrl(ctx context.Context, id string, offchainUrl string) (*pb_nft.NFTMintRequestWithStatus, error) {
	num := 0
	mr := &pb_nft.NFTMintRequestWithStatus{}
	for i, m := range ts.mintRequest {
		if m.NftMintRequest.Id == id {
			num = i
			mr = m
		}
	}

	if offchainUrl == "" {
		return nil, fmt.Errorf("UpdateOffchainUrl:offchainUrl is empty")
	}

	if !(mr.Status == pb_nft.Status_Pending) {
		return nil, fmt.Errorf("UpdateOffchainUrl:can change url only for uploaded onchain & offchain")
	}

	mr.OffchainUrl = offchainUrl
	mr.Status = pb_nft.Status_UploadedOffchain

	ts.mintRequest[num] = mr
	return mr, nil
}

func (ts *testStore) DeleteIpfsUrl(ctx context.Context, id string) (*pb_nft.NFTMintRequestWithStatus, error) {
	num := 0
	mr := &pb_nft.NFTMintRequestWithStatus{}
	for i, m := range ts.mintRequest {
		if mr.NftMintRequest.Id == id {
			num = i
			mr = m
		}
	}

	if !(mr.Status == pb_nft.Status_Uploaded) {
		return nil, fmt.Errorf("DeleteNFTOnchainUrl:can delete onchain url only for uploaded")
	}

	mr.Status = pb_nft.Status_UploadedOffchain

	ts.mintRequest[num] = mr
	return mr, nil
}

func (ts *testStore) UpdateShippingInfo(ctx context.Context, req *pb_nft.BurnRequest) (*pb_nft.NFTMintRequestWithStatus, error) {
	num := 0
	mr := &pb_nft.NFTMintRequestWithStatus{}
	for i, m := range ts.mintRequest {
		if m.NftMintRequest.Id == req.Id {
			num = i
			mr = m
		}
	}
	if !(mr.Status == pb_nft.Status_Uploaded) {
		return nil, fmt.Errorf("UpdateShippingInfo:can update shipping info only for uploaded")
	}

	mr.Shipping = &pb_nft.Shipping{
		Shipping: req.Shipping,
	}
	mr.Status = pb_nft.Status_Burned

	ts.mintRequest[num] = mr
	return mr, nil
}

func (ts *testStore) UpdateTrackingNumber(ctx context.Context, req *pb_nft.SetTrackingNumberRequest) (*pb_nft.NFTMintRequestWithStatus, error) {
	num := 0
	mr := &pb_nft.NFTMintRequestWithStatus{}
	for i, m := range ts.mintRequest {
		if m.NftMintRequest.Id == req.Id {
			num = i
			mr = m
		}
	}

	if !(mr.Status == pb_nft.Status_Burned) {
		return nil, fmt.Errorf("UpdateTrackingNumber:can update tracking number only for burned")
	}

	mr.Shipping.TrackingNumber = req.TrackingNumber
	mr.Status = pb_nft.Status_Shipped

	ts.mintRequest[num] = mr
	return mr, nil
}

func (ts *testStore) UpdateStatus(ctx context.Context, id string, status pb_nft.Status) (*pb_nft.NFTMintRequestWithStatus, error) {
	ok := false
	m := &pb_nft.NFTMintRequestWithStatus{}
	for _, mr := range ts.mintRequest {
		if mr.NftMintRequest.Id == id {
			ok = true
			mr.Status = status
			m = mr
			break
		}
	}
	if !ok {
		return nil, fmt.Errorf("no such id")
	}
	return m, nil
}

func newMr(id int, st pb_nft.Status) *pb_nft.NFTMintRequest {
	return &pb_nft.NFTMintRequest{
		Id:                 fmt.Sprintf("%s:%d", st.String(), id),
		EthAddress:         "0x1234567890123456789012345678901234567890",
		MintSequenceNumber: int32(id),
		Description:        fmt.Sprintf("description:%d", id),
	}
}

func newImages() []*pb_nft.ImageList {
	return []*pb_nft.ImageList{
		{
			FullSize: "https://mint.com/img.jpg",
		},
		{
			FullSize: "https://mint2.com/img.jpg",
		},
	}
}
func newShipping(trackN string) *pb_nft.Shipping {
	return &pb_nft.Shipping{
		Shipping: &pb_nft.ShippingTo{
			FullName: "John",
			Address:  "123 Main St",
			City:     "San Francisco",
			ZipCode:  "94105",
			Country:  "USA",
			Email:    "test@solutions.com",
		},
		TrackingNumber: trackN,
	}
}

func (ts *testStore) AddMockData(sts []pb_nft.Status, amountPerStatus int) {
	totalAmount := len(sts) * amountPerStatus
	ts.mintRequest = make([]*pb_nft.NFTMintRequestWithStatus, totalAmount+1)
	for i, st := range sts {
		for j := 0; j <= amountPerStatus; j++ {
			// fmt.Println("adding", st, j+1)
			ts.mintRequest[(i*amountPerStatus)+j] = GetMockMrByStatus(st, j+1)
		}
	}
	ts.mintRequest = append(ts.mintRequest)
}

func GetMockMrByStatus(st pb_nft.Status, num int) *pb_nft.NFTMintRequestWithStatus {
	switch st {
	case pb_nft.Status_Unknown:
		return &pb_nft.NFTMintRequestWithStatus{
			Status:         pb_nft.Status_Unknown,
			NftMintRequest: newMr(num, st),
			SampleImages:   newImages(),
		}

	case pb_nft.Status_Pending:
		return &pb_nft.NFTMintRequestWithStatus{
			Status:         pb_nft.Status_Pending,
			NftMintRequest: newMr(num, st),
			SampleImages:   newImages(),
		}

	case pb_nft.Status_UploadedOffchain:
		return &pb_nft.NFTMintRequestWithStatus{
			OffchainUrl:    fmt.Sprintf("https://offchain.com/%d", num),
			Status:         pb_nft.Status_UploadedOffchain,
			NftMintRequest: newMr(num, st),
			SampleImages:   newImages(),
		}

	case pb_nft.Status_Uploaded:
		return &pb_nft.NFTMintRequestWithStatus{
			OffchainUrl:    fmt.Sprintf("https://offchain.com/%d", num),
			Status:         pb_nft.Status_Uploaded,
			NftMintRequest: newMr(num, st),
			SampleImages:   newImages(),
		}

	case pb_nft.Status_Burned:
		return &pb_nft.NFTMintRequestWithStatus{
			OffchainUrl:    fmt.Sprintf("https://offchain.com/%d", num),
			Status:         pb_nft.Status_Burned,
			NftMintRequest: newMr(num, st),
			SampleImages:   newImages(),
			Shipping:       newShipping(""),
		}

	case pb_nft.Status_Shipped:
		return &pb_nft.NFTMintRequestWithStatus{
			OffchainUrl:    fmt.Sprintf("https://offchain.com/%d", num),
			Status:         pb_nft.Status_Shipped,
			NftMintRequest: newMr(num, st),
			SampleImages:   newImages(),
			Shipping:       newShipping("228"),
		}
	}
	return nil
}
