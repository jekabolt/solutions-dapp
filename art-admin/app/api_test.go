package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jekabolt/solutions-dapp/art-admin/auth"
	"github.com/jekabolt/solutions-dapp/art-admin/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/config"
	"github.com/jekabolt/solutions-dapp/art-admin/store/bunt"
	"github.com/jekabolt/solutions-dapp/art-admin/store/nft"
	"github.com/matryer/is"
)

var (
	S3AccessKey       = "xxx"
	S3SecretAccessKey = "xxx"
	S3Endpoint        = "fra1.digitaloceanspaces.com"
	bucketName        = "grbpwr"
	bucketLocation    = "fra-1"
	imageStorePrefix  = "grbpwr-com"

	serverPort = "8080"

	jwtSecret   = "jwtSecret"
	adminSecret = "adminSecret"
	hosts       = []string{"*"}
)

func BucketFromConst() (*bucket.Bucket, error) {
	bucketConf := &bucket.Config{
		S3AccessKey:       S3AccessKey,
		S3SecretAccessKey: S3SecretAccessKey,
		S3Endpoint:        S3Endpoint,
		S3BucketName:      bucketName,
		S3BucketLocation:  bucketLocation,
		ImageStorePrefix:  imageStorePrefix,
	}
	return bucketConf.Init()
}

func buntFromConst() (*bunt.BuntDB, error) {
	c := &bunt.Config{
		DBPath: ":memory:",
	}
	return c.InitDB()
}

func InitServerFromConst() (*Server, error) {
	bucket, err := BucketFromConst()
	if err != nil {
		return nil, err
	}

	b, err := buntFromConst()
	if err != nil {
		return nil, err
	}

	ac := &auth.Config{
		AdminSecret: adminSecret,
		JWTSecret:   jwtSecret,
	}

	return &Server{
		nftStore:      b.NFTStore(),
		metadataStore: b.MetadataStore(),
		Bucket:        bucket,
		Auth:          ac.New(),
		Config: &config.Config{
			Port:  serverPort,
			Hosts: hosts,
			Auth:  ac,
			Debug: true,
		},
	}, err
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader, response interface{}, at string) (*http.Response, interface{}) {

	is := is.New(t)

	req, err := http.NewRequest(method, ts.URL+path, body)
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", at))

	resp, err := http.DefaultClient.Do(req)
	is.NoErr(err)

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(response)
	is.NoErr(err)

	return resp, response
}

func (s *Server) getAuthRequestPw() *bytes.Reader {
	a := auth.AuthRequest{
		Password: s.Auth.AdminSecret,
	}
	aBytes, _ := json.Marshal(a)
	return bytes.NewReader(aBytes)
}

func (s *Server) getAuthRequestWrongPw() *bytes.Reader {
	a := auth.AuthRequest{
		Password: s.Auth.AdminSecret + "wrong",
	}
	aBytes, _ := json.Marshal(a)
	return bytes.NewReader(aBytes)
}

func (s *Server) getAuthRequestRefresh(rt string) *bytes.Reader {
	a := auth.AuthRequest{
		RefreshToken: rt,
	}
	aBytes, _ := json.Marshal(a)
	return bytes.NewReader(aBytes)
}

func (s *Server) getAuthRequestWrongRefresh(rt string) *bytes.Reader {
	a := auth.AuthRequest{
		RefreshToken: rt + "wrong",
	}
	aBytes, _ := json.Marshal(a)
	return bytes.NewReader(aBytes)

}

func TestAuthTokenByPasswordAndRefresh(t *testing.T) {
	is := is.New(t)

	s, err := InitServerFromConst()
	is.NoErr(err)

	ts := httptest.NewServer(s.Router())
	defer ts.Close()

	// auth w password
	authResp := &AuthResponse{}
	res, ar := testRequest(t, ts, http.MethodPost, "/auth", s.getAuthRequestPw(), authResp, "")
	authResp = ar.(*AuthResponse)
	is.Equal(res.StatusCode, http.StatusOK)
	// t.Logf("%+v", authResp)

	// auth w refresh
	res, ar = testRequest(t, ts, http.MethodPost, "/auth", s.getAuthRequestRefresh(authResp.RefreshToken), authResp, "")
	authResp = ar.(*AuthResponse)
	is.Equal(res.StatusCode, http.StatusOK)
	t.Logf("%+v", authResp)

	// auth with wrong password
	authResp = &AuthResponse{}
	res, ar = testRequest(t, ts, http.MethodPost, "/auth", s.getAuthRequestWrongPw(), authResp, "")
	authResp = ar.(*AuthResponse)
	is.Equal(res.StatusCode, http.StatusUnauthorized)
	// t.Logf("%+v", authResp)

	// auth w  wrong refresh
	res, ar = testRequest(t, ts, http.MethodPost, "/auth", s.getAuthRequestWrongRefresh(authResp.RefreshToken), authResp, "")
	authResp = ar.(*AuthResponse)
	is.Equal(res.StatusCode, http.StatusUnauthorized)

	t.Logf("%+v", authResp)
}

