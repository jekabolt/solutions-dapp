package ipfs

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"

	"github.com/jekabolt/solutions-dapp/art-admin/descriptions"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/jekabolt/solutions-dapp/art-admin/store/bunt"
	"github.com/matryer/is"
)

const (
	testImage = "data:image/jpeg;base64,/9j/2wCEAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAf/AABEIABQAFAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/APMfjJ+y1+zf8Nf+CCHws/aP8IfsK+OPFv7QXxg+G3hv4reN/wBtbwTCxvfhPMv7RXgkNJ43m8XfG1PHMfjP4g+Divg94/AltcSPLcZM8CJHG/ofxk+H/wDwTS+NUH7D3i79k3/gkJ8erXwb4s/b5+Gnwv8AjJ4Jf4kfBaLxx48t4fg18cFu/wBn/wCIPwjuP2y77xJ8OPH8E/hdt RM3xnt/hbZQR/DeWGfxjOZo4Lj4jb4J+LfjL/wSe+GPw4/Z4+KH/Bbj9oX9ofxTb/De0f8AZnnf4un/AIJ8ww/8LlFxdf8ACvvGL/Bv/hX1vpkSKk/hCQfFtoDOrSCRZrdUPHfFf9oH9mfxB8Xfh9+0HH+3P/wU4+LV78M/2uDb2H7SHhb4lOfj9Y/8E5b34CWl0L/whIvgzRoo/iMf2kB8TfhxPc3cjQeFrbwxciaOBrhb+0APmr/gptexfs6/txfHD4OfA7SP2nf+Ce/w58HD4Yf2N+zP4f8A2j9ckbw43iH4M/DvxVP4l15vCXxd+JPhn/hI/GR1xde1NdI8XaltF1ANVKa9/au74O/4Xx8WP+jxP2vv/Ei/Hv8A80Vfdf7WXgvwfe/tZftCf8JX8cf23NEaHX/hU3h3/hte+0T49/tP6x4V1P8AZs+B+vaPr3xR+Jnwr1nx94H1w6murTy+Bv7H8U6ksPwzXwWZmimkaJfBv+FffCP/AKON8bf+Go8Uf/KagD9NfBP/AAUh/wCCiH7CNn4h/ZY/Z1/bG8eeHvgd8ANYHhP4Z+D/ABB8JP2VPFkmj6FvkkFlNr2tfs93HiC6RXmlZQ2poitJIwQNJIzfCkv7Xf7TH7LK/DRfhD8VU0U6doPh/XtDZvhb8DYD4d/4Z9+MWv8A7ZfgbRtCGi/C/R1sNBvfj18e/idr/ivTlV/+Eg8Pa+fA800PhCBdGPQftH/8l5/aK/7Hwf0r5T/aR6fDz/sQ/il/6rr4dUAfp/8AsW/Bvwn8Sfh342+JWtRwaXrHxE+KOseNdT0Xwp4T+HHhjwZoF5r3hTwXdy6F4H8IaV4Ig0TwX4N0nK2Xhvwh4ftbLQPDmlxQaXo9laWNvDBH9g/8M0eAP+fnVP8AwV+Bf/mMrxf/AIJ1/wDJvFr/ANjBF/6hHgmvu6gD/9k="
)

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

func buntFromConst() (*bunt.BuntDB, error) {
	c := &bunt.Config{
		DBPath: ":memory:",
	}
	return c.InitDB()
}

const (
	MORALIS_API_KEY  = "YNK1oJejsgzJ1L5Gxaszqd1fOH5t5h595ksVu5bvE6nyCDl9Bb7WD7N18Gb6mglz"
	MORALIS_TIMEOUT  = "10s"
	MORALIS_BASE_URL = "https://deep-index.moralis.io/api/v2/"
)

func InitMoralisFromConst() (IPFS, error) {
	cfgMoralis := &Config{
		APIKey:  MORALIS_API_KEY,
		Timeout: MORALIS_TIMEOUT,
		BaseURL: MORALIS_BASE_URL,
	}
	cfgDecription := &descriptions.Config{
		Path:           "etc/descriptions.json",
		CollectionName: "Solutions #",
	}
	desc, err := cfgDecription.Init()
	if err != nil {
		return nil, err
	}
	return cfgMoralis.Init(desc)
}

func getTestMrs() []*pb_nft.NFTMintRequestWithStatus {
	return []*pb_nft.NFTMintRequestWithStatus{
		{
			NftOffchainUrl: testImage,
		},
		{
			NftOffchainUrl: testImage,
		},
		{
			NftOffchainUrl: testImage,
		},
	}
}

func TestUploadToIPFSFolder(t *testing.T) {
	skipCI(t)

	is := is.New(t)
	m, err := InitMoralisFromConst()
	is.NoErr(err)

	resp, err := m.BulkUploadIPFS(getTestMrs())
	is.NoErr(err)
	t.Logf("%+v", resp)

}

type adj struct {
	Adjectives []string `json:"adjectives"`
}

type desc struct {
	Description []string `json:"description"`
	MintNumber  int      `json:"mintNumber"`
}

func TestGenDescriptions(t *testing.T) {
	skipCI(t)
	is := is.New(t)

	a, err := os.Open("../etc/adjectives.json")
	is.NoErr(err)

	bs, err := ioutil.ReadAll(a)
	is.NoErr(err)

	adj := adj{}

	err = json.Unmarshal(bs, &adj)
	is.NoErr(err)

	totalDescriptions := 10000
	wordsInDescription := 3
	descs := make([]desc, totalDescriptions)

	for i := 0; i < totalDescriptions; i++ {
		for j := 0; j < wordsInDescription; j++ {
			descs[i].Description = append(descs[i].Description, adj.Adjectives[rand.Intn(len(adj.Adjectives))])
			descs[i].MintNumber = i + 1
		}
	}

	bs, err = json.Marshal(descs)
	is.NoErr(err)

	f, err := os.Create("../etc/descriptions.json")
	is.NoErr(err)

	_, err = f.WriteString(string(bs))
	is.NoErr(err)

}
