package bucket

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"strings"

	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/minio/minio-go"
)

type PathExtra struct {
	TxHash       string
	EthAddr      string
	MintSequence string
}

// // type Image struct {
// // 	RawB64Image *string `json:"raw,omitempty"`
// // 	FullSize    string  `json:"fullSize"`
// // 	Compressed  string  `json:"compressed"`
// // }

// func (i *Image) Validate() error {
// 	if i == nil {
// 		return fmt.Errorf("missing Image")
// 	}
// 	if len(i.FullSize) == 0 {
// 		return fmt.Errorf("missing Image FullSize")
// 	}
// 	if len(i.Compressed) == 0 {
// 		return fmt.Errorf("missing Image Compressed")
// 	}
// 	return nil
// }

// upload image to bucket return url
func (b *Bucket) UploadImageToBucket(img io.Reader, contentType string, prefix string, pe *PathExtra) (string, error) {

	fp := b.getImageFullPath(fileExtensionFromContentType(contentType), prefix, pe)

	userMetaData := map[string]string{"x-amz-acl": "public-read"} // make it public
	cacheControl := "max-age=31536000"

	bs, _ := ioutil.ReadAll(img)

	r := bytes.NewReader(bs)

	_, err := b.Client.PutObject(b.S3BucketName, fp, r, int64(len(bs)), minio.PutObjectOptions{ContentType: contentType, CacheControl: cacheControl, UserMetadata: userMetaData})
	if err != nil {
		return "", fmt.Errorf("PutObject:err [%v]", err.Error())
	}

	return b.GetCDNURL(fp), nil
}

func GetB64ImageFromString(rawB64Image string) (*B64Image, error) {
	ss := strings.Split(rawB64Image, ";base64,")
	if len(ss) != 2 {
		return nil, fmt.Errorf("GetB64ImageFromString:bad base64 image")
	}
	return &B64Image{
		Content:     []byte(ss[1]),
		ContentType: ss[0],
	}, nil

}

func (b64Img *B64Image) B64ToImage() (image.Image, error) {
	var img image.Image
	var err error
	switch b64Img.ContentType {
	case "data:image/jpeg":
		img, err = jpgFromB64(b64Img.Content)
		if err != nil {
			return nil, fmt.Errorf("B64ToImage:JPGFromB64: [%v]", err.Error())
		}
	case "data:image/png":
		img, err = pngFromB64(b64Img.Content)
		if err != nil {
			return nil, fmt.Errorf("B64ToImage:PNGFromB64: [%v]", err.Error())
		}
	default:
		return nil, fmt.Errorf("B64ToImage:PNGFromB64: File type is not supported [%s]", b64Img.ContentType)
	}
	return img, err
}

func imageFromString(rawB64Image string) (image.Image, error) {
	b64Img, err := GetB64ImageFromString(rawB64Image)
	if err != nil {
		return nil, err
	}
	return b64Img.B64ToImage()
}

// upload single image with defined quality and	prefix to bucket
func (b *Bucket) UploadSingleImage(img image.Image, quality int, prefix string, pe *PathExtra) (string, error) {
	var buf bytes.Buffer
	imgWriter := bufio.NewWriter(&buf)

	err := encodeJPG(imgWriter, img, quality)
	if err != nil {
		return "", fmt.Errorf("Upload:EncodeJPG: [%v]", err.Error())
	}

	imgReader := bufio.NewReader(&buf)
	url, err := b.UploadImageToBucket(imgReader, contentTypeJSON, prefix, pe)
	if err != nil {
		return "", fmt.Errorf("Upload:UploadImageToBucket: [%v]", err.Error())
	}
	return url, nil
}

// compose internal image object (with FullSize & Compressed formats) and upload it to S3
func (b *Bucket) UploadImageObj(img image.Image, pe *PathExtra) (*pb_nft.ImageList, error) {
	imgObj := &pb_nft.ImageList{}
	var err error

	imgObj.FullSize, err = b.UploadSingleImage(img, 100, "og", pe)
	if err != nil {
		return nil, fmt.Errorf("UploadProductImage:Upload:FullSize [%v]", err.Error())
	}

	imgObj.Compressed, err = b.UploadSingleImage(img, 60, "compressed", pe)
	if err != nil {
		return nil, fmt.Errorf("UploadProductImage:Upload:Compressed [%v]", err.Error())
	}
	return imgObj, nil
}

// get raw image from b64 encoded string and upload fullsize and compressed images to s3
func (b *Bucket) UploadContentImage(rawB64Image string, pe *PathExtra) (*pb_nft.ImageList, error) {
	img, err := imageFromString(rawB64Image)
	if err != nil {
		return nil, err
	}
	return b.UploadImageObj(img, pe)
}