func getNFTMintRequestReq(t *testing.T, ethAddr string) *bytes.Reader {
	is := is.New(t)
	nft := &nft.NFTMintRequest{
		Id:         0,
		ETHAddress: ethAddr,
		TxHash:     "0x0",
		SampleImages: []bucket.Image{
			{
				FullSize: "https://ProductImages.com/img.jpg",
			},
			{
				FullSize: "https://ProductImages2.com/img.jpg",
			},
		},
		MintSequenceNumber: 1,
		Description:        "test",
		Status:             nft.StatusUnknown,
	}

	nftBytes, err := json.Marshal(nft)

	// t.Logf("%s", nftBytes)
	is.NoErr(err)

	return bytes.NewReader(nftBytes)
}

func NFTMintRequestFromRespReq(t *testing.T, nft nft.NFTMintRequest) *bytes.Reader {
	is := is.New(t)

	nftBytes, err := json.Marshal(nft)

	t.Logf("%s", nftBytes)
	is.NoErr(err)

	return bytes.NewReader(nftBytes)

}

func TestNFTMintRequestCRUDWAuth(t *testing.T) {
	is := is.New(t)

	s, err := InitServerFromConst()
	is.NoErr(err)

	ts := httptest.NewServer(s.Router())
	defer ts.Close()

	// jwt token
	authData, err := s.Auth.GetJWT()
	is.NoErr(err)
	// t.Log(authData)

	// add nft mint
	nftResp := &NFTMintResponse{}
	ethAddr1 := "0x1"
	res, pr := testRequest(t, ts, http.MethodPut, "/api/mint/requests", getNFTMintRequestReq(t, ethAddr1), nftResp, authData.AccessToken)
	nftResp = pr.(*NFTMintResponse)
	is.Equal(res.StatusCode, http.StatusOK)
	is.Equal(nftResp.NFTMintRequest.ETHAddress, ethAddr1)

	// get all
	allNFTMintRequestResp := &[]NFTMintResponse{}

	res, ar := testRequest(t, ts, http.MethodGet, "/api/mint/requests", nil, allNFTMintRequestResp, authData.AccessToken)
	allNFTMintRequestResp = ar.(*[]NFTMintResponse)
	is.Equal(res.StatusCode, http.StatusOK)
	is.Equal(len(*allNFTMintRequestResp), 1)

	// add nft
	nftWithOffchain := nftResp.NFTMintRequest
	nftWithOffchain.NFTOffchain = "some offchain url"

	res, pr = testRequest(t, ts, http.MethodPut, "/api/nft", NFTMintRequestFromRespReq(t, *nftWithOffchain), nftResp, authData.AccessToken)
	is.Equal(res.StatusCode, http.StatusOK)

	// get all
	allNFTMintRequestResp = &[]NFTMintResponse{}
	res, ar = testRequest(t, ts, http.MethodGet, "/api/mint/requests", nil, allNFTMintRequestResp, authData.AccessToken)
	allNFTMintRequestResp = ar.(*[]NFTMintResponse)
	is.Equal(res.StatusCode, http.StatusOK)
	is.Equal(len(*allNFTMintRequestResp), 1)
	all := *allNFTMintRequestResp
	is.Equal(all[0].NFTMintRequest.NFTOffchain, nftWithOffchain.NFTOffchain)

	// delte nft
	res, pr = testRequest(t, ts, http.MethodDelete, fmt.Sprintf("/api/nft/%d", all[0].NFTMintRequest.Id), NFTMintRequestFromRespReq(t, *nftWithOffchain), nftResp, authData.AccessToken)
	is.Equal(res.StatusCode, http.StatusOK)

	// get all
	allNFTMintRequestResp = &[]NFTMintResponse{}
	res, ar = testRequest(t, ts, http.MethodGet, "/api/mint/requests", nil, allNFTMintRequestResp, authData.AccessToken)
	allNFTMintRequestResp = ar.(*[]NFTMintResponse)
	is.Equal(res.StatusCode, http.StatusOK)
	is.Equal(len(*allNFTMintRequestResp), 1)
	all = *allNFTMintRequestResp
	is.Equal(all[0].NFTMintRequest.NFTOffchain, "")

	// delete mint by id
	res, pr = testRequest(t, ts, http.MethodDelete, fmt.Sprintf("/api/mint/requests/%d", all[0].NFTMintRequest.Id), NFTMintRequestFromRespReq(t, *nftWithOffchain), nftResp, authData.AccessToken)
	is.Equal(res.StatusCode, http.StatusOK)

	allNFTMintRequestResp = &[]NFTMintResponse{}
	res, ar = testRequest(t, ts, http.MethodGet, "/api/mint/requests", nil, allNFTMintRequestResp, authData.AccessToken)
	allNFTMintRequestResp = ar.(*[]NFTMintResponse)
	is.Equal(res.StatusCode, http.StatusOK)
	is.Equal(len(*allNFTMintRequestResp), 0)

	t.Logf("%+v", allNFTMintRequestResp)
}
